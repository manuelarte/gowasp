package repositories

import (
	"context"
	"gorm.io/gorm"
	"gowasp/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email string, password string) (models.User, error)
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

func (u UserRepositoryDB) Login(ctx context.Context, email string, password string) (models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = '" + email + "' AND password = '" + password + "';"

	row := u.DB.WithContext(ctx).Raw(query).Row()
	if row.Err() != nil {
		return models.User{}, row.Err()
	}
	var user models.User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
