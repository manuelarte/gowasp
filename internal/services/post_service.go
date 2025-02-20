package services

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/repositories"
)

type PostService interface {
	GetAll(ctx context.Context, pagination *pagorminator.Pagination) (models.PageResponse[*models.Post], error)
	GetById(ctx context.Context, id int) (models.Post, error)
}

var _ PostService = new(PostServiceImpl)

type PostServiceImpl struct {
	Repository repositories.PostRepository
}

func (b PostServiceImpl) GetAll(ctx context.Context, pagination *pagorminator.Pagination) (models.PageResponse[*models.Post], error) {
	posts, err := b.Repository.GetAll(ctx, pagination)
	if err != nil {
		return models.PageResponse[*models.Post]{}, err
	}
	return models.PageResponse[*models.Post]{
		Data: posts,
		Metadata: models.PageMetadata{
			Page:       pagination.GetPage(),
			Size:       pagination.GetSize(),
			TotalCount: pagination.GetTotalElements(),
			TotalPages: pagination.GetTotalPages(),
		},
	}, nil
}

func (b PostServiceImpl) GetById(ctx context.Context, id int) (models.Post, error) {
	return b.Repository.GetById(ctx, id)
}
