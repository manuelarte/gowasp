package repositories

import (
	"context"
	"database/sql"
	"gowasp/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
}

var _ UserRepository = new(UserRepositoryDB)

type UserRepositoryDB struct {
	DB *sql.DB
}

func (u UserRepositoryDB) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users(created_at, updated_at, username, password) VALUES (current_timestamp, current_timestamp, ?, ?)"
	if _, err := u.DB.ExecContext(ctx, query, user.Username, user.Password); err != nil {
		return err
	}
	return nil
}
