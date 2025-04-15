package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0022_update0023(db *sql.DB, getInfo bool) (string, string, string, error) {
	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Create permissions, roles, user_roles, role_permissions, organization_keys tables (UUID ids)"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}

	query := `
		BEGIN;
		CREATE EXTENSION IF NOT EXISTS "pgcrypto";
		CREATE TABLE IF NOT EXISTS permissions (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255) NOT NULL UNIQUE,
			description TEXT
		);
		CREATE TABLE IF NOT EXISTS roles (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255) NOT NULL UNIQUE,
			description TEXT
		);
		CREATE TABLE IF NOT EXISTS user_roles (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			team_id UUID NULL REFERENCES teams(id),
			user_id UUID NOT NULL REFERENCES users(id),
			role_id UUID NOT NULL REFERENCES roles(id)
		);
		CREATE TABLE IF NOT EXISTS role_permissions (
			role_id UUID NOT NULL REFERENCES roles(id),
			permission_id UUID NOT NULL REFERENCES permissions(id),
			PRIMARY KEY (role_id, permission_id)
		);
		COMMIT;
	`
	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Problem with exec, query: " + query + "\n   error:" + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	return fromUpdateId, toUpdateId, description, nil
}
