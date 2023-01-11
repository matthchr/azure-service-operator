/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

//import (
//	"testing"
//
//	. "github.com/onsi/gomega"
//	"sigs.k8s.io/controller-runtime/pkg/client"
//
//	"github.com/Azure/go-autorest/autorest/to"
//
//	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
//	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
//	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101"
//	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
//)
//
//// Some examples drawn from: https://learn.microsoft.com/azure/azure-sql/managed-instance/create-template-quickstart?view=azuresql
//
//func Test_SQL_ManagedInstance_CRUD(t *testing.T) {
//	t.Parallel()
//
//	tc := globalTestContext.ForTest(t)
//	// Use a different region where we have quota
//	tc.AzureRegion = to.StringPtr("eastus")
//
//	secretName := "sqlsecret"
//	adminPasswordKey := "adminPassword"
//	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)
//
//	rg := tc.CreateTestResourceGroupAndWait()
//
//	// SQL managed instance must be deployed into a subnet
//	subnet := makeSubnetForSQLManagedInstance(tc, rg)
//
//	licenseType := sql.ManagedInstanceProperties_LicenseType_LicenseIncluded
//	managedInstance := &sql.ManagedInstance{
//		ObjectMeta: tc.MakeObjectMeta("sqlmanagedinstance"),
//		Spec: sql.ManagedInstance_Spec{
//			Location:                   tc.AzureRegion,
//			Owner:                      testcommon.AsOwner(rg),
//			AdministratorLogin:         to.StringPtr("myadmin"),
//			AdministratorLoginPassword: &adminPasswordSecretRef,
//			Sku: &sql.Sku{
//				Name: to.StringPtr("GP_Gen5"),
//			},
//			StorageSizeInGB: to.IntPtr(32),
//			VCores:          to.IntPtr(8),
//			SubnetReference: tc.MakeReferenceFromResource(subnet),
//			LicenseType:     &licenseType,
//		},
//	}
//
//	tc.CreateResourceAndWait(managedInstance)
//	tc.Expect(managedInstance.Status.Id).ToNot(BeNil())
//	tc.Expect(managedInstance.Status.Location).To(Equal(tc.AzureRegion))
//
//	tc.RunParallelSubtests(testcommon.Subtest{
//		Name: "SQL Database CRUD",
//		Test: func(tc *testcommon.KubePerTestContext) {
//			SQL_ManagedInstance_Database_CRUD(tc, managedInstance)
//		},
//	})
//
//	// TODO: Test update (here and in below tests where it makes sense)
//
//	armId := *managedInstance.Status.Id
//	tc.DeleteResourceAndWait(managedInstance)
//
//	exists, _, err := tc.AzureClient.HeadByID(
//		tc.Ctx,
//		armId,
//		string(sql.APIVersion_Value))
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}
//
//func SQL_ManagedInstance_Database_CRUD(tc *testcommon.KubePerTestContext, managedInstance *sql.ManagedInstance) {
//	db := &sql.ManagedInstancesDatabase{
//		ObjectMeta: tc.MakeObjectMeta("db"),
//		Spec: sql.ManagedInstances_Database_Spec{
//			Owner:    testcommon.AsOwner(managedInstance),
//			Location: tc.AzureRegion,
//		},
//	}
//
//	tc.CreateResourceAndWait(db)
//
//	tc.Expect(db.Status.Id).ToNot(BeNil())
//	armId := *managedInstance.Status.Id
//
//	tc.RunParallelSubtests(testcommon.Subtest{
//		Name: "SQL Database BackupLongTermRetentionPolicy CRUD",
//		Test: func(tc *testcommon.KubePerTestContext) {
//			SQL_ManagedInstance_BackupLongTermRetention_CRUD(tc, db)
//		},
//	})
//
//	tc.DeleteResourceAndWait(db)
//
//	exists, _, err := tc.AzureClient.HeadByID(
//		tc.Ctx,
//		armId,
//		string(sql.APIVersion_Value))
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}
//
//func SQL_ManagedInstance_BackupLongTermRetention_CRUD(tc *testcommon.KubePerTestContext, db *sql.ManagedInstancesDatabase) {
//	policy := &sql.ManagedInstancesDatabasesBackupLongTermRetentionPolicy{
//		ObjectMeta: tc.MakeObjectMeta("db"),
//		Spec: sql.ManagedInstances_Databases_BackupLongTermRetentionPolicy_Spec{
//			Owner:           testcommon.AsOwner(db),
//			WeeklyRetention: to.StringPtr("PT4W"),
//		},
//	}
//
//	tc.CreateResourceAndWait(policy)
//	defer tc.DeleteResourceAndWait(policy)
//	tc.Expect(policy.Status.Id).ToNot(BeNil())
//}
//
//func makeSubnetForSQLManagedInstance(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *network.VirtualNetworksSubnet {
//	vnet := &network.VirtualNetwork{
//		ObjectMeta: tc.MakeObjectMeta("vn"),
//		Spec: network.VirtualNetwork_Spec{
//			Owner:    testcommon.AsOwner(rg),
//			Location: tc.AzureRegion,
//			AddressSpace: &network.AddressSpace{
//				AddressPrefixes: []string{"10.0.0.0/16"},
//			},
//		},
//	}
//	tc.CreateResourceAndWait(vnet)
//
//	// TODO: Put all of this setup into a helper so we can hide it
//	routeTable := &network.RouteTable{
//		ObjectMeta: tc.MakeObjectMeta("routetable"),
//		Spec: network.RouteTable_Spec{
//			Owner:    testcommon.AsOwner(rg),
//			Location: tc.AzureRegion,
//		},
//	}
//	nsg := &network.NetworkSecurityGroup{
//		ObjectMeta: tc.MakeObjectMeta("nsg"),
//		Spec: network.NetworkSecurityGroup_Spec{
//			Owner:    testcommon.AsOwner(rg),
//			Location: tc.AzureRegion,
//		},
//	}
//	rules := makeSQLManagedInstanceNSGSecurityRules(tc, nsg)
//
//	tc.CreateResourceAndWait(nsg)
//	tc.CreateResourcesAndWait(rules...)
//
//	subnet := &network.VirtualNetworksSubnet{
//		ObjectMeta: tc.MakeObjectMeta("subnet"),
//		Spec: network.VirtualNetworks_Subnet_Spec{
//			Owner:         testcommon.AsOwner(vnet),
//			AddressPrefix: to.StringPtr("10.0.0.0/24"),
//			Delegations: []network.Delegation{
//				{
//					Name:        to.StringPtr("managedInstanceDelegation"),
//					ServiceName: to.StringPtr("Microsoft.Sql/managedInstances"),
//				},
//			},
//			RouteTable: &network.RouteTableSpec_VirtualNetworks_Subnet_SubResourceEmbedded{
//				Reference: tc.MakeReferenceFromResource(routeTable),
//			},
//			NetworkSecurityGroup: &network.NetworkSecurityGroupSpec_VirtualNetworks_Subnet_SubResourceEmbedded{
//				Reference: tc.MakeReferenceFromResource(nsg),
//			},
//		},
//	}
//	tc.CreateResourcesAndWait(routeTable, subnet)
//
//	return subnet
//}
//
//func makeSQLManagedInstanceNSGSecurityRules(tc *testcommon.KubePerTestContext, nsg *network.NetworkSecurityGroup) []client.Object {
//	tcp := network.SecurityRulePropertiesFormat_Protocol_Tcp
//	allow := network.SecurityRuleAccess_Allow
//	inbound := network.SecurityRuleDirection_Inbound
//	allowTDSInbound := &network.NetworkSecurityGroupsSecurityRule{
//		ObjectMeta: tc.MakeObjectMetaWithName("allowtdsinbound"),
//		Spec: network.NetworkSecurityGroups_SecurityRule_Spec{
//			Owner:                    testcommon.AsOwner(nsg),
//			Description:              to.StringPtr("Allow access to data"),
//			Protocol:                 &tcp,
//			SourcePortRange:          to.StringPtr("*"),
//			DestinationPortRange:     to.StringPtr("1433"),
//			SourceAddressPrefix:      to.StringPtr("VirtualNetwork"),
//			DestinationAddressPrefix: to.StringPtr("*"),
//			Access:                   &allow,
//			Priority:                 to.IntPtr(1000),
//			Direction:                &inbound,
//		},
//	}
//	allowRedirectInbound := &network.NetworkSecurityGroupsSecurityRule{
//		ObjectMeta: tc.MakeObjectMetaWithName("allowredirectinbound"),
//		Spec: network.NetworkSecurityGroups_SecurityRule_Spec{
//			Owner:                    testcommon.AsOwner(nsg),
//			Description:              to.StringPtr("Allow inbound redirect traffic to Managed Instance inside the virtual network"),
//			Protocol:                 &tcp,
//			SourcePortRange:          to.StringPtr("*"),
//			DestinationPortRange:     to.StringPtr("11000-11999"),
//			SourceAddressPrefix:      to.StringPtr("VirtualNetwork"),
//			DestinationAddressPrefix: to.StringPtr("*"),
//			Access:                   &allow,
//			Priority:                 to.IntPtr(1100),
//			Direction:                &inbound,
//		},
//	}
//
//	return []client.Object{allowTDSInbound, allowRedirectInbound}
//}
