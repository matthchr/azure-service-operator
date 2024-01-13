/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package sql

import (
	"context"
	"database/sql"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	asosql "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	"github.com/Azure/azure-service-operator/v2/internal/identity"
	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	azuresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/azuresql"
)

const Scope = "https://database.windows.net/.default"

type aadUser struct {
	user               *asosql.User
	resourceResolver   *resolver.Resolver
	credentialProvider identity.CredentialProvider
	log                logr.Logger
}

var _ Connector = &aadUser{}

func (u *aadUser) CreateOrUpdate(ctx context.Context) error {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	u.log.V(Status).Info("Creating AzureSQL AAD user")

	username := u.username()
	err = azuresqlutil.CreateOrUpdateAADUser(ctx, db, username)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	err = azuresqlutil.ReconcileUserRoles(ctx, db, username, u.user.Spec.Roles)
	if err != nil {
		return errors.Wrap(err, "ensuring database roles")
	}

	u.log.V(Status).Info("Successfully reconciled SQLUser")
	return nil
}

func (u *aadUser) Delete(ctx context.Context) error {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return err
	}
	defer db.Close()

	// TODO: There's still probably some ways that this user can be deleted but that we don't detect (and
	// TODO: so might cause an error triggering the resource to get stuck).
	// TODO: Cases where the server is in the process of being
	// TODO: deleted (or all system tables have been wiped?) might also exist...
	username := u.username()
	err = azuresqlutil.DropUser(ctx, db, username)
	if err != nil {
		return err
	}

	return nil
}

func (u *aadUser) Exists(ctx context.Context) (bool, error) {
	db, err := u.connectToDB(ctx)
	if err != nil {
		return false, err
	}
	defer db.Close()

	username := u.username()
	exists, err := azuresqlutil.DoesUserExist(ctx, db, username)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u *aadUser) connectToDB(ctx context.Context) (*sql.DB, error) {
	details, err := getOwnerDetails(ctx, u.resourceResolver, u.user)
	if err != nil {
		return nil, err
	}

	if u.user.Spec.AADUser == nil {
		return nil, errors.Errorf("AAD User missing $.spec.aadUser field")
	}
	//adminUser := u.user.Spec.AADUser.ServerAdminUsername
	//if len(adminUser) == 0 {
	//	return nil, errors.Errorf("AAD User must specify $.spec.aadUser.serverAdminUsernamed")
	//}

	return connectToDBAAD(ctx, u.credentialProvider, u.log, u.user, details)
}

func (u *aadUser) username() string {
	return u.user.Spec.AzureName
}

func connectToDBAAD(
	ctx context.Context,
	credentialProvider identity.CredentialProvider,
	log logr.Logger,
	user *asosql.User,
	ownerDetails ownerDetails,
	// adminUser string,
) (*sql.DB, error) {

	tokenProvider := func() (string, error) {
		credential, err := credentialProvider.GetCredential(ctx, user)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get credential")
		}
		token, err := credential.TokenCredential().GetToken(ctx, policy.TokenRequestOptions{Scopes: []string{Scope}})
		if err != nil {
			return "", errors.Wrap(err, "failed to get token from credential")
		}
		log.V(Verbose).Info("Retrieved token for Azure SQL", "scope", Scope, "expires", token.ExpiresOn)

		return token.Token, nil
	}

	// Connect to the DB
	// TODO: How does this work without the username?
	db, err := azuresqlutil.ConnectToDBAAD(ctx, ownerDetails.fqdn, ownerDetails.database, azuresqlutil.ServerPort, tokenProvider)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to connect database. Server: %s, Database: %s, Port: %d",
			ownerDetails.fqdn,
			ownerDetails.database,
			azuresqlutil.ServerPort)
	}

	return db, nil
}
