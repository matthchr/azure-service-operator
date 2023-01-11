/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1beta20211101"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-logical-server/azuredeploy.json

func Test_SQL_Server_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Use a different region where we have quota
	tc.AzureRegion = to.StringPtr("eastus")

	secretName := "sqlsecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	rg := tc.CreateTestResourceGroupAndWait()

	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location:                   tc.AzureRegion,
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.StringPtr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Version:                    to.StringPtr("12.0"),
		},
	}

	tc.CreateResourceAndWait(server)
	tc.Expect(server.Status.Id).ToNot(BeNil())
	tc.Expect(server.Status.Location).To(Equal(tc.AzureRegion))

	storageDetails := makeStorageAccountForSQLVulnerabilityAssessment(tc, rg)

	// Need to run some setup first for some of the follow subtests to pass

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL Database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Databases_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL ConnectionPolicy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ConnectionPolicy_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL AdvancedThreatDetection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_AdvancedThreatProtection_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VulnerabilityAssessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VulnerabilityAssessments_CRUD(tc, server, rg, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_FirewallRules_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL ElasticPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ElasticPool_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VirtualNetworkRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VirtualNetworkRule_CRUD(tc, server, rg)
			},
		})

	// TODO: Test update (here and in below tests where it makes sense)

	armId := *server.Status.Id
	tc.DeleteResourceAndWait(server)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func SQL_Server_ConnectionPolicy_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	connectionType := sql.ServerConnectionPolicyProperties_ConnectionType_Default
	policy := &sql.ServersConnectionPolicy{
		ObjectMeta: tc.MakeObjectMeta("connpolicy"),
		Spec: sql.Servers_ConnectionPolicy_Spec{
			Owner:          testcommon.AsOwner(server),
			ConnectionType: &connectionType,
		},
	}

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())
	//armId := *policy.Status.Id

	// TOOD: Delete is not allowed
	//tc.DeleteResourceAndWait(policy)
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

func SQL_Server_AdvancedThreatProtection_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	advancedThreatProtection := sql.AdvancedThreatProtectionProperties_State_Enabled
	policy := &sql.ServersAdvancedThreatProtectionSetting{
		ObjectMeta: tc.MakeObjectMeta("atp"),
		Spec: sql.Servers_AdvancedThreatProtectionSetting_Spec{
			Owner: testcommon.AsOwner(server),
			State: &advancedThreatProtection,
		},
	}

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())
	//armId := *policy.Status.Id

	// tc.DeleteResourceAndWait(policy) TODO: Doesn't support delete
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

func SQL_Server_VulnerabilityAssessments_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup, storageDetails vulnStorageAccountDetails) {
	enabled := sql.ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled
	policy := &sql.ServersSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.Servers_SecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(server),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.Servers_VulnerabilityAssessment_Spec{
			Owner: testcommon.AsOwner(server),
			RecurringScans: &sql.VulnerabilityAssessmentRecurringScansProperties{
				IsEnabled: to.BoolPtr(false),
			},
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			// TODO: Make this easier to build a combined path in ASO itself
			StorageContainerPath: to.StringPtr(fmt.Sprintf("%s%s", blobEndpoint, storageDetails.container)),
		},
	}

	tc.CreateResourceAndWait(vulnerabilityAssessment)

	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(vulnerabilityAssessment)
	// TODO: It seems like delete of this resource isn't actually honored by the service
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

func SQL_Server_FirewallRules_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	firewall := &sql.ServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.Servers_FirewallRule_Spec{
			Owner:          testcommon.AsOwner(server),
			StartIpAddress: to.StringPtr("0.0.0.0"),
			EndIpAddress:   to.StringPtr("0.0.0.0"),
		},
	}

	tc.CreateResourceAndWait(firewall)

	tc.Expect(firewall.Status.Id).ToNot(BeNil())
	armId := *firewall.Status.Id

	tc.DeleteResourceAndWait(firewall)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

//
//func SQL_Server_Administrator_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
//	mi := &managedidentity2018.UserAssignedIdentity{
//		ObjectMeta: tc.MakeObjectMeta("mi"),
//		Spec: managedidentity2018.UserAssignedIdentity_Spec{
//			Location: tc.AzureRegion,
//			Owner:    testcommon.AsOwner(rg),
//			OperatorSpec: &managedidentity2018.UserAssignedIdentityOperatorSpec{
//				ConfigMaps: &managedidentity2018.UserAssignedIdentityOperatorConfigMaps{
//					PrincipalId: &genruntime.ConfigMapDestination{
//						Name: "mimap",
//						Key:  "oid",
//					},
//					TenantId: &genruntime.ConfigMapDestination{
//						Name: "mimap",
//						Key:  "tenant",
//					},
//				},
//			},
//		},
//	}
//
//	tc.CreateResourceAndWait(mi)
//
//	adminType := sql.AdministratorProperties_AdministratorType_ActiveDirectory
//	firewall := &sql.ServersAdministrator{
//		ObjectMeta: tc.MakeObjectMeta("atp"),
//		Spec: sql.Servers_Administrator_Spec{
//			Owner:          testcommon.AsOwner(server),
//			SidFromConfig: &genruntime.ConfigMapReference{
//				Name: "mimap",
//				Key:  "oid",
//			},
//			TenantIdFromConfig: &genruntime.ConfigMapReference{
//				Name: "mimap",
//				Key:  "tenant",
//			},
//			AdministratorType: &adminType,
//			Login:
//		},
//	}
//
//	tc.CreateResourceAndWait(firewall)
//
//	tc.Expect(firewall.Status.Id).ToNot(BeNil())
//	armId := *firewall.Status.Id
//
//	tc.DeleteResourceAndWait(firewall)
//	exists, _, err := tc.AzureClient.HeadByID(
//		tc.Ctx,
//		armId,
//		string(sql.APIVersion_Value))
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}

// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-elastic-pool-create/azuredeploy.json
func SQL_Server_ElasticPool_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	elasticPool := &sql.ServersElasticPool{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.Servers_ElasticPool_Spec{
			Owner:    testcommon.AsOwner(server),
			Location: tc.AzureRegion,
			Sku: &sql.Sku{
				Name:     to.StringPtr("StandardPool"),
				Tier:     to.StringPtr("Standard"),
				Capacity: to.IntPtr(100),
			},
			PerDatabaseSettings: &sql.ElasticPoolPerDatabaseSettings{
				MinCapacity: to.Float64Ptr(0),
				MaxCapacity: to.Float64Ptr(100),
			},
		},
	}

	tc.CreateResourceAndWait(elasticPool)

	tc.Expect(elasticPool.Status.Id).ToNot(BeNil())
	armId := *elasticPool.Status.Id

	tc.DeleteResourceAndWait(elasticPool)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-with-failover-group/azuredeploy.json
//func SQL_Server_Failover_GroupCRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
//	elasticPool := &sql.ServersFailoverGroup{
//		ObjectMeta: tc.MakeObjectMeta("pool"),
//		Spec: sql.Servers_FailoverGroup_Spec{
//			Owner:    testcommon.AsOwner(server),
//			Location: tc.AzureRegion,
//			Sku: &sql.Sku{
//				Name:     to.StringPtr("StandardPool"),
//				Tier:     to.StringPtr("Standard"),
//				Capacity: to.IntPtr(100),
//			},
//			PerDatabaseSettings: &sql.ElasticPoolPerDatabaseSettings{
//				MinCapacity: to.Float64Ptr(0),
//				MaxCapacity: to.Float64Ptr(100),
//			},
//		},
//	}
//
//	tc.CreateResourceAndWait(elasticPool)
//
//	tc.Expect(elasticPool.Status.Id).ToNot(BeNil())
//	armId := *elasticPool.Status.Id
//
//	tc.DeleteResourceAndWait(elasticPool)
//	exists, _, err := tc.AzureClient.HeadByID(
//		tc.Ctx,
//		armId,
//		string(sql.APIVersion_Value))
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}

func makeSubnetForSQLServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *network.VirtualNetworksSubnet {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMeta("vn"),
		Spec: network.VirtualNetwork_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
	privateEndpointNetworkPolicyDisabled := network.SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Disabled
	privateLinkNetworkPolicyDisabled := network.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.StringPtr("10.0.0.0/24"),
			ServiceEndpoints: []network.ServiceEndpointPropertiesFormat{
				{
					Service: to.StringPtr("Microsoft.Sql"),
				},
			},
			PrivateEndpointNetworkPolicies:    &privateEndpointNetworkPolicyDisabled,
			PrivateLinkServiceNetworkPolicies: &privateLinkNetworkPolicyDisabled,
		},
	}
	tc.CreateResourcesAndWait(vnet, subnet)

	return subnet
}

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.web/asev2-appservice-sql-vpngw/azuredeploy.json
func SQL_Server_VirtualNetworkRule_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
	subnet := makeSubnetForSQLServer(tc, rg)

	vnetRule := &sql.ServersVirtualNetworkRule{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.Servers_VirtualNetworkRule_Spec{
			Owner:                         testcommon.AsOwner(server),
			VirtualNetworkSubnetReference: tc.MakeReferenceFromResource(subnet),
		},
	}

	tc.CreateResourceAndWait(vnetRule)

	tc.Expect(vnetRule.Status.Id).ToNot(BeNil())
	armId := *vnetRule.Status.Id

	tc.DeleteResourceAndWait(vnetRule)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func SQL_Database_VulnerabilityAssessment_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase, storageDetails vulnStorageAccountDetails) {
	// First set up pre-reqs:
	enabled := sql.DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled
	securityAlertPolicy := &sql.ServersDatabasesSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.Servers_Databases_SecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(securityAlertPolicy)

	tc.Expect(securityAlertPolicy.Status.Id).ToNot(BeNil())

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersDatabasesVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.Servers_Databases_VulnerabilityAssessment_Spec{
			Owner: testcommon.AsOwner(db),
			RecurringScans: &sql.VulnerabilityAssessmentRecurringScansProperties{
				IsEnabled: to.BoolPtr(false),
			},
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			// TODO: Make this easier to build a combined path in ASO itself
			StorageContainerPath: to.StringPtr(fmt.Sprintf("%s%s", blobEndpoint, storageDetails.container)),
		},
	}

	tc.CreateResourceAndWait(vulnerabilityAssessment)

	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(vulnerabilityAssessment)
	// TODO: It seems like delete of this resource isn't actually honored by the service
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-database-transparent-encryption-create/azuredeploy.json
func SQL_Databases_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Database_Spec{
			Owner:     testcommon.AsOwner(server),
			Location:  tc.AzureRegion,
			Collation: to.StringPtr("SQL_Latin1_General_CP1_CI_AS"),
		},
	}

	tc.CreateResourceAndWait(db)

	tc.Expect(db.Status.Id).ToNot(BeNil())
	armId := *db.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL Database BackupLongTermRetentionPolicy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_BackupLongTermRetention_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Vulnerability Assessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_VulnerabilityAssessment_CRUD(tc, db, storageDetails)
			},
		})

	// TODO: Update?

	tc.DeleteResourceAndWait(db)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func SQL_BackupLongTermRetention_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	policy := &sql.ServersDatabasesBackupLongTermRetentionPolicy{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Databases_BackupLongTermRetentionPolicy_Spec{
			Owner:           testcommon.AsOwner(db),
			WeeklyRetention: to.StringPtr("P30D"),
		},
	}

	tc.CreateResourceAndWait(policy)
	//defer tc.DeleteResourceAndWait(policy) // This resource doesn't support delete
	tc.Expect(policy.Status.Id).ToNot(BeNil())
}

type vulnStorageAccountDetails struct {
	container             string
	secretName            string
	blobEndpointSecretKey string
	keySecretKey          string
}

func makeStorageAccountForSQLVulnerabilityAssessment(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) vulnStorageAccountDetails {
	const vulnerabilityAssessmentsContainerName = "vulnerabilityassessments"
	const secretName = "storagesecret"

	acct := newStorageAccount(tc, rg)
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		Secrets: &storage.StorageAccountOperatorSecrets{
			BlobEndpoint: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "blobEndpoint",
			},
			Key1: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "key1",
			},
		},
	}
	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccounts_BlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(vulnerabilityAssessmentsContainerName),
		Spec: storage.StorageAccounts_BlobServices_Container_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourcesAndWait(acct, blobService, blobContainer)

	return vulnStorageAccountDetails{
		container:             vulnerabilityAssessmentsContainerName,
		secretName:            secretName,
		blobEndpointSecretKey: "blobEndpoint",
		keySecretKey:          "key1",
	}
}
