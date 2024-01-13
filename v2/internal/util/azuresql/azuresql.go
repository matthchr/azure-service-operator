/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package azuresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/pkg/errors"
)

// ServerPort is the default server port for sql server
const ServerPort = 1433

// driverName is driver name for db connection
const driverName = "sqlserver"

func ConnectToDB(
	ctx context.Context,
	serverAddress string,
	database string,
	port int,
	user string,
	password string,
) (*sql.DB, error) {
	// Make sure to set connection timeout: https://github.com/denisenkom/go-mssqldb/issues/609
	connString := fmt.Sprintf(
		"server=%s;database=%s;user id=%s;password=%s;port=%d;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30",
		serverAddress,
		database,
		user,
		password,
		port)

	db, err := sql.Open(driverName, connString)
	if err != nil {
		return db, err
	}

	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}

// ConnectToDBAAD connects to the Azure SQL database using the specified token provider.
// TODO: Does this work for any form of user? Don't need a username at all?
func ConnectToDBAAD(
	ctx context.Context,
	serverAddress string,
	database string,
	port int,
	tokenProvider func() (string, error),
) (*sql.DB, error) {
	// Make sure to set connection timeout: https://github.com/denisenkom/go-mssqldb/issues/609
	connString := fmt.Sprintf("server=%s;database=%s;port=%d;Connection Timeout=30", serverAddress, database, port)

	connector, err := mssql.NewAccessTokenConnector(connString, tokenProvider)
	if err != nil {
		return nil, errors.Wrap(err, "NewAccessTokenConnector failed")
	}

	db := sql.OpenDB(connector)

	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrap(err, "PingContext failed")
	}

	return db, err
}

// TODO: I don't think hostname is a thing
//func HostnameOrDefault(hostname string) string {
//	if hostname == "" {
//		hostname = "%"
//	}
//
//	return hostname
//}

func CreateOrUpdateUser(ctx context.Context, db *sql.DB, username string, password string) error {
	// make an effort to prevent sql injection
	if err := findBadChars(username); err != nil {
		return errors.Wrap(err, "problem found with username")
	}
	if err := findBadChars(password); err != nil {
		return errors.Wrap(err, "problem found with password")
	}

	tsql := `
IF NOT EXISTS (SELECT name FROM sysusers WHERE name='%[1]s')
	BEGIN
		CREATE USER "%[1]s" WITH PASSWORD='%[2]s';
	END
ELSE
	BEGIN
		ALTER USER "%[1]s" WITH PASSWORD='%[2]s';
	END;
`
	tsql = fmt.Sprintf(tsql, username, password)
	_, err := db.ExecContext(ctx, tsql)
	if err != nil {
		return err
	}

	return nil
}

// CreateOrUpdateAADUser creates or updates an AAD user.
// See https://learn.microsoft.com/en-us/azure/azure-sql/database/authentication-aad-configure for details on how to create AAD
// users.
// username can be either the actual AAD username (for real AAD users), the group name for groups, or
// the managed identity name (app) for managed identities.
func CreateOrUpdateAADUser(ctx context.Context, db *sql.DB, username string) error {

	//"CREATE USER [Azure_AD_user_name] FROM EXTERNAL PROVIDER;"
	//"CREATE USER [bob@contoso.com] FROM EXTERNAL PROVIDER;"
	//"CREATE USER [alice@fabrikam.onmicrosoft.com] FROM EXTERNAL PROVIDER;"

	//"CREATE USER [appName] FROM EXTERNAL PROVIDER;" -- TODO: This seems to work for UMI... how can I test logging in?

	if err := findBadChars(username); err != nil {
		return errors.Wrap(err, "problem found with managed identity username")
	}

	// TODO: There doesn't seem to be a need to update (and the FROM EXTERNAL PROVIDER syntax isn't valid for ALTER anyway).
	tsql := `
IF NOT EXISTS (SELECT name FROM sysusers WHERE name='%[1]s')
	BEGIN
		CREATE USER [%[1]s] FROM EXTERNAL PROVIDER;
	END
`

	tsql = fmt.Sprintf(tsql, username)
	_, err := db.ExecContext(ctx, tsql)
	if err != nil {
		return err
	}

	return nil
}

// GrantUserRoles grants roles to a user for a given database
// TODO: Need to make this merge/delete existing roles
//func GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
//	var errs []error
//	for _, role := range roles {
//		// tsql := "sp_addrolemember @role, @user"
//		tsql := "ALTER ROLE @role ADD MEMBER @user"
//
//		_, err := db.ExecContext(
//			ctx, tsql,
//			sql.Named("role", role),
//			sql.Named("user", user),
//		)
//		if err != nil {
//			errs = append(errs)
//		}
//	}
//
//	return kerrors.NewAggregate(errs)
//}

// DoesUserExist checks if db contains user
func DoesUserExist(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(
		ctx,
		"SELECT * FROM sysusers WHERE NAME=@user",
		sql.Named("user", username),
	)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// DropUser drops a user from db
func DropUser(ctx context.Context, db *sql.DB, username string) error {
	if err := findBadChars(username); err != nil {
		return errors.Wrap(err, "Problem found with username")
	}
	tsql := fmt.Sprintf("DROP USER [%s]", username)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// TODO: This is horrible we should see if we can do something better
//func convertToSid(msiClientId string) string {
//	var byteguid string
//	guid, err := uuid.FromString(msiClientId)
//	if err != nil {
//		return ""
//	}
//	bytes := guid.Bytes()
//
//	// Encode first 3 sets of bytes in little endian format
//	for i := 3; i >= 0; i-- {
//		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
//	}
//	for i := 5; i > 3; i-- {
//		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
//	}
//	for i := 7; i > 5; i-- {
//		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
//	}
//
//	// Encode last 2 sets of bytes in big endian format
//	for i := 8; i < len(bytes); i++ {
//		byteguid = byteguid + fmt.Sprintf("%02X", bytes[i])
//	}
//	return "0x" + byteguid
//}

func findBadChars(stack string) error {
	badChars := []string{
		"'",
		"\"",
		";",
		"--",
		"/*",
	}

	for _, s := range badChars {
		if idx := strings.Index(stack, s); idx > -1 {
			return fmt.Errorf("potentially dangerous character seqience found: '%s' at pos: %d", s, idx)
		}
	}
	return nil
}
