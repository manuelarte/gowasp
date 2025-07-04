package repositories

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

var ErrUserNotFound = errors.New("user and password not found")

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username, password string) (models.User, error)
}

var _ UserRepository = new(UserRepositoryDB)

type UserRepositoryDB struct {
	DB *gorm.DB
}

func (u UserRepositoryDB) Create(ctx context.Context, user *models.User) error {
	if err := u.DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepositoryDB) Login(ctx context.Context, username, password string) (models.User, error) {
	// CWE-89: Improper Neutralization of Special Elements used in an SQL Command
	// ('SQL Injection') https://cwe.mitre.org/data/definitions/89.html
	query := fmt.Sprintf("SELECT id, username, password FROM users "+
		"WHERE username = '%s' AND PASSWORD = '%s';", username, password)

	row := u.DB.WithContext(ctx).Raw(query).Row()
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
