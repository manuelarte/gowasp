package postcomments

import (
	"context"

	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
)

type Service interface {
	Create(ctx context.Context, postComment *models.PostComment) error
	GetAllForPostID(ctx context.Context, postID uint64,
		pagination *pagorminator.Pagination) (models.PageResponse[*models.PostComment], error)
}

var _ Service = new(serviceImpl)

type serviceImpl struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{Repository: repository}
}

func (b serviceImpl) GetAllForPostID(ctx context.Context, postID uint64,
	pagination *pagorminator.Pagination,
) (models.PageResponse[*models.PostComment], error) {
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

func (b serviceImpl) Create(ctx context.Context, postComment *models.PostComment) error {
	return b.Repository.Create(ctx, postComment)
}
