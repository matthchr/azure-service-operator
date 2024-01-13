/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"os"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	managedidentity2022 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20220131preview"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	azuresqlv1 "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	azuresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/azuresql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AzureSQL_AAD_User(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePasswordOfLength(60) // Use a long password to ensure we meet complexity requirements

	secret := newSecret(tc, adminPasswordKey, adminPassword)
	tc.CreateResource(secret)

	server := newAzureSQLServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
	database := newAzureSQLServerDatabase(tc, server)
	firewallRule := newSQLServerOpenFirewallRule(tc, server)

	tc.CreateResourcesAndWait(server, database, firewallRule)

	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *server.Status.FullyQualifiedDomainName

	// Ensure that firewall rule access has worked. It can take up to 5 minutes to take effect
	tc.G.Eventually(
		func() error {
			db, err := azuresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				database.AzureName(),
				azuresqlutil.ServerPort,
				adminUsername,
				adminPassword)
			if err != nil {
				return err
			}

			return db.Close()
		},
		5*time.Minute,
	).Should(Succeed())

	// These must run sequentially as they're mutating SQL state
	// TODO: Ought to test the AzureSQL User via AAD Admin too?
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "AzureSQL Administrator CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				SQL_Server_Administrator_CRUD(testContext, server, rg)
			},
		},
		testcommon.Subtest{
			Name: "AzureSQL AAD User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_AAD_User_CRUD(testContext, rg, server, database, adminPassword)
			},
		},
	)
}

// TODO: read env for either principalID or OIDC URL, or from OIDC file... which maybe seems best?
// TODO: Move this to TC for sure
func readOIDCIssuer() (string, error) {
	issuer, err := os.ReadFile("../out/workload-identity/azure/saissuer.txt")
	if err != nil {
		return "", errors.Wrapf(err, "failed to read OIDC issuer file. Is the kind cluster using workload identity?")
	}
	return strings.TrimSpace(string(issuer)), nil
}

func SQL_Server_Administrator_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
	issuer, err := readOIDCIssuer()
	tc.Expect(err).ToNot(HaveOccurred())

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: "mimap",
						Key:  "oid",
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: "mimap",
						Key:  "tenant",
					},
				},
			},
		},
	}
	tc.CreateResourceAndWait(mi)

	fic := &managedidentity2022.FederatedIdentityCredential{
		ObjectMeta: tc.MakeObjectMeta("fic"),
		Spec: managedidentity2022.UserAssignedIdentities_FederatedIdentityCredential_Spec{
			Owner: testcommon.AsOwner(mi),
			// For Workload Identity, Audiences should always be "api://AzureADTokenExchange"
			Audiences: []string{
				"api://AzureADTokenExchange",
			},
			Issuer: to.Ptr(issuer),
			// For Workload Identity, Subject should always be system:serviceaccount:<namespace>:<serviceaccount>
			Subject: to.Ptr("system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default"),
		},
	}
	tc.CreateResourceAndWait(fic)

	adminType := sql.AdministratorProperties_AdministratorType_ActiveDirectory
	admin := &sql.ServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("admin"),
		Spec: sql.Servers_Administrator_Spec{
			Owner: testcommon.AsOwner(server),
			SidFromConfig: &genruntime.ConfigMapReference{
				Name: "mimap",
				Key:  "oid",
			},
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: "mimap",
				Key:  "tenant",
			},
			AdministratorType: &adminType,
			Login:             to.Ptr(mi.Name),
		},
	}

	tc.CreateResourceAndWait(admin)
}

func AzureSQL_AAD_User_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, server *sql.Server, database *sql.ServersDatabase, adminPassword string) {
	issuer, err := readOIDCIssuer()
	tc.Expect(err).ToNot(HaveOccurred())

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: "misecret",
						Key:  "AZURE_PRINCIPAL_ID",
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: "misecret",
						Key:  "AZURE_TENANT_ID",
					},
					ClientId: &genruntime.ConfigMapDestination{
						Name: "misecret",
						Key:  "AZURE_CLIENT_ID",
					},
				},
			},
		},
	}
	tc.CreateResourceAndWait(mi)

	// TODO: The fact you can't get subscription ID into a secret easily like this is problematic, as it makes it hard to use
	// TODO: directly via credential-from?
	credentialFromSecret := "misecret-with-subscription"
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMetaWithName(credentialFromSecret),
		StringData: map[string]string{
			"AZURE_CLIENT_ID":       to.Value(mi.Status.ClientId),
			"AZURE_TENANT_ID":       to.Value(mi.Status.TenantId),
			"AZURE_SUBSCRIPTION_ID": tc.AzureSubscription,
		},
	}
	tc.CreateResource(secret)

	fic := &managedidentity2022.FederatedIdentityCredential{
		ObjectMeta: tc.MakeObjectMeta("fic"),
		Spec: managedidentity2022.UserAssignedIdentities_FederatedIdentityCredential_Spec{
			Owner: testcommon.AsOwner(mi),
			// For Workload Identity, Audiences should always be "api://AzureADTokenExchange"
			Audiences: []string{
				"api://AzureADTokenExchange",
			},
			Issuer: to.Ptr(issuer),
			// For Workload Identity, Subject should always be system:serviceaccount:<namespace>:<serviceaccount>
			Subject: to.Ptr("system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default"),
		},
	}
	tc.CreateResourceAndWait(fic)

	adminType := sql.AdministratorProperties_AdministratorType_ActiveDirectory
	admin := &sql.ServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("admin"),
		Spec: sql.Servers_Administrator_Spec{
			Owner: testcommon.AsOwner(server),
			SidFromConfig: &genruntime.ConfigMapReference{
				Name: "misecret",
				Key:  "AZURE_PRINCIPAL_ID",
			},
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: "misecret",
				Key:  "AZURE_TENANT_ID",
			},
			AdministratorType: &adminType,
			Login:             to.Ptr(mi.Name),
		},
	}
	tc.CreateResourceAndWait(admin)

	user := &azuresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithNameAndCredentialFrom(mi.Name, credentialFromSecret),
		Spec: azuresqlv1.UserSpec{
			Owner: testcommon.AsOwner(database),
			Roles: []string{
				"db_datareader",
			},
			// TODO: Weird
			AADUser: &azuresqlv1.AADUserSpec{},
		},
	}
	tc.CreateResourcesAndWait(user)

	// Connect to the DB
	fqdn := to.Value(server.Status.FullyQualifiedDomainName)
	conn, err := azuresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		database.AzureName(),
		azuresqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right roles
	roles, err := azuresqlutil.GetUserRoles(tc.Ctx, conn, mi.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Update the user
	old := user.DeepCopy()
	user.Spec.Roles = []string{
		"db_datareader",
		"db_datawriter",
		"db_securityadmin",
	}
	tc.PatchResourceAndWait(old, user)

	// Confirm that we have the right roles
	roles, err = azuresqlutil.GetUserRoles(tc.Ctx, conn, mi.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Unfortunately we can't easily test that we can actually connect as this user

	originalUser := user.DeepCopy()

	// Confirm that we cannot change the user type from AAD to local
	user = originalUser.DeepCopy()
	old = user.DeepCopy()
	user.Spec.AADUser = nil
	user.Spec.LocalUser = &azuresqlv1.LocalUserSpec{
		ServerAdminUsername: "someadminuser",
	}
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change from AAD User to local user"))

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)

	conn, err = azuresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		database.AzureName(),
		azuresqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	exists, err := azuresqlutil.DoesUserExist(tc.Ctx, conn, user.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
