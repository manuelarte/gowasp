package users

import (
	"context"
	"fmt"

	//#nosec G501
	"crypto/md5"
	"encoding/hex"

	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/models/gerrors"
)

//nolint:iface // repository does not need to have the same methods as service
type Service interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (models.User, error)
	Login(ctx context.Context, username, password string) (models.User, error)
}

var _ Service = new(serviceImpl)

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (u serviceImpl) Create(ctx context.Context, user *models.User) error {
	if err := isValidPassword(user.Password); err != nil {
		return err
	}

	user.Password = hashit(user.Password)
	if err := u.repository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u serviceImpl) GetByID(ctx context.Context, id uint) (models.User, error) {
	user, err := u.repository.GetByID(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("error retrieving user: %w", err)
	}

	return user, nil
}

func (u serviceImpl) Login(ctx context.Context, username, password string) (models.User, error) {
	return u.repository.Login(ctx, username, hashit(password))
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
		return gerrors.PasswordNotValidError{Message: "Password must be at least 4 characters"}
	}

	return nil
}
