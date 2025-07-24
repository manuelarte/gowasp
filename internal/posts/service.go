package posts

import (
	"context"

	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
)

//nolint:iface // separate repository from service
type Service interface {
	GetAll(ctx context.Context, pagination *pagorminator.Pagination) ([]*models.Post, error)
	GetByID(ctx context.Context, id uint64) (models.Post, error)
}

var _ Service = new(serviceImpl)

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{repository: repository}
}

func (b serviceImpl) GetAll(ctx context.Context,
	pagination *pagorminator.Pagination,
) ([]*models.Post, error) {
	posts, err := b.repository.GetAll(ctx, pagination)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (b serviceImpl) GetByID(ctx context.Context, id uint64) (models.Post, error) {
	return b.repository.GetByID(ctx, id)
}
