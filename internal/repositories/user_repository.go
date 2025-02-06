package repositories

import (
	"context"
	"gorm.io/gorm"
	"gowasp/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	Login(ctx context.Context, user *models.User) error
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

func (u UserRepositoryDB) Login(ctx context.Context, user *models.User) error {
	if err := u.DB.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}
