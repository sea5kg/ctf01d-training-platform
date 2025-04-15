package casbinrepo

import (
	"database/sql"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

const CasbinModel = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _  # user, role, domain

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (p.sub == "*" || g(r.sub, p.sub, r.dom)) && (p.dom == "*" || r.dom == p.dom) && keyMatch(r.obj, p.obj) && (p.act == "*" || regexMatch(r.act, p.act))
`

func NewEnforcer() (*casbin.Enforcer, error) {
	m, err := model.NewModelFromString(CasbinModel)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Casbin model: %w", err)
	}

	enforcer, err := casbin.NewEnforcer(m)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Casbin enforcer: %w", err)
	}

	return enforcer, nil
}

func LoadPolicies(enforcer *casbin.Enforcer, db *sql.DB) error {
	// "subject" "domain" "object" "action"
	// teams
	if _, err := enforcer.AddPolicy("*", "global", "teams", "create"); err != nil {
		return fmt.Errorf("failed to add policy for creating teams: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "global", "teams", "view"); err != nil {
		return fmt.Errorf("failed to add policy for viewing teams: %w", err)
	}

	// users
	if _, err := enforcer.AddPolicy("*", "global", "users", "create"); err != nil {
		return fmt.Errorf("failed to add policy for creating users: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "global", "users", "view"); err != nil {
		return fmt.Errorf("failed to add policy for viewing users: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "global", "users", "edit"); err != nil {
		return fmt.Errorf("failed to add policy for editing users: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "global", "users", "delete"); err != nil {
		return fmt.Errorf("failed to add policy for deleting users: %w", err)
	}

	// roles
	if _, err := enforcer.AddPolicy("*", "*", "roles", "view"); err != nil {
		return fmt.Errorf("failed to add policy for viewing roles: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "*", "permissions", "view"); err != nil {
		return fmt.Errorf("failed to add policy for viewing permissions: %w", err)
	}

	// auth
	if _, err := enforcer.AddPolicy("*", "*", "logout", "create"); err != nil {
		return fmt.Errorf("failed to add policy for creating logout: %w", err)
	}
	if _, err := enforcer.AddPolicy("*", "*", "pong", "view"); err != nil {
		return fmt.Errorf("failed to add policy for viewing pong: %w", err)
	}

	rolePermissionsQuery := `
	    SELECT r.name AS role_name, p.name AS permission_name
        FROM role_permissions rp
        JOIN roles r ON rp.role_id = r.id
        JOIN permissions p ON rp.permission_id = p.id
    `
	rolePermissions := make(map[string][]string) // role_name -> []permission_name
	rows, err := db.Query(rolePermissionsQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var roleName, permissionName string
		if err := rows.Scan(&roleName, &permissionName); err != nil {
			return err
		}

		rolePermissions[roleName] = append(rolePermissions[roleName], permissionName)
	}

	userRolesQuery := `
		SELECT ur.user_id, r.name AS role_name, ur.team_id
		FROM user_roles ur
		JOIN roles r ON ur.role_id = r.id
	`
	roleDomains := make(map[string]map[string]bool) // role_name -> set of domains
	userRoles, err := db.Query(userRolesQuery)
	if err != nil {
		return err
	}
	defer userRoles.Close()

	for userRoles.Next() {
		var userID string
		var roleName string
		var teamID sql.NullString
		if err := userRoles.Scan(&userID, &roleName, &teamID); err != nil {
			return err
		}

		var domain string
		if teamID.Valid {
			domain = fmt.Sprintf("team_%s", teamID.String)
		} else {
			domain = "global"
		}

		subject := fmt.Sprintf("user_%s", userID)
		if _, err := enforcer.AddGroupingPolicy(subject, roleName, domain); err != nil {
			return fmt.Errorf("failed to add grouping policy for subject %s, role %s, domain %s: %w", subject, roleName, domain, err)
		}

		if _, ok := roleDomains[roleName]; !ok {
			roleDomains[roleName] = make(map[string]bool)
		}
		roleDomains[roleName][domain] = true
	}

	for subject, actions := range rolePermissions {
		domains, ok := roleDomains[subject]
		if !ok {
			continue
		}

		for domain := range domains {
			for _, action := range actions {
				objects := getObjectsForPermission(action)
				for _, object := range objects {
					if _, err := enforcer.AddPolicy(subject, domain, object, action); err != nil {
						return fmt.Errorf("failed to add policy for role %s, domain %s, object %s, permission %s: %w", subject, domain, object, action, err)
					}
				}
			}
		}
	}
	return nil
}

func getObjectsForPermission(permission string) []string {
	switch permission {
	case "create":
		return []string{"teams"}
	case "view":
		return []string{"teams"}
	case "edit":
		return []string{"teams"}
	case "delete":
		return []string{"teams"}
	default:
		return []string{}
	}
}
