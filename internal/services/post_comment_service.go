package services

import (
	"context"
	"github.com/manuelarte/gowasp/internal/models"
	"github.com/manuelarte/gowasp/internal/repositories"

	"github.com/manuelarte/pagorminator"
)

type PostCommentService interface {
	Create(ctx context.Context, postComment *models.PostComment) error
	GetAllForPostID(ctx context.Context, postID uint64,
		pagination *pagorminator.Pagination) (models.PageResponse[*models.PostComment], error)
}

var _ PostCommentService = new(PostCommentServiceImpl)

type PostCommentServiceImpl struct {
	Repository repositories.PostCommentRepository
}

func (b PostCommentServiceImpl) GetAllForPostID(ctx context.Context, postID uint64,
	pagination *pagorminator.Pagination) (models.PageResponse[*models.PostComment], error) {
	postComments, err := b.Repository.GetAllForPostID(ctx, postID, pagination)
	if err != nil {
		return models.PageResponse[*models.PostComment]{}, err
	}
	return models.PageResponse[*models.PostComment]{
		Data: postComments,
		Metadata: models.PageMetadata{
			Page:       pagination.GetPage(),
			Size:       pagination.GetSize(),
			TotalCount: pagination.GetTotalElements(),
			TotalPages: pagination.GetTotalPages(),
		},
	}, nil
}

func (b PostCommentServiceImpl) Create(ctx context.Context, postComment *models.PostComment) error {
	return b.Repository.Create(ctx, postComment)
}
