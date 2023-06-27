package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/bearatol/lg"
)

func (r *Repository) GetUser(ctx context.Context, login string) (*mapping.User, error) {
	query := `
	SELECT * FROM ` + usersTableName + ` WHERE login = $1
	`
	user := &mapping.User{}
	err := r.db.GetContext(ctx, user, query, login)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("user is not found")
	}

	return user, err
}

func (r *Repository) CreateUser(ctx context.Context, user *mapping.UserAvailableFileds) error {
	lg.Warnf("%+v", user)
	query := `
	INSERT INTO ` + usersTableName + `
		(login, password, name, email) VALUES
			(:login, :password, :name, :email)
	`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

func (r *Repository) UpdateUser(ctx context.Context, login string, user *mapping.UserAvailableFileds) error {
	query := `
	UPDATE ` + usersTableName + `
	SET
		login = :login, password = :password, name = :name, email = :email, updated_at = DEFAULT
	WHERE login = '` + login + `'
	`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}
func (r *Repository) DeleteUser(ctx context.Context, login string) error {
	query := `
	DELETE FROM ` + usersTableName + `
	WHERE login = $1
	`
	_, err := r.db.ExecContext(ctx, query, login)
	return err
}
