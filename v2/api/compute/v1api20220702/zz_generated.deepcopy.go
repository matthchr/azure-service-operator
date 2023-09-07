//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20220702

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiErrorBase_STATUS) DeepCopyInto(out *ApiErrorBase_STATUS) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiErrorBase_STATUS.
func (in *ApiErrorBase_STATUS) DeepCopy() *ApiErrorBase_STATUS {
	if in == nil {
		return nil
	}
	out := new(ApiErrorBase_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiErrorBase_STATUS_ARM) DeepCopyInto(out *ApiErrorBase_STATUS_ARM) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiErrorBase_STATUS_ARM.
func (in *ApiErrorBase_STATUS_ARM) DeepCopy() *ApiErrorBase_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ApiErrorBase_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiError_STATUS) DeepCopyInto(out *ApiError_STATUS) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]ApiErrorBase_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Innererror != nil {
		in, out := &in.Innererror, &out.Innererror
		*out = new(InnerError_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiError_STATUS.
func (in *ApiError_STATUS) DeepCopy() *ApiError_STATUS {
	if in == nil {
		return nil
	}
	out := new(ApiError_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiError_STATUS_ARM) DeepCopyInto(out *ApiError_STATUS_ARM) {
	*out = *in
	if in.Code != nil {
		in, out := &in.Code, &out.Code
		*out = new(string)
		**out = **in
	}
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make([]ApiErrorBase_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Innererror != nil {
		in, out := &in.Innererror, &out.Innererror
		*out = new(InnerError_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiError_STATUS_ARM.
func (in *ApiError_STATUS_ARM) DeepCopy() *ApiError_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(ApiError_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSet) DeepCopyInto(out *DiskEncryptionSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSet.
func (in *DiskEncryptionSet) DeepCopy() *DiskEncryptionSet {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskEncryptionSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSetList) DeepCopyInto(out *DiskEncryptionSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DiskEncryptionSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSetList.
func (in *DiskEncryptionSetList) DeepCopy() *DiskEncryptionSetList {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DiskEncryptionSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSet_STATUS) DeepCopyInto(out *DiskEncryptionSet_STATUS) {
	*out = *in
	if in.ActiveKey != nil {
		in, out := &in.ActiveKey, &out.ActiveKey
		*out = new(KeyForDiskEncryptionSet_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoKeyRotationError != nil {
		in, out := &in.AutoKeyRotationError, &out.AutoKeyRotationError
		*out = new(ApiError_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(DiskEncryptionSetType_STATUS)
		**out = **in
	}
	if in.FederatedClientId != nil {
		in, out := &in.FederatedClientId, &out.FederatedClientId
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(EncryptionSetIdentity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.LastKeyRotationTimestamp != nil {
		in, out := &in.LastKeyRotationTimestamp, &out.LastKeyRotationTimestamp
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
	if in.PreviousKeys != nil {
		in, out := &in.PreviousKeys, &out.PreviousKeys
		*out = make([]KeyForDiskEncryptionSet_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.RotationToLatestKeyVersionEnabled != nil {
		in, out := &in.RotationToLatestKeyVersionEnabled, &out.RotationToLatestKeyVersionEnabled
		*out = new(bool)
		**out = **in
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSet_STATUS.
func (in *DiskEncryptionSet_STATUS) DeepCopy() *DiskEncryptionSet_STATUS {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSet_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSet_STATUS_ARM) DeepCopyInto(out *DiskEncryptionSet_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(EncryptionSetIdentity_STATUS_ARM)
		(*in).DeepCopyInto(*out)
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
		*out = new(EncryptionSetProperties_STATUS_ARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSet_STATUS_ARM.
func (in *DiskEncryptionSet_STATUS_ARM) DeepCopy() *DiskEncryptionSet_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSet_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSet_Spec) DeepCopyInto(out *DiskEncryptionSet_Spec) {
	*out = *in
	if in.ActiveKey != nil {
		in, out := &in.ActiveKey, &out.ActiveKey
		*out = new(KeyForDiskEncryptionSet)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(DiskEncryptionSetType)
		**out = **in
	}
	if in.FederatedClientId != nil {
		in, out := &in.FederatedClientId, &out.FederatedClientId
		*out = new(string)
		**out = **in
	}
	if in.FederatedClientIdFromConfig != nil {
		in, out := &in.FederatedClientIdFromConfig, &out.FederatedClientIdFromConfig
		*out = new(genruntime.ConfigMapReference)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(EncryptionSetIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.RotationToLatestKeyVersionEnabled != nil {
		in, out := &in.RotationToLatestKeyVersionEnabled, &out.RotationToLatestKeyVersionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSet_Spec.
func (in *DiskEncryptionSet_Spec) DeepCopy() *DiskEncryptionSet_Spec {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSet_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSet_Spec_ARM) DeepCopyInto(out *DiskEncryptionSet_Spec_ARM) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(EncryptionSetIdentity_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(EncryptionSetProperties_ARM)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSet_Spec_ARM.
func (in *DiskEncryptionSet_Spec_ARM) DeepCopy() *DiskEncryptionSet_Spec_ARM {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSet_Spec_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity) DeepCopyInto(out *EncryptionSetIdentity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(EncryptionSetIdentity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity.
func (in *EncryptionSetIdentity) DeepCopy() *EncryptionSetIdentity {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity_ARM) DeepCopyInto(out *EncryptionSetIdentity_ARM) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(EncryptionSetIdentity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserAssignedIdentityDetails_ARM, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity_ARM.
func (in *EncryptionSetIdentity_ARM) DeepCopy() *EncryptionSetIdentity_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity_STATUS) DeepCopyInto(out *EncryptionSetIdentity_STATUS) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(EncryptionSetIdentity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]EncryptionSetIdentity_UserAssignedIdentities_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity_STATUS.
func (in *EncryptionSetIdentity_STATUS) DeepCopy() *EncryptionSetIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity_STATUS_ARM) DeepCopyInto(out *EncryptionSetIdentity_STATUS_ARM) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(EncryptionSetIdentity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity_STATUS_ARM.
func (in *EncryptionSetIdentity_STATUS_ARM) DeepCopy() *EncryptionSetIdentity_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity_UserAssignedIdentities_STATUS) DeepCopyInto(out *EncryptionSetIdentity_UserAssignedIdentities_STATUS) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity_UserAssignedIdentities_STATUS.
func (in *EncryptionSetIdentity_UserAssignedIdentities_STATUS) DeepCopy() *EncryptionSetIdentity_UserAssignedIdentities_STATUS {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity_UserAssignedIdentities_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM) DeepCopyInto(out *EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM.
func (in *EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM) DeepCopy() *EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetProperties_ARM) DeepCopyInto(out *EncryptionSetProperties_ARM) {
	*out = *in
	if in.ActiveKey != nil {
		in, out := &in.ActiveKey, &out.ActiveKey
		*out = new(KeyForDiskEncryptionSet_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(DiskEncryptionSetType)
		**out = **in
	}
	if in.FederatedClientId != nil {
		in, out := &in.FederatedClientId, &out.FederatedClientId
		*out = new(string)
		**out = **in
	}
	if in.RotationToLatestKeyVersionEnabled != nil {
		in, out := &in.RotationToLatestKeyVersionEnabled, &out.RotationToLatestKeyVersionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetProperties_ARM.
func (in *EncryptionSetProperties_ARM) DeepCopy() *EncryptionSetProperties_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetProperties_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSetProperties_STATUS_ARM) DeepCopyInto(out *EncryptionSetProperties_STATUS_ARM) {
	*out = *in
	if in.ActiveKey != nil {
		in, out := &in.ActiveKey, &out.ActiveKey
		*out = new(KeyForDiskEncryptionSet_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoKeyRotationError != nil {
		in, out := &in.AutoKeyRotationError, &out.AutoKeyRotationError
		*out = new(ApiError_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
	if in.EncryptionType != nil {
		in, out := &in.EncryptionType, &out.EncryptionType
		*out = new(DiskEncryptionSetType_STATUS)
		**out = **in
	}
	if in.FederatedClientId != nil {
		in, out := &in.FederatedClientId, &out.FederatedClientId
		*out = new(string)
		**out = **in
	}
	if in.LastKeyRotationTimestamp != nil {
		in, out := &in.LastKeyRotationTimestamp, &out.LastKeyRotationTimestamp
		*out = new(string)
		**out = **in
	}
	if in.PreviousKeys != nil {
		in, out := &in.PreviousKeys, &out.PreviousKeys
		*out = make([]KeyForDiskEncryptionSet_STATUS_ARM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.RotationToLatestKeyVersionEnabled != nil {
		in, out := &in.RotationToLatestKeyVersionEnabled, &out.RotationToLatestKeyVersionEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSetProperties_STATUS_ARM.
func (in *EncryptionSetProperties_STATUS_ARM) DeepCopy() *EncryptionSetProperties_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(EncryptionSetProperties_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerError_STATUS) DeepCopyInto(out *InnerError_STATUS) {
	*out = *in
	if in.Errordetail != nil {
		in, out := &in.Errordetail, &out.Errordetail
		*out = new(string)
		**out = **in
	}
	if in.Exceptiontype != nil {
		in, out := &in.Exceptiontype, &out.Exceptiontype
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerError_STATUS.
func (in *InnerError_STATUS) DeepCopy() *InnerError_STATUS {
	if in == nil {
		return nil
	}
	out := new(InnerError_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InnerError_STATUS_ARM) DeepCopyInto(out *InnerError_STATUS_ARM) {
	*out = *in
	if in.Errordetail != nil {
		in, out := &in.Errordetail, &out.Errordetail
		*out = new(string)
		**out = **in
	}
	if in.Exceptiontype != nil {
		in, out := &in.Exceptiontype, &out.Exceptiontype
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InnerError_STATUS_ARM.
func (in *InnerError_STATUS_ARM) DeepCopy() *InnerError_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(InnerError_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyForDiskEncryptionSet) DeepCopyInto(out *KeyForDiskEncryptionSet) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.KeyUrlFromConfig != nil {
		in, out := &in.KeyUrlFromConfig, &out.KeyUrlFromConfig
		*out = new(genruntime.ConfigMapReference)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyForDiskEncryptionSet.
func (in *KeyForDiskEncryptionSet) DeepCopy() *KeyForDiskEncryptionSet {
	if in == nil {
		return nil
	}
	out := new(KeyForDiskEncryptionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyForDiskEncryptionSet_ARM) DeepCopyInto(out *KeyForDiskEncryptionSet_ARM) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyForDiskEncryptionSet_ARM.
func (in *KeyForDiskEncryptionSet_ARM) DeepCopy() *KeyForDiskEncryptionSet_ARM {
	if in == nil {
		return nil
	}
	out := new(KeyForDiskEncryptionSet_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyForDiskEncryptionSet_STATUS) DeepCopyInto(out *KeyForDiskEncryptionSet_STATUS) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyForDiskEncryptionSet_STATUS.
func (in *KeyForDiskEncryptionSet_STATUS) DeepCopy() *KeyForDiskEncryptionSet_STATUS {
	if in == nil {
		return nil
	}
	out := new(KeyForDiskEncryptionSet_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyForDiskEncryptionSet_STATUS_ARM) DeepCopyInto(out *KeyForDiskEncryptionSet_STATUS_ARM) {
	*out = *in
	if in.KeyUrl != nil {
		in, out := &in.KeyUrl, &out.KeyUrl
		*out = new(string)
		**out = **in
	}
	if in.SourceVault != nil {
		in, out := &in.SourceVault, &out.SourceVault
		*out = new(SourceVault_STATUS_ARM)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyForDiskEncryptionSet_STATUS_ARM.
func (in *KeyForDiskEncryptionSet_STATUS_ARM) DeepCopy() *KeyForDiskEncryptionSet_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(KeyForDiskEncryptionSet_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault) DeepCopyInto(out *SourceVault) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault.
func (in *SourceVault) DeepCopy() *SourceVault {
	if in == nil {
		return nil
	}
	out := new(SourceVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault_ARM) DeepCopyInto(out *SourceVault_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault_ARM.
func (in *SourceVault_ARM) DeepCopy() *SourceVault_ARM {
	if in == nil {
		return nil
	}
	out := new(SourceVault_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault_STATUS) DeepCopyInto(out *SourceVault_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault_STATUS.
func (in *SourceVault_STATUS) DeepCopy() *SourceVault_STATUS {
	if in == nil {
		return nil
	}
	out := new(SourceVault_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceVault_STATUS_ARM) DeepCopyInto(out *SourceVault_STATUS_ARM) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceVault_STATUS_ARM.
func (in *SourceVault_STATUS_ARM) DeepCopy() *SourceVault_STATUS_ARM {
	if in == nil {
		return nil
	}
	out := new(SourceVault_STATUS_ARM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails) DeepCopyInto(out *UserAssignedIdentityDetails) {
	*out = *in
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityDetails.
func (in *UserAssignedIdentityDetails) DeepCopy() *UserAssignedIdentityDetails {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails_ARM) DeepCopyInto(out *UserAssignedIdentityDetails_ARM) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityDetails_ARM.
func (in *UserAssignedIdentityDetails_ARM) DeepCopy() *UserAssignedIdentityDetails_ARM {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityDetails_ARM)
	in.DeepCopyInto(out)
	return out
}