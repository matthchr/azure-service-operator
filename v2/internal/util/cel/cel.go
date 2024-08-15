/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package cel

import (
	"reflect"
	"strings"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/ast"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/ext"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// TODO: if we ever make a breaking change here, we should allow users to control which "version" of CEL parsing they
// TODO: get with an annotation or property on the genruntime type.

// TODO:
// - Write CEL documentation.

const (
	SelfIdent   = "self"
	SecretIdent = "secret"
)

func AllowedOutputTypes() []*cel.Type {
	return []*cel.Type{
		cel.StringType,
		cel.MapType(cel.StringType, cel.StringType),
	}
}

// ExpressionEvaluator defines an interface for evaluating expressions based on self (the resource) and an optional
// secret parameter (containing the secrets)
type ExpressionEvaluator interface {
	CompileAndRun(expression string, self any, secret map[string]string) (*ExpressionResult, error)
	FindSecretUsage(expression string, self any) (set.Set[string], error)
}

type ExpressionEvaluatorOption func(e *expressionEvaluator) (*expressionEvaluator, error)

func Cache(cache ProgramCacher) ExpressionEvaluatorOption {
	return func(e *expressionEvaluator) (*expressionEvaluator, error) {
		e.programCache = cache
		return e, nil
	}
}

var _ ExpressionEvaluator = &expressionEvaluator{}

type expressionEvaluator struct {
	programCache ProgramCacher
}

func NewExpressionEvaluator(
	opts ...ExpressionEvaluatorOption,
) (ExpressionEvaluator, error) {
	result := &expressionEvaluator{}

	for _, opt := range opts {
		var err error
		result, err = opt(result)
		if err != nil {
			return nil, err
		}
	}

	if result.programCache == nil {
		envCache := NewEnvCache(NewEnv)
		result.programCache = NewProgramCache(envCache, Compile)
	}

	return result, nil
}

func (e *expressionEvaluator) Start() {
	go e.programCache.Start()
}

func parseStructTag(field reflect.StructField, tag string) string {
	tag, found := field.Tag.Lookup(tag)
	if found {
		splits := strings.Split(tag, ",")
		if len(splits) > 0 {
			// We make the assumption that the leftmost entry in the tag is the name.
			// This seems to be true for most tags that have the concept of a name/key, such as:
			// https://pkg.go.dev/encoding/xml#Marshal
			// https://pkg.go.dev/encoding/json#Marshal
			// https://pkg.go.dev/go.mongodb.org/mongo-driver/bson#hdr-Structs
			// https://pkg.go.dev/gopkg.in/yaml.v2#Marshal
			name := splits[0]
			return name
		}
	}

	return field.Name
}

func EscapeFieldName(field string) string {
	replacer := strings.NewReplacer("__", "__underscores__", ".", "__dot__", "-", "__dash__", "/", "__slash__")
	return replacer.Replace(field)
}

func ParseStructTag(field reflect.StructField) string {
	tag := parseStructTag(field, "json")

	// We follow the same escaping sequences as Kubernetes
	// https://kubernetes.io/docs/reference/using-api/cel/#escaping
	return EscapeFieldName(tag)
}

func NewEnv(resource reflect.Type) (*cel.Env, error) {
	types := findTypesRecursive(resource)
	selfPath := getTypeImportPath(resource)

	// coerce between type slices
	typesList := coerceList(types)
	typesList = append(typesList, ext.ParseStructField(ParseStructTag))

	// The typesList of supported env functions was taken from Kubernetes here:
	// https://kubernetes.io/docs/reference/using-api/cel/#language-overview.
	return cel.NewEnv(
		ext.Strings(),
		cel.DefaultUTCTimeZone(true),
		// Kubernetes specifies HomogenousAggregateLiterals, but it could cause issues for
		// construction of dynamic JSON payloads. This isn't a common operation, but then again
		// neither is having Heterogeneous Aggregate (list|map) literals, as we currently restrict output
		// types to string or map[string]string.
		// cel.HomogeneousAggregateLiterals(),

		// I'm not sure that this actually does anything for us given that all of our declarations are defined as DynType,
		// but I don't think setting it harms anything and it keeps us in-sync with upstream Kubernetes.
		cel.EagerlyValidateDeclarations(true),
		cel.OptionalTypes(),
		cel.CrossTypeNumericComparisons(true),

		ext.NativeTypes(typesList...),
		newJSONProvider(), // This must come after ext.NativeTypes

		// TODO: We could consider adding support for the Kubernetes List Library
		// TODO: https://kubernetes.io/docs/reference/using-api/cel/#kubernetes-list-library
		// TODO: and the Kubernetes Regex Library
		// TODO: https://kubernetes.io/docs/reference/using-api/cel/#kubernetes-regex-library
		cel.Variable("self", cel.ObjectType(selfPath)),
		cel.Variable("secret", cel.MapType(cel.StringType, cel.StringType)), // TODO: is there a better way to do conditional inputs?
	)
	// cel.ClearMacros can be used to disable macros.
	// See https://github.com/google/cel-spec/blob/master/doc/langdef.md#macros
	// In our case many of the macros such as has, all, etc may be useful so we don't
	// disable them. We don't support any custom macros though, only the builtins.
}

// Compile builds the specified expression and returns a ProgramCacheItem containing
// the cel.AST and cel.Program.
func Compile(env *cel.Env, expression string) (*ProgramCacheItem, error) {
	ast, iss := env.Compile(expression)
	if iss.Err() != nil {
		return nil, errors.Wrapf(iss.Err(), "failed to compile CEL expression: %q", expression)
	}

	err := CheckOutputTypeAllowed(ast, AllowedOutputTypes()...)
	if err != nil {
		return nil, err
	}

	program, err := env.Program(ast)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate program from CEL AST: %q", expression)
	}

	return &ProgramCacheItem{
		AST:     ast,
		Program: program,
	}, nil
}

type ExpressionResult struct {
	Value  string
	Values map[string]string
}

// CompileAndRun compiles the specified expression and returns the result of the compilation.
// expression is a CEL expression that must return either a string or map[string]string.
// self is the resource being reconciled.
// secret is the set of secrets associated with the resource (which may be empty).
func (e *expressionEvaluator) CompileAndRun(expression string, self any, secret map[string]string) (*ExpressionResult, error) {
	// Cache lookup also compiles and checks the program for errors
	program, err := e.programCache.Get(reflect.TypeOf(self), expression)
	if err != nil {
		return nil, err
	}

	input := map[string]any{
		SelfIdent:   self,
		SecretIdent: secret,
	}
	// TODO: We may want to use ContextEval here, alongside prgm, err := env.cel.Program(ast, cel.InterruptCheckFrequency(10))
	out, _, err := program.Program.Eval(input)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to eval CEL expression: %q", expression)
	}

	// TODO: When inputs were Dyn, it was hard to statically typecheck based on out.Type. We technically can now
	// TODO: that inputs are well-typed, but there's maybe not much advantage to doing so because we've got access
	// TODO: to the actual output here too. If we compile (but don't execute) the expression in a webhook then we
	// TODO: may want to bring back the type-checking based on CEL output type rather than Golang output type.
	outStr, strOK := out.Value().(string)
	outMap, mapErr := valToMap(out)

	if strOK {
		return &ExpressionResult{
			Value: outStr,
		}, nil
	} else if mapErr == nil {
		return &ExpressionResult{
			Values: outMap,
		}, nil
	}

	// If we get here, the result type was something unexpected
	return nil, makeUnexpectedResultError(program.AST, AllowedOutputTypes()...)
}

func searchExpr(expr ast.Expr) set.Set[string] {
	result := set.Set[string]{}
	searchExprImpl(expr, result)
	return result
}

func searchExprImpl(expr ast.Expr, result set.Set[string]) {
	if expr == nil {
		// Shouldn't happen but let's just be safe
		return
	}

	switch expr.Kind() {
	case ast.CallKind:
		t := expr.AsCall()
		searchExprImpl(t.Target(), result)
		for _, item := range t.Args() {
			searchExprImpl(item, result)
		}
	case ast.UnspecifiedExprKind:
		return
	case ast.ComprehensionKind:
		t := expr.AsComprehension()
		searchExprImpl(t.IterRange(), result)
		searchExprImpl(t.AccuInit(), result)
		searchExprImpl(t.LoopStep(), result)
		searchExprImpl(t.LoopCondition(), result)
	case ast.IdentKind:
		return
	case ast.ListKind:
		t := expr.AsList()
		for _, item := range t.Elements() {
			searchExprImpl(item, result)
		}
	case ast.LiteralKind:
		return
	case ast.MapKind:
		t := expr.AsMap()
		for _, item := range t.Entries() {
			ent := item.AsMapEntry()
			searchExprImpl(ent.Value(), result)
			searchExprImpl(ent.Key(), result)
		}
	case ast.SelectKind:
		t := expr.AsSelect()
		op := t.Operand()
		if op.Kind() != ast.IdentKind {
			searchExprImpl(op, result)
		} else {
			ident := op.AsIdent()
			if ident == SecretIdent {
				result.Add(t.FieldName())
			}
		}
	case ast.StructKind:
		t := expr.AsStruct()
		for _, item := range t.Fields() {
			ent := item.AsStructField()
			searchExprImpl(ent.Value(), result)
		}
	}
}

// FindSecretUsage finds CEL expressions using `secret.xyz` and returns the set of secret values (xyz in the example).
func (e *expressionEvaluator) FindSecretUsage(expression string, self any) (set.Set[string], error) {
	// Cache lookup also compiles and checks the program for errors
	program, err := e.programCache.Get(reflect.TypeOf(self), expression)
	if err != nil {
		return nil, err
	}

	result := searchExpr(program.AST.NativeRep().Expr())
	return result, nil
}

func CheckOutputTypeAllowed(ast *cel.Ast, allowed ...*cel.Type) error {
	matched := false
	for _, t := range allowed {
		if ast.OutputType().IsExactType(t) {
			matched = true
		}
	}

	if !matched {
		return makeUnexpectedResultError(ast, allowed...)
	}

	return nil
}

func makeUnexpectedResultError(ast *cel.Ast, allowed ...*cel.Type) error {
	expectedTypes := lo.Map(
		allowed,
		func(item *cel.Type, _ int) string {
			return item.String()
		})
	expectedTypesStr := strings.Join(expectedTypes, ",")
	return errors.Errorf("expression must return one of [%s], but was %s", expectedTypesStr, ast.OutputType().String())
}

func valToMap(val ref.Val) (map[string]string, error) {
	// Convert the CEL value to a native Go map
	nativeVal, err := val.ConvertToNative(reflect.TypeOf(map[string]string{}))
	if err != nil {
		return nil, errors.Wrap(err, "error converting CEL value to native")
	}

	// Type assert the native value to a map[string]string
	result, ok := nativeVal.(map[string]string)
	if !ok {
		return nil, errors.Errorf("expectedStr map[string]string but got %T", nativeVal)
	}

	return result, nil
}

func simplePkgAlias(pkgPath string) string {
	paths := strings.Split(pkgPath, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[len(paths)-1]
}

// getResourceTypename should return the package name +
func getTypeImportPath(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return simplePkgAlias(t.PkgPath()) + "." + t.Name()
}

func findTypesRecursive(t reflect.Type) []reflect.Type { // Returns any here because that's what cel.Native expects
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var result []reflect.Type

	switch t.Kind() {
	case reflect.Struct:
		// Special case to not include v1.JSON type in the exported reflect data. This is required because in Go
		// this is just a string (a single .Raw field), but we want to transform it to a JSON map in CEL so that
		// users can use it the same way they specified it in the CRD. See json_type_provider.go.
		if t.PkgPath() == "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1" && t.Name() == "JSON" {
			break
		}

		result = append(result, t)

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if !field.IsExported() {
				continue
			}

			fieldType := field.Type
			// Recurse into nested structs
			result = append(result, findTypesRecursive(fieldType)...)
		}
	case reflect.Array, reflect.Slice:
		result = append(result, findTypesRecursive(t.Elem())...)
	case reflect.Map:
		result = append(result, findTypesRecursive(t.Elem())...)
	default:
		// noop on all other types
	}

	return result
}
