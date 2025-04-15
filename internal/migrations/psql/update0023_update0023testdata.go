package migration

import (
	"database/sql"
	"log/slog"
	"runtime"
)

func DatabaseUpdate_update0023_update0023testdata(db *sql.DB, getInfo bool) (string, string, string, error) {
	// WARNING!!!
	// Do not change the update if it has already been installed by other developers or in production.
	// To correct the database, create a new update and register it in the list of updates.

	fromUpdateId, toUpdateId := ParseNameFuncUpdate(runtime.Caller(0))
	description := "Add test data with new roles"
	if getInfo {
		return fromUpdateId, toUpdateId, description, nil
	}

	tx, err := db.Begin()
	if err != nil {
		slog.Error("Failed to begin transaction: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}
	// Вставка данных в permissions
	permissions := []struct {
		Name, Description string
	}{
		{"view", "Разрешение на просмотр"},
		{"edit", "Разрешение на редактирование"},
		{"create", "Разрешение на создание"},
		{"delete", "Разрешение на удаление"},
	}
	for _, p := range permissions {
		_, err := db.Exec(`INSERT INTO permissions (id, name, description) VALUES (gen_random_uuid(), $1, $2) ON CONFLICT (name) DO NOTHING`, p.Name, p.Description)
		if err != nil {
			return fromUpdateId, toUpdateId, description, err
		}
	}
	// Вставка данных в roles
	roles := []struct {
		Name, Description string
	}{
		{"guest", "Только просмотр"},
		{"player", "Просмотр и редактирование"},
		{"admin", "Просмотр, редактирование, создание и удаление"},
	}
	for _, r := range roles {
		_, err := db.Exec(`INSERT INTO roles (id, name, description) VALUES (gen_random_uuid(), $1, $2) ON CONFLICT (name) DO NOTHING`, r.Name, r.Description)
		if err != nil {
			return fromUpdateId, toUpdateId, description, err
		}
	}

	// Получаем id ролей и прав
	roleIDs := make(map[string]string)
	rows, err := db.Query(`SELECT id, name FROM roles`)
	if err != nil {
		return fromUpdateId, toUpdateId, description, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return fromUpdateId, toUpdateId, description, err
		}
		roleIDs[name] = id
	}

	permIDs := make(map[string]string)
	rows, err = db.Query(`SELECT id, name FROM permissions`)
	if err != nil {
		return fromUpdateId, toUpdateId, description, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			return fromUpdateId, toUpdateId, description, err
		}
		permIDs[name] = id
	}

	// Вставка данных в role_permissions
	rolePerms := []struct {
		Role, Perm string
	}{
		{"guest", "view"},
		{"player", "view"},
		{"player", "edit"},
		{"admin", "view"},
		{"admin", "edit"},
		{"admin", "create"},
		{"admin", "delete"},
	}
	for _, rp := range rolePerms {
		_, err := db.Exec(`INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, roleIDs[rp.Role], permIDs[rp.Perm])
		if err != nil {
			return fromUpdateId, toUpdateId, description, err
		}
	}

	_, err = db.Exec(`
		INSERT INTO user_roles (
			user_id, role_id
		) VALUES (
			(SELECT id FROM users WHERE user_name = 'admin'),
			(SELECT id FROM roles WHERE name = 'admin')
		)
	`)
	if err != nil {
		return fromUpdateId, toUpdateId, description, err
	}
	err = tx.Commit()
	if err != nil {
		slog.Error("Failed to commit transaction: " + err.Error())
		return fromUpdateId, toUpdateId, description, err
	}

	return fromUpdateId, toUpdateId, description, nil
}
