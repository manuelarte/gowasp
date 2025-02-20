package services

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/repositories"
)

type BlogService interface {
	GetAll(ctx context.Context, pagination *pagorminator.Pagination) (models.PageResponse[*models.Blog], error)
	GetById(ctx context.Context, id int) (models.Blog, error)
}

var _ BlogService = new(BlogServiceImpl)

type BlogServiceImpl struct {
	Repository repositories.BlogRepository
}

func (b BlogServiceImpl) GetAll(ctx context.Context, pagination *pagorminator.Pagination) (models.PageResponse[*models.Blog], error) {
	blogs, err := b.Repository.GetAll(ctx, pagination)
	if err != nil {
		return models.PageResponse[*models.Blog]{}, err
	}
	return models.PageResponse[*models.Blog]{
		Data: blogs,
		Metadata: models.PageMetadata{
			Page:       pagination.GetPage(),
			Size:       pagination.GetSize(),
			TotalCount: pagination.GetTotalElements(),
			TotalPages: pagination.GetTotalPages(),
		},
	}, nil
}

func (b BlogServiceImpl) GetById(ctx context.Context, id int) (models.Blog, error) {
	return b.Repository.GetById(ctx, id)
}
