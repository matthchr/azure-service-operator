//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CMKIdentityDefinition) DeepCopyInto(out *CMKIdentityDefinition) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UserAssignedIdentityReference != nil {
		in, out := &in.UserAssignedIdentityReference, &out.UserAssignedIdentityReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CMKIdentityDefinition.
func (in *CMKIdentityDefinition) DeepCopy() *CMKIdentityDefinition {
	if in == nil {
		return nil
	}
	out := new(CMKIdentityDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CMKIdentityDefinition_STATUS) DeepCopyInto(out *CMKIdentityDefinition_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UserAssignedIdentity != nil {
		in, out := &in.UserAssignedIdentity, &out.UserAssignedIdentity
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CMKIdentityDefinition_STATUS.
func (in *CMKIdentityDefinition_STATUS) DeepCopy() *CMKIdentityDefinition_STATUS {
	if in == nil {
		return nil
	}
	out := new(CMKIdentityDefinition_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(CMKIdentityDefinition)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.KeyVersion != nil {
		in, out := &in.KeyVersion, &out.KeyVersion
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VaultBaseUrl != nil {
		in, out := &in.VaultBaseUrl, &out.VaultBaseUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration_STATUS) DeepCopyInto(out *EncryptionConfiguration_STATUS) {
	*out = *in
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(CMKIdentityDefinition_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.KeyVersion != nil {
		in, out := &in.KeyVersion, &out.KeyVersion
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.VaultBaseUrl != nil {
		in, out := &in.VaultBaseUrl, &out.VaultBaseUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration_STATUS.
func (in *EncryptionConfiguration_STATUS) DeepCopy() *EncryptionConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Factory) DeepCopyInto(out *Factory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Factory.
func (in *Factory) DeepCopy() *Factory {
	if in == nil {
		return nil
	}
	out := new(Factory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Factory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryGitHubConfiguration) DeepCopyInto(out *FactoryGitHubConfiguration) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(GitHubClientSecret)
		(*in).DeepCopyInto(*out)
	}
	if in.CollaborationBranch != nil {
		in, out := &in.CollaborationBranch, &out.CollaborationBranch
		*out = new(string)
		**out = **in
	}
	if in.DisablePublish != nil {
		in, out := &in.DisablePublish, &out.DisablePublish
		*out = new(bool)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.LastCommitId != nil {
		in, out := &in.LastCommitId, &out.LastCommitId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryGitHubConfiguration.
func (in *FactoryGitHubConfiguration) DeepCopy() *FactoryGitHubConfiguration {
	if in == nil {
		return nil
	}
	out := new(FactoryGitHubConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryGitHubConfiguration_STATUS) DeepCopyInto(out *FactoryGitHubConfiguration_STATUS) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.ClientSecret != nil {
		in, out := &in.ClientSecret, &out.ClientSecret
		*out = new(GitHubClientSecret_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.CollaborationBranch != nil {
		in, out := &in.CollaborationBranch, &out.CollaborationBranch
		*out = new(string)
		**out = **in
	}
	if in.DisablePublish != nil {
		in, out := &in.DisablePublish, &out.DisablePublish
		*out = new(bool)
		**out = **in
	}
	if in.HostName != nil {
		in, out := &in.HostName, &out.HostName
		*out = new(string)
		**out = **in
	}
	if in.LastCommitId != nil {
		in, out := &in.LastCommitId, &out.LastCommitId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryGitHubConfiguration_STATUS.
func (in *FactoryGitHubConfiguration_STATUS) DeepCopy() *FactoryGitHubConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(FactoryGitHubConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryIdentity) DeepCopyInto(out *FactoryIdentity) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryIdentity.
func (in *FactoryIdentity) DeepCopy() *FactoryIdentity {
	if in == nil {
		return nil
	}
	out := new(FactoryIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryIdentity_STATUS) DeepCopyInto(out *FactoryIdentity_STATUS) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryIdentity_STATUS.
func (in *FactoryIdentity_STATUS) DeepCopy() *FactoryIdentity_STATUS {
	if in == nil {
		return nil
	}
	out := new(FactoryIdentity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryList) DeepCopyInto(out *FactoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Factory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryList.
func (in *FactoryList) DeepCopy() *FactoryList {
	if in == nil {
		return nil
	}
	out := new(FactoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FactoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryOperatorSpec) DeepCopyInto(out *FactoryOperatorSpec) {
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryOperatorSpec.
func (in *FactoryOperatorSpec) DeepCopy() *FactoryOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(FactoryOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryRepoConfiguration) DeepCopyInto(out *FactoryRepoConfiguration) {
	*out = *in
	if in.FactoryGitHub != nil {
		in, out := &in.FactoryGitHub, &out.FactoryGitHub
		*out = new(FactoryGitHubConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.FactoryVSTS != nil {
		in, out := &in.FactoryVSTS, &out.FactoryVSTS
		*out = new(FactoryVSTSConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryRepoConfiguration.
func (in *FactoryRepoConfiguration) DeepCopy() *FactoryRepoConfiguration {
	if in == nil {
		return nil
	}
	out := new(FactoryRepoConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryRepoConfiguration_STATUS) DeepCopyInto(out *FactoryRepoConfiguration_STATUS) {
	*out = *in
	if in.FactoryGitHub != nil {
		in, out := &in.FactoryGitHub, &out.FactoryGitHub
		*out = new(FactoryGitHubConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.FactoryVSTS != nil {
		in, out := &in.FactoryVSTS, &out.FactoryVSTS
		*out = new(FactoryVSTSConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryRepoConfiguration_STATUS.
func (in *FactoryRepoConfiguration_STATUS) DeepCopy() *FactoryRepoConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(FactoryRepoConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryVSTSConfiguration) DeepCopyInto(out *FactoryVSTSConfiguration) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.CollaborationBranch != nil {
		in, out := &in.CollaborationBranch, &out.CollaborationBranch
		*out = new(string)
		**out = **in
	}
	if in.DisablePublish != nil {
		in, out := &in.DisablePublish, &out.DisablePublish
		*out = new(bool)
		**out = **in
	}
	if in.LastCommitId != nil {
		in, out := &in.LastCommitId, &out.LastCommitId
		*out = new(string)
		**out = **in
	}
	if in.ProjectName != nil {
		in, out := &in.ProjectName, &out.ProjectName
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
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
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryVSTSConfiguration.
func (in *FactoryVSTSConfiguration) DeepCopy() *FactoryVSTSConfiguration {
	if in == nil {
		return nil
	}
	out := new(FactoryVSTSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FactoryVSTSConfiguration_STATUS) DeepCopyInto(out *FactoryVSTSConfiguration_STATUS) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.CollaborationBranch != nil {
		in, out := &in.CollaborationBranch, &out.CollaborationBranch
		*out = new(string)
		**out = **in
	}
	if in.DisablePublish != nil {
		in, out := &in.DisablePublish, &out.DisablePublish
		*out = new(bool)
		**out = **in
	}
	if in.LastCommitId != nil {
		in, out := &in.LastCommitId, &out.LastCommitId
		*out = new(string)
		**out = **in
	}
	if in.ProjectName != nil {
		in, out := &in.ProjectName, &out.ProjectName
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RepositoryName != nil {
		in, out := &in.RepositoryName, &out.RepositoryName
		*out = new(string)
		**out = **in
	}
	if in.RootFolder != nil {
		in, out := &in.RootFolder, &out.RootFolder
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
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FactoryVSTSConfiguration_STATUS.
func (in *FactoryVSTSConfiguration_STATUS) DeepCopy() *FactoryVSTSConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(FactoryVSTSConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Factory_STATUS) DeepCopyInto(out *Factory_STATUS) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.ETag != nil {
		in, out := &in.ETag, &out.ETag
		*out = new(string)
		**out = **in
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.GlobalParameters != nil {
		in, out := &in.GlobalParameters, &out.GlobalParameters
		*out = make(map[string]GlobalParameterSpecification_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(FactoryIdentity_STATUS)
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
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PurviewConfiguration != nil {
		in, out := &in.PurviewConfiguration, &out.PurviewConfiguration
		*out = new(PurviewConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.RepoConfiguration != nil {
		in, out := &in.RepoConfiguration, &out.RepoConfiguration
		*out = new(FactoryRepoConfiguration_STATUS)
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
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Factory_STATUS.
func (in *Factory_STATUS) DeepCopy() *Factory_STATUS {
	if in == nil {
		return nil
	}
	out := new(Factory_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Factory_Spec) DeepCopyInto(out *Factory_Spec) {
	*out = *in
	if in.AdditionalProperties != nil {
		in, out := &in.AdditionalProperties, &out.AdditionalProperties
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Encryption != nil {
		in, out := &in.Encryption, &out.Encryption
		*out = new(EncryptionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.GlobalParameters != nil {
		in, out := &in.GlobalParameters, &out.GlobalParameters
		*out = make(map[string]GlobalParameterSpecification, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(FactoryIdentity)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(FactoryOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PublicNetworkAccess != nil {
		in, out := &in.PublicNetworkAccess, &out.PublicNetworkAccess
		*out = new(string)
		**out = **in
	}
	if in.PurviewConfiguration != nil {
		in, out := &in.PurviewConfiguration, &out.PurviewConfiguration
		*out = new(PurviewConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.RepoConfiguration != nil {
		in, out := &in.RepoConfiguration, &out.RepoConfiguration
		*out = new(FactoryRepoConfiguration)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Factory_Spec.
func (in *Factory_Spec) DeepCopy() *Factory_Spec {
	if in == nil {
		return nil
	}
	out := new(Factory_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubClientSecret) DeepCopyInto(out *GitHubClientSecret) {
	*out = *in
	if in.ByoaSecretAkvUrl != nil {
		in, out := &in.ByoaSecretAkvUrl, &out.ByoaSecretAkvUrl
		*out = new(string)
		**out = **in
	}
	if in.ByoaSecretName != nil {
		in, out := &in.ByoaSecretName, &out.ByoaSecretName
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubClientSecret.
func (in *GitHubClientSecret) DeepCopy() *GitHubClientSecret {
	if in == nil {
		return nil
	}
	out := new(GitHubClientSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubClientSecret_STATUS) DeepCopyInto(out *GitHubClientSecret_STATUS) {
	*out = *in
	if in.ByoaSecretAkvUrl != nil {
		in, out := &in.ByoaSecretAkvUrl, &out.ByoaSecretAkvUrl
		*out = new(string)
		**out = **in
	}
	if in.ByoaSecretName != nil {
		in, out := &in.ByoaSecretName, &out.ByoaSecretName
		*out = new(string)
		**out = **in
	}
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubClientSecret_STATUS.
func (in *GitHubClientSecret_STATUS) DeepCopy() *GitHubClientSecret_STATUS {
	if in == nil {
		return nil
	}
	out := new(GitHubClientSecret_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalParameterSpecification) DeepCopyInto(out *GlobalParameterSpecification) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalParameterSpecification.
func (in *GlobalParameterSpecification) DeepCopy() *GlobalParameterSpecification {
	if in == nil {
		return nil
	}
	out := new(GlobalParameterSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalParameterSpecification_STATUS) DeepCopyInto(out *GlobalParameterSpecification_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make(map[string]v1.JSON, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalParameterSpecification_STATUS.
func (in *GlobalParameterSpecification_STATUS) DeepCopy() *GlobalParameterSpecification_STATUS {
	if in == nil {
		return nil
	}
	out := new(GlobalParameterSpecification_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurviewConfiguration) DeepCopyInto(out *PurviewConfiguration) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PurviewResourceReference != nil {
		in, out := &in.PurviewResourceReference, &out.PurviewResourceReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurviewConfiguration.
func (in *PurviewConfiguration) DeepCopy() *PurviewConfiguration {
	if in == nil {
		return nil
	}
	out := new(PurviewConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PurviewConfiguration_STATUS) DeepCopyInto(out *PurviewConfiguration_STATUS) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PurviewResourceId != nil {
		in, out := &in.PurviewResourceId, &out.PurviewResourceId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PurviewConfiguration_STATUS.
func (in *PurviewConfiguration_STATUS) DeepCopy() *PurviewConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(PurviewConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails) DeepCopyInto(out *UserAssignedIdentityDetails) {
	*out = *in
	if in.PropertyBag != nil {
		in, out := &in.PropertyBag, &out.PropertyBag
		*out = make(genruntime.PropertyBag, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
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
