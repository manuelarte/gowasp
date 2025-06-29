package services

import (
	"context"
	//#nosec G501
	"crypto/md5"
	"encoding/hex"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/models/errors"
	"github.com/manuelarte/gowasp/internal/repositories"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	LoginUser(ctx context.Context, username, password string) (models.User, error)
}

var _ UserService = new(UserServiceImpl)

type UserServiceImpl struct {
	Repository repositories.UserRepository
}

func (u UserServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	if err := isValidPassword(user.Password); err != nil {
		return err
	}
	user.Password = hashit(user.Password)
	if err := u.Repository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u UserServiceImpl) LoginUser(ctx context.Context, username, password string) (models.User, error) {
	hashedPassword := hashit(password)
	return u.Repository.Login(ctx, username, hashedPassword)
}

// CWE-328: Use of Weak Hash https://cwe.mitre.org/data/definitions/328.html
func hashit(str string) string {
	//#nosec G401
	hash := md5.Sum([]byte(str))

	return hex.EncodeToString(hash[:])
}

// CWE-521: Weak Password Requirements https://cwe.mitre.org/data/definitions/521.html
func isValidPassword(password string) error {
	minPasswordLength := 4
	if len(password) < minPasswordLength {
		return errors.PasswordNotValidError{Message: "Password must be at least 4 characters"}
	}

	return nil
}
