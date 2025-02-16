package services

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gowasp/internal/models"
	"gowasp/internal/repositories"
)

type BlogCommentService interface {
	GetAllForBlog(ctx context.Context, blogID uint, pagination *pagorminator.Pagination) (models.PageResponse[*models.BlogComment], error)
}

var _ BlogCommentService = new(BlogCommentServiceImpl)

type BlogCommentServiceImpl struct {
	Repository repositories.BlogCommentRepository
}

func (b BlogCommentServiceImpl) GetAllForBlog(ctx context.Context, blogID uint, pagination *pagorminator.Pagination) (models.PageResponse[*models.BlogComment], error) {
	blogComments, err := b.Repository.GetAllForBlog(ctx, blogID, pagination)
	if err != nil {
		return models.PageResponse[*models.BlogComment]{}, err
	}
	return models.PageResponse[*models.BlogComment]{
		Data: blogComments,
		Metadata: models.PageMetadata{
			Page:       pagination.GetPage(),
			Size:       pagination.GetSize(),
			TotalCount: pagination.GetTotalElements(),
			TotalPages: pagination.GetTotalPages(),
		},
	}, nil
}
