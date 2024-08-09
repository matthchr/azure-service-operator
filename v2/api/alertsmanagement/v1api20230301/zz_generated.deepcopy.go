//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20230301

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRule) DeepCopyInto(out *PrometheusRule) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]PrometheusRuleGroupAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.For != nil {
		in, out := &in.For, &out.For
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Record != nil {
		in, out := &in.Record, &out.Record
		*out = new(string)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(PrometheusRuleResolveConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRule.
func (in *PrometheusRule) DeepCopy() *PrometheusRule {
	if in == nil {
		return nil
	}
	out := new(PrometheusRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroup) DeepCopyInto(out *PrometheusRuleGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroup.
func (in *PrometheusRuleGroup) DeepCopy() *PrometheusRuleGroup {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusRuleGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupAction) DeepCopyInto(out *PrometheusRuleGroupAction) {
	*out = *in
	if in.ActionGroupReference != nil {
		in, out := &in.ActionGroupReference, &out.ActionGroupReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupAction.
func (in *PrometheusRuleGroupAction) DeepCopy() *PrometheusRuleGroupAction {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupAction_ARM) DeepCopyInto(out *PrometheusRuleGroupAction_ARM) {
	*out = *in
	if in.ActionGroupId != nil {
		in, out := &in.ActionGroupId, &out.ActionGroupId
		*out = new(string)
		**out = **in
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupAction_ARM.
func (in *PrometheusRuleGroupAction_ARM) DeepCopy() *PrometheusRuleGroupAction_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupAction_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupAction_STATUS) DeepCopyInto(out *PrometheusRuleGroupAction_STATUS) {
	*out = *in
	if in.ActionGroupId != nil {
		in, out := &in.ActionGroupId, &out.ActionGroupId
		*out = new(string)
		**out = **in
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupAction_STATUS.
func (in *PrometheusRuleGroupAction_STATUS) DeepCopy() *PrometheusRuleGroupAction_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupAction_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupAction_STATUS_ARM) DeepCopyInto(out *PrometheusRuleGroupAction_STATUS_ARM) {
	*out = *in
	if in.ActionGroupId != nil {
		in, out := &in.ActionGroupId, &out.ActionGroupId
		*out = new(string)
		**out = **in
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupAction_STATUS_ARM.
func (in *PrometheusRuleGroupAction_STATUS_ARM) DeepCopy() *PrometheusRuleGroupAction_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupAction_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupList) DeepCopyInto(out *PrometheusRuleGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusRuleGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupList.
func (in *PrometheusRuleGroupList) DeepCopy() *PrometheusRuleGroupList {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusRuleGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupOperatorSpec) DeepCopyInto(out *PrometheusRuleGroupOperatorSpec) {
	*out = *in
	if in.ConfigMapExpressions != nil {
		in, out := &in.ConfigMapExpressions, &out.ConfigMapExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
	if in.SecretExpressions != nil {
		in, out := &in.SecretExpressions, &out.SecretExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupOperatorSpec.
func (in *PrometheusRuleGroupOperatorSpec) DeepCopy() *PrometheusRuleGroupOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupProperties_ARM) DeepCopyInto(out *PrometheusRuleGroupProperties_ARM) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PrometheusRule_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupProperties_ARM.
func (in *PrometheusRuleGroupProperties_ARM) DeepCopy() *PrometheusRuleGroupProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroupProperties_STATUS_ARM) DeepCopyInto(out *PrometheusRuleGroupProperties_STATUS_ARM) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PrometheusRule_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroupProperties_STATUS_ARM.
func (in *PrometheusRuleGroupProperties_STATUS_ARM) DeepCopy() *PrometheusRuleGroupProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroupProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroup_STATUS) DeepCopyInto(out *PrometheusRuleGroup_STATUS) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PrometheusRule_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroup_STATUS.
func (in *PrometheusRuleGroup_STATUS) DeepCopy() *PrometheusRuleGroup_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroup_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroup_STATUS_ARM) DeepCopyInto(out *PrometheusRuleGroup_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PrometheusRuleGroupProperties_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroup_STATUS_ARM.
func (in *PrometheusRuleGroup_STATUS_ARM) DeepCopy() *PrometheusRuleGroup_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroup_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroup_Spec) DeepCopyInto(out *PrometheusRuleGroup_Spec) {
	*out = *in
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(PrometheusRuleGroupOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]PrometheusRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScopesReferences != nil {
		in, out := &in.ScopesReferences, &out.ScopesReferences
		*out = make([]genruntime.ResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroup_Spec.
func (in *PrometheusRuleGroup_Spec) DeepCopy() *PrometheusRuleGroup_Spec {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroup_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleGroup_Spec_ARM) DeepCopyInto(out *PrometheusRuleGroup_Spec_ARM) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(PrometheusRuleGroupProperties_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleGroup_Spec_ARM.
func (in *PrometheusRuleGroup_Spec_ARM) DeepCopy() *PrometheusRuleGroup_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleGroup_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleResolveConfiguration) DeepCopyInto(out *PrometheusRuleResolveConfiguration) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleResolveConfiguration.
func (in *PrometheusRuleResolveConfiguration) DeepCopy() *PrometheusRuleResolveConfiguration {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleResolveConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleResolveConfiguration_ARM) DeepCopyInto(out *PrometheusRuleResolveConfiguration_ARM) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleResolveConfiguration_ARM.
func (in *PrometheusRuleResolveConfiguration_ARM) DeepCopy() *PrometheusRuleResolveConfiguration_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleResolveConfiguration_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleResolveConfiguration_STATUS) DeepCopyInto(out *PrometheusRuleResolveConfiguration_STATUS) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleResolveConfiguration_STATUS.
func (in *PrometheusRuleResolveConfiguration_STATUS) DeepCopy() *PrometheusRuleResolveConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleResolveConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRuleResolveConfiguration_STATUS_ARM) DeepCopyInto(out *PrometheusRuleResolveConfiguration_STATUS_ARM) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRuleResolveConfiguration_STATUS_ARM.
func (in *PrometheusRuleResolveConfiguration_STATUS_ARM) DeepCopy() *PrometheusRuleResolveConfiguration_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRuleResolveConfiguration_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRule_ARM) DeepCopyInto(out *PrometheusRule_ARM) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]PrometheusRuleGroupAction_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.For != nil {
		in, out := &in.For, &out.For
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Record != nil {
		in, out := &in.Record, &out.Record
		*out = new(string)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(PrometheusRuleResolveConfiguration_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRule_ARM.
func (in *PrometheusRule_ARM) DeepCopy() *PrometheusRule_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRule_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRule_STATUS) DeepCopyInto(out *PrometheusRule_STATUS) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]PrometheusRuleGroupAction_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.For != nil {
		in, out := &in.For, &out.For
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Record != nil {
		in, out := &in.Record, &out.Record
		*out = new(string)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(PrometheusRuleResolveConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRule_STATUS.
func (in *PrometheusRule_STATUS) DeepCopy() *PrometheusRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(PrometheusRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusRule_STATUS_ARM) DeepCopyInto(out *PrometheusRule_STATUS_ARM) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]PrometheusRuleGroupAction_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Alert != nil {
		in, out := &in.Alert, &out.Alert
		*out = new(string)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.For != nil {
		in, out := &in.For, &out.For
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Record != nil {
		in, out := &in.Record, &out.Record
		*out = new(string)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(PrometheusRuleResolveConfiguration_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusRule_STATUS_ARM.
func (in *PrometheusRule_STATUS_ARM) DeepCopy() *PrometheusRule_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(PrometheusRule_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS_ARM) DeepCopyInto(out *SystemData_STATUS_ARM) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS_ARM)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS_ARM)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS_ARM.
func (in *SystemData_STATUS_ARM) DeepCopy() *SystemData_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}
