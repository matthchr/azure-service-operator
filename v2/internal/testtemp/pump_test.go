/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package wop_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/go-logr/logr"
	"github.com/onsi/gomega/format"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2/textlogger"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/controllers"
	"github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/arm"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers/generic"
	asocel "github.com/Azure/azure-service-operator/v2/internal/util/cel"
	"github.com/Azure/azure-service-operator/v2/internal/util/interval"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/internal/util/lockedrand"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/registration"

	//nolint:staticcheck // ignoring deprecation (SA1019) to unblock CI builds
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"

	. "github.com/onsi/gomega"
)

//var (
//	globalTestContext LocalTestContext
//)

//func setup() error {
//	log.Println("Running test setup")
//
//	format.TruncateThreshold = 4000 // Force a longer truncate threshold
//
//	nameConfig := testcommon.NewResourceNameConfig(
//		testcommon.ResourcePrefix,
//		"-",
//		6,
//		testcommon.ResourceNamerModeRandomBasedOnTestName)
//
//	// set global context var
//	newGlobalTestContext, err := testcommon.NewKubeContext(
//		options.useEnvTest,
//		options.recordReplay,
//		testcommon.DefaultTestRegion,
//		nameConfig)
//	if err != nil {
//		return err
//	}
//
//	log.Print("Done with test setup")
//	globalTestContext = newGlobalTestContext
//	return nil
//}

//func teardown() error {
//	return globalTestContext.Cleanup()
//}


//func TestMain(m *testing.M) {
//	os.Exit(testcommon.SetupTeardownTestMain(m, setup, teardown))
//}

func newReconciler(config *LocalTestContext, info registration.StorageType) (*generic.GenericReconciler, error) {
	//logger := textlogger.NewLogger(textlogger.NewConfig(textlogger.Verbosity(logging.Debug)))

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(info.Obj, config.scheme)
	if err != nil {
		return nil, errors.Wrapf(err, "creating GVK for obj %T", info.Obj)
	}

	loggerFactory := func(mo genruntime.MetaObject) logr.Logger {
		logger := testcommon.NewTestLogger(config.t)
		return logger.WithName(info.Name)
	}

	broadcaster := record.NewBroadcasterForTests(1 * time.Second)
	recorder := broadcaster.NewRecorder(config.scheme, corev1.EventSource{Component: info.Name})

	requeueDelay := 10 * time.Millisecond
	minBackoff := 5 * time.Millisecond
	maxBackoff := 5 * time.Millisecond

	reconciler := &generic.GenericReconciler{
		Reconciler:         info.Reconciler,
		KubeClient:         config.kubeClient,
		Config:             config.config,
		LoggerFactory:      loggerFactory,
		Recorder:           recorder,
		GVK:                gvk,
		PositiveConditions: conditions.NewPositiveConditionBuilder(clock.New()),
		// Specified here because usually controller-runtime logging would detect panics and log them for us
		// but in the case of envtest we disable those logs because they're too verbose.
		PanicHandler: func() {
			if e := recover(); e != nil {
				stack := debug.Stack()
				log.Printf("panic: %s\nstack:%s\n", e, stack)
			}
		},
		RequeueIntervalCalculator: interval.NewCalculator(
			interval.CalculatorParameters{
				//nolint:gosec // do not want cryptographic randomness here
				Rand:                 rand.New(lockedrand.NewSource(time.Now().UnixNano())),
				ErrorBaseDelay:       minBackoff,
				ErrorMaxFastDelay:    maxBackoff,
				ErrorMaxSlowDelay:    maxBackoff,
				RequeueDelayOverride: requeueDelay,
			}),
	}

	return reconciler, nil
	//
	//var cacheFunc cache.NewCacheFunc
	//if cfg.TargetNamespaces != nil && cfg.OperatorMode.IncludesWatchers() {
	//	cacheFunc = func(config *rest.Config, opts cache.Options) (cache.Cache, error) {
	//		opts.DefaultNamespaces = make(map[string]cache.Config, len(cfg.TargetNamespaces))
	//		for _, ns := range cfg.TargetNamespaces {
	//			opts.DefaultNamespaces[ns] = cache.Config{}
	//		}
	//
	//		return cache.New(config, opts)
	//	}
	//}
	//
	//loggerFactory := func(obj metav1.Object) logr.Logger {
	//	result := namespaceResources.Lookup(obj.GetNamespace())
	//	if result == nil {
	//		panic(fmt.Sprintf("no logger registered for %s: %s", obj.GetNamespace(), obj.GetName()))
	//	}
	//
	//	return result.logger
	//}
	//
	//// We use a custom indexer here so that we can simulate the caching client behavior for indexing even though
	//// for our tests we are not using the caching client
	//testIndexer := NewIndexer(mgr.GetScheme())
	//indexer := kubeclient.NewAndIndexer(mgr.GetFieldIndexer(), testIndexer)
	//kubeClient := kubeclient.NewClient(NewClient(mgr.GetClient(), testIndexer))
	//expressionEvaluator, err := asocel.NewExpressionEvaluator(asocel.Log(logger))
	//// Note that we don't start expressionEvaluator here because we're in a test context and turning cache eviction
	//// on is probably overkill.
	//if err != nil {
	//	return nil, errors.Wrapf(err, "creating expression evaluator")
	//}
	//
	//// This means a single evaluator will be used for all envtests. For the purposes of testing that's probably OK...
	////asocel.RegisterEvaluator(expressionEvaluator)
	//
	//credentialProviderWrapper := &credentialProviderWrapper{namespaceResources: namespaceResources}
	//
	//var clientFactory arm.ARMConnectionFactory = func(ctx context.Context, mo genruntime.ARMMetaObject) (arm.Connection, error) {
	//	result := namespaceResources.Lookup(mo.GetNamespace())
	//	if result == nil {
	//		panic(fmt.Sprintf("unable to locate ARM client for namespace %s; tests should only create resources in the namespace they are assigned or have declared via TargetNamespaces",
	//			mo.GetNamespace()))
	//	}
	//
	//	return result.armClientCache.GetConnection(ctx, mo)
	//}

}

func newLocalTestContext(t *testing.T) *LocalTestContext {
	scheme := controllers.CreateScheme()

	ctrl.SetLogger(logr.Discard())

	fakeClient := fake.NewClientBuilder().WithScheme(scheme).Build()
	kubeClient :=  kubeclient.NewClient(fakeClient)

	config := config.Values{
		ClientID: "1234",
		SubscriptionID: "1234",
		TenantID: "1234",
		// Test configs never want SyncPeriod set as it introduces jitter
		SyncPeriod: nil,
		// Simulate pod namespace being set, as we're not running in a pod context so we don't have this env variable
		// injected automatically
		PodNamespace: "azureserviceoperator-system",
	}

	return &LocalTestContext{
		config: config,
		kubeClient: kubeClient,
		scheme: scheme,
		t: t,
	}
}

type LocalTestContext struct {
	config config.Values
	kubeClient kubeclient.Client
	scheme *runtime.Scheme
	t *testing.T
}

func Test_FindOptionalConfigMapReferences(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := newLocalTestContext(t)
	info :=
	operator := newReconciler(tc, obj)

	operator :=

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref1,
			},
			RefSlice: []genruntime.ResourceReference{
				ref2,
				ref3,
			},
			RefMap: map[string]genruntime.ResourceReference{
				"a": ref4,
				"b": ref5,
			},
			PropertyWithTag: to.Ptr("hello"),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	results, err := reflecthelpers.FindOptionalConfigMapReferences(res)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(results).To(HaveLen(1))
	g.Expect(results[0].Name).To(Equal("Spec.PropertyWithTag"))
	g.Expect(results[0].Value).To(Equal(to.Ptr("hello")))
	g.Expect(results[0].RefName).To(Equal("Spec.PropertyWithTagFromConfig"))
	g.Expect(results[0].Ref).To(Equal((*genruntime.ConfigMapReference)(nil)))
}

// defaultResourceReferencesName exists to showcase an example where ReflectVisitor is used to modify the object in question
func defaultResourceReferencesName(transformer genruntime.ARMTransformer, name string) error {
	visitor := reflecthelpers.NewReflectVisitor()
	visitor.VisitStruct = func(this *reflecthelpers.ReflectVisitor, it reflect.Value, ctx interface{}) error {
		if it.Type() == reflect.TypeOf(genruntime.ResourceReference{}) {
			if it.CanInterface() {
				reference := it.Interface().(genruntime.ResourceReference)
				if reference.Name == "" {
					// Cannot do assignment on the reference variable as it is a copy
					f := it.FieldByName("Name")
					if !f.CanSet() {
						return errors.New("cannot set 'Name' field of 'genruntime.ResourceReference'")
					}
					f.SetString(name)
				}
			} else {
				// This should be impossible given how the visitor works
				return errors.New("genruntime.ResourceReference field was unexpectedly nil")
			}
			return nil
		}

		return reflecthelpers.IdentityVisitStruct(this, it, ctx)
	}

	err := visitor.Visit(transformer, nil)
	if err != nil {
		return errors.Wrap(err, "defaulting genruntime.ResourceReference")
	}

	return nil
}

func Test_CanUseReflectVisitorToModifyResource(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	ref := genruntime.ResourceReference{Group: "microsoft.keyvault", Kind: "keyvault"}

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
			Ref: &ResourceReference{
				Reference: ref,
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	err := defaultResourceReferencesName(&res.Spec, "myname")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(res.Spec.Ref.Reference.Name).To(Equal("myname"))
}

func Test_GetObjectListItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	res := ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	list := &ResourceWithReferencesList{
		Items: []ResourceWithReferences{
			res,
		},
	}

	items, err := reflecthelpers.GetObjectListItems(list)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(items).To(HaveLen(1))
	g.Expect(items[0].GetName()).To(Equal("test-group"))
}

func Test_SetObjectListItems(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	res := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
			Location:  "westus",
			Owner: genruntime.KnownResourceReference{
				Name: "myrg",
			},
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-group",
			Namespace: "test-namespace",
		},
	}

	list := &ResourceWithReferencesList{}

	itemList := []client.Object{res}
	err := reflecthelpers.SetObjectListItems(list, itemList)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(list.Items).To(HaveLen(1))
	g.Expect(list.Items[0].GetName()).To(Equal("test-group"))
}

func Test_SetProperty_TargetingStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName: "azureName",
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "AzureName", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.AzureName).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleProperties_MakesChanges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName:       "azureName",
		Location:        "westus",
		PropertyWithTag: to.Ptr("hello"),
	}

	name := "dorothy"
	location := "land-of-oz"
	property := to.Ptr("flying-house")

	g.Expect(reflecthelpers.SetProperty(subject, "AzureName", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Location", location)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "PropertyWithTag", property)).To(Succeed())

	g.Expect(subject.AzureName).To(Equal(name))
	g.Expect(subject.Location).To(Equal(location))
	g.Expect(subject.PropertyWithTag).To(Equal(property))
}

func Test_SetProperty_TargetingNestedStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "Spec.AzureName", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.Spec.AzureName).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleNestedProperties_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName:       "azureName",
			Location:        "westus",
			PropertyWithTag: to.Ptr("hello"),
		},
	}

	name := "dorothy"
	location := "land-of-oz"
	property := to.Ptr("flying-house")

	g.Expect(reflecthelpers.SetProperty(subject, "Spec.AzureName", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Location", location)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.PropertyWithTag", property)).To(Succeed())

	g.Expect(subject.Spec.AzureName).To(Equal(name))
	g.Expect(subject.Spec.Location).To(Equal(location))
	g.Expect(subject.Spec.PropertyWithTag).To(Equal(property))
}

func Test_SetProperty_TargetingNestedNestedStringProperty_MakesChange(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	newValue := "newName"
	err := reflecthelpers.SetProperty(subject, "Spec.Owner.Name", newValue)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(subject.Spec.Owner.Name).To(Equal(newValue))
}

func Test_SetProperty_TargetingMultipleNestedNestedProperties_MakesChanges(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{}

	name := "lock"
	key := "key"

	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Secret.Name", name)).To(Succeed())
	g.Expect(reflecthelpers.SetProperty(subject, "Spec.Secret.Key", key)).To(Succeed())

	g.Expect(subject.Spec.Secret.Name).To(Equal(name))
	g.Expect(subject.Spec.Secret.Key).To(Equal(key))
}

func Test_SetProperty_TargetingUnknownProperty_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{
		AzureName: "azureName",
	}

	err := reflecthelpers.SetProperty(subject, "UnknownProperty", "newValue")
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("UnknownProperty")))
}

func Test_SetProperty_TargetingUnknownNestedProperty_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferencesSpec{}

	err := reflecthelpers.SetProperty(subject, "Owner.UnknownProperty", "newValue")
	g.Expect(err).To(HaveOccurred())

	g.Expect(err).To(MatchError(ContainSubstring("Owner")))
	g.Expect(err).To(MatchError(ContainSubstring("UnknownProperty")))
}

func Test_SetProperty_WhenValueOfWrongType_ReturnsExpectedError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Spec: ResourceWithReferencesSpec{
			AzureName: "azureName",
		},
	}

	err := reflecthelpers.SetProperty(subject, "Spec.AzureName", make([]int, 0, 10))
	g.Expect(err).To(HaveOccurred())
	g.Expect(err).To(MatchError(ContainSubstring("Spec")))
	g.Expect(err).To(MatchError(ContainSubstring("AzureName")))
	g.Expect(err).To(MatchError(ContainSubstring("kind []")))
	g.Expect(err).To(MatchError(ContainSubstring("not compatible")))
}

func Test_SetProperty_WhenValueOfCompatibleType_ModifiesValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	subject := &ResourceWithReferences{
		Status: ResourceWithReferencesStatus{
			ProvisioningState: to.Ptr(ProvisioningStateSucceeded),
		},
	}

	err := reflecthelpers.SetProperty(subject, "Status.ProvisioningState", to.Ptr(string(ProvisioningStateFailed)))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(*subject.Status.ProvisioningState).To(Equal(ProvisioningStateFailed))
}
