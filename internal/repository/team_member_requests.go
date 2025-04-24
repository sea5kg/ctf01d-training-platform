package repository

import (
	"context"
	"database/sql"
	"fmt"

	"ctf01d/internal/model"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	StatusPending  = "pending"
	StatusApproved = "approved"
	StatusRejected = "rejected"
)

type TeamMemberRequestRepository interface {
	ConnectUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID, role string) error
	ApproveUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID) error
	RejectUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID) error
	LeaveUserFromTeam(ctx context.Context, teamID, userID openapi_types.UUID) error
	TeamMembers(ctx context.Context, teamID openapi_types.UUID) ([]*model.User, error)
}

func NewTeamMemberRequestRepository(db *sql.DB) TeamMemberRequestRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) ConnectUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID, role string) error {
	query := `INSERT INTO team_member_requests (team_id, user_id, role, status)
	          VALUES ($1, $2, $3, 'pending')`
	_, err := r.db.ExecContext(ctx, query, teamID, userID, role)
	return err
}

func (r *teamRepo) ApproveUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `
		UPDATE team_member_requests SET status = 'approved' WHERE team_id = $1 AND user_id = $2 AND status = $3
	`
	_, err = tx.ExecContext(ctx, query, teamID, userID, StatusPending)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	var role string
	query = `
		SELECT role FROM team_member_requests WHERE team_id = $1 AND user_id = $2 AND status = $3
	`
	err = tx.QueryRowContext(ctx, query, teamID, userID, StatusApproved).Scan(&role)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}
	// fixme - обновить team_history
	query = `
		INSERT INTO profiles (current_team_id, user_id, role) VALUES ($1, $2, $3)
	`
	_, err = tx.ExecContext(ctx, query, teamID, userID, role)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (r *teamRepo) RejectUserTeamRequest(ctx context.Context, teamID, userID openapi_types.UUID) error {
	query := `
		UPDATE team_member_requests
		SET status = $3
		WHERE team_id = $1 AND user_id = $2 AND status = $4
	`
	result, err := r.db.ExecContext(ctx, query, teamID, userID, StatusRejected, StatusPending)
	if err != nil {
		return fmt.Errorf("failed to execute reject query for user %s in team %s: %w", userID, teamID, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected after rejecting user %s in team %s: %w", userID, teamID, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no pending request found to reject for user %s in team %s", userID, teamID)
	}
	return nil
}

func (r *teamRepo) LeaveUserFromTeam(ctx context.Context, teamID, userID openapi_types.UUID) error {
	// fixme - обновить team_history
	query := `
		DELETE FROM profiles WHERE current_team_id = $1 AND user_id = $2
	`
	_, err := r.db.ExecContext(ctx, query, teamID, userID)
	return err
}

func (r *teamRepo) TeamMembers(ctx context.Context, teamID openapi_types.UUID) ([]*model.User, error) {
	query := `
		SELECT u.id, u.display_name, u.user_name, tm.role, u.avatar_url, u.status
		FROM profiles tm
		JOIN users u ON tm.user_id = u.id
		WHERE tm.current_team_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*model.User
	for rows.Next() {
		var member model.User
		if err := rows.Scan(&member.Id, &member.DisplayName, &member.Username, &member.Role, &member.AvatarUrl, &member.Status); err != nil {
			return nil, err
		}
		members = append(members, &member)
	}
	return members, nil
}
