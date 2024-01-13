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

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pkg/errors"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

type SQLRoleDelta struct {
	AddedRoles   set.Set[string]
	DeletedRoles set.Set[string]
}

func DiffCurrentAndExpectedSQLRoles(currentRoles set.Set[string], expectedRoles set.Set[string]) SQLRoleDelta {
	result := SQLRoleDelta{
		AddedRoles:   set.Make[string](),
		DeletedRoles: set.Make[string](),
	}

	for role := range expectedRoles {
		// If an expected role isn't in the current role set, we need to add it
		if !currentRoles.Contains(role) {
			result.AddedRoles.Add(role)
		}
	}

	for role := range currentRoles {
		// If a current role isn't in the expected set, we need to remove it
		if !expectedRoles.Contains(role) {
			result.DeletedRoles.Add(role)
		}
	}

	return result
}

//
//// IsSQLAll returns whether the string matches the special privilege value ALL.
//func IsSQLAll(privilege string) bool {
//	return strings.EqualFold(privilege, sqlAll)
//}

// GetUserDatabaseRoles gets the per-database privileges that the
// user has. The user can have different permissions to each
// database. The details of access are returned in the map, keyed by
// database name.
//func GetUserDatabaseRoles(ctx context.Context, db *sql.DB, user string) (map[string]set.Set[string], error) {
//	//hostname = HostnameOrDefault(hostname)
//
//	tsql := `
//SELECT r.name role_principal_name
//FROM sys.database_role_members rm
//JOIN sys.database_principals r
//    ON rm.role_principal_id = r.principal_id
//JOIN sys.database_principals m
//    ON rm.member_principal_id = m.principal_id
//where m.name = @user
//`
//
//	// Note: This works because we only assign permissions at the DB level, not at the table, column, etc levels -- if we assigned
//	// permissions at more levels we would need to do something else here such as join multiple tables or
//	// parse SHOW GRANTS with a regex.
//	rows, err := db.QueryContext(
//		ctx,
//		"SELECT TABLE_SCHEMA, PRIVILEGE_TYPE FROM INFORMATION_SCHEMA.SCHEMA_PRIVILEGES WHERE GRANTEE = ?",
//		user, hostname)
//	if err != nil {
//		return nil, errors.Wrapf(err, "listing database grants for user %s", user)
//	}
//	defer rows.Close()
//
//	results := make(map[string]set.Set[string])
//	for rows.Next() {
//		var database, privilege string
//		err := rows.Scan(&database, &privilege)
//		if err != nil {
//			return nil, errors.Wrapf(err, "extracting privilege row")
//		}
//
//		var privileges set.Set[string]
//		if existingPrivileges, found := results[database]; found {
//			privileges = existingPrivileges
//		} else {
//			privileges = make(set.Set[string])
//			results[database] = privileges
//		}
//		privileges.Add(privilege)
//	}
//
//	if rows.Err() != nil {
//		return nil, errors.Wrapf(rows.Err(), "iterating database privileges")
//	}
//
//	return results, nil
//}

// GetUserRoles gets the roles assigned t othe user
func GetUserRoles(ctx context.Context, db *sql.DB, user string) (set.Set[string], error) {
	tsql := `
SELECT r.name role_principal_name
FROM sys.database_role_members rm 
JOIN sys.database_principals r 
    ON rm.role_principal_id = r.principal_id
JOIN sys.database_principals m 
    ON rm.member_principal_id = m.principal_id
where m.name = @user
`

	rows, err := db.QueryContext(
		ctx,
		tsql,
		sql.Named("user", user))
	if err != nil {
		return nil, errors.Wrapf(err, "listing roles for user %s", user)
	}
	defer rows.Close()

	result := make(set.Set[string])
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return nil, errors.Wrapf(err, "extracting role field")
		}

		result.Add(row)
	}
	if rows.Err() != nil {
		return nil, errors.Wrapf(rows.Err(), "iterating roles")
	}

	return result, nil
}

// ReconcileUserRoles revokes and grants roles as
// needed so the roles for the user match those passed in.
func ReconcileUserRoles(ctx context.Context, db *sql.DB, user string, roles []string) error {
	var errs []error
	desiredRoles := set.Make[string](roles...)

	currentRoles, err := GetUserRoles(ctx, db, user)
	if err != nil {
		return errors.Wrapf(err, "couldn't get existing roles for user %s", user)
	}

	rolesDiff := DiffCurrentAndExpectedSQLRoles(currentRoles, desiredRoles)
	err = addRoles(ctx, db, user, rolesDiff.AddedRoles)
	if err != nil {
		errs = append(errs, err)
	}
	err = deleteRoles(ctx, db, user, rolesDiff.DeletedRoles)
	if err != nil {
		errs = append(errs, err)
	}

	err = kerrors.NewAggregate(errs)
	if err != nil {
		return err
	}
	return nil
}

// TODO: Fix privs comment
// ReconcileUserDatabaseRoles revokes and grants database privileges as needed
// so they match the ones passed in. If there's an error applying
// privileges for one database it will still continue to apply
// privileges for subsequent databases (before reporting all errors).
//func ReconcileUserDatabaseRoles(ctx context.Context, conn *sql.DB, user string, dbRoles map[string][]string) error {
//	desiredRoles := make(map[string]set.Set[string])
//	for database, roles := range dbRoles {
//		desiredRoles[database] = set.Make[string](roles...)
//	}
//
//	currentPrivs, err := GetUserDatabaseRoles(ctx, conn, user, hostname)
//	if err != nil {
//		return errors.Wrapf(err, "couldn't get existing database privileges for user %s", user)
//	}
//
//	allDatabases := make(set.Set[string])
//	for db := range desiredRoles {
//		allDatabases.Add(db)
//	}
//	for db := range currentPrivs {
//		allDatabases.Add(db)
//	}
//
//	var dbErrors []error
//	for db := range allDatabases {
//		privsDiff := DiffCurrentAndExpectedSQLRoles(
//			currentPrivs[db],
//			desiredRoles[db],
//		)
//
//		err = addRoles(ctx, conn, db, user, privsDiff.AddedRoles)
//		if err != nil {
//			dbErrors = append(dbErrors, errors.Wrap(err, db))
//		}
//		err = deleteRoles(ctx, conn, db, user, privsDiff.DeletedRoles)
//		if err != nil {
//			dbErrors = append(dbErrors, errors.Wrap(err, db))
//		}
//	}
//
//	return kerrors.NewAggregate(dbErrors)
//}

func addRoles(ctx context.Context, db *sql.DB, user string, roles set.Set[string]) error {
	return alterRoles(ctx, db, user, roles, addMember)
}

func deleteRoles(ctx context.Context, db *sql.DB, user string, roles set.Set[string]) error {
	return alterRoles(ctx, db, user, roles, dropMember)
}

type alterRoleMode string

const (
	addMember  = alterRoleMode("ADD")
	dropMember = alterRoleMode("DROP")
)

func alterRoles(ctx context.Context, db *sql.DB, user string, roles set.Set[string], mode alterRoleMode) error {
	if len(roles) == 0 {
		// Nothing to do
		return nil
	}

	builder := strings.Builder{}
	for role := range roles {
		// TODO: Either need escaping here, or we need to use sql.Named?
		_, err := builder.WriteString(fmt.Sprintf("ALTER ROLE %s %s MEMBER %s;\n", role, string(mode), user))
		if err != nil {
			return errors.Wrapf(err, "failed to build T-SQL ALTER ROLE %s statement", mode)
		}
	}

	_, err := db.ExecContext(
		ctx,
		builder.String(),
	)
	if err != nil {
		return errors.Wrapf(err, "failed to execute T-SQL ALTER ROLE %s statement", mode)
	}

	//for role := range roles {
	//	// TODO: Either need escaping here, or we need to use sql.Named?
	//	tsql := fmt.Sprintf("ALTER ROLE %s %s MEMBER @user;", role, string(mode))
	//
	//	_, err := db.ExecContext(
	//		ctx,
	//		tsql,
	//		sql.Named("user", user),
	//	)
	//	if err != nil {
	//		return errors.Wrapf(err, "failed to execute T-SQL ALTER ROLE %s statement", mode)
	//	}
	//}

	return nil
}
