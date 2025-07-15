package users

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

var ErrUserNotFound = errors.New("user and password not found")

//nolint:iface // repository does not need to have the same methods as service
type Repository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username, password string) (models.User, error)
}

var _ Repository = new(gormRepository)

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (u gormRepository) Create(ctx context.Context, user *models.User) error {
	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u gormRepository) Login(ctx context.Context, username, password string) (models.User, error) {
	// CWE-89: Improper Neutralization of Special Elements used in an SQL Command
	// ('SQL Injection') https://cwe.mitre.org/data/definitions/89.html
	query := fmt.Sprintf("SELECT id, username, password FROM users "+
		"WHERE username = '%s' AND PASSWORD = '%s';", username, password)

	row := u.db.WithContext(ctx).Raw(query).Row()
	if row.Err() != nil {
		return models.User{}, row.Err()
	}
	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return models.User{}, ErrUserNotFound
	}

	return user, nil
}
