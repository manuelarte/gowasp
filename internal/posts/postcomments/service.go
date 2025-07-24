package postcomments

import (
	"context"

	"github.com/manuelarte/pagorminator"

	"github.com/manuelarte/gowasp/internal/models"
)

//nolint:iface // separate repository from service
type Service interface {
	Create(ctx context.Context, postComment *models.PostComment) error
	GetAllForPostID(ctx context.Context, postID uint,
		pagination *pagorminator.Pagination) ([]*models.PostComment, error)
}

var _ Service = new(serviceImpl)

type serviceImpl struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return &serviceImpl{Repository: repository}
}

func (b serviceImpl) GetAllForPostID(ctx context.Context, postID uint,
	pagination *pagorminator.Pagination,
) ([]*models.PostComment, error) {
	postComments, err := b.Repository.GetAllForPostID(ctx, postID, pagination)
	if err != nil {
		return nil, err
	}

	return postComments, nil
}

func (b serviceImpl) Create(ctx context.Context, postComment *models.PostComment) error {
	return b.Repository.Create(ctx, postComment)
}
