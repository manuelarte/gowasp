package postcomments

import (
	"context"

	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

type Repository interface {
	Create(ctx context.Context, postComment *models.PostComment) error
	GetAllForPostID(ctx context.Context, postID uint,
		pageRequest *pagorminator.Pagination) ([]*models.PostComment, error)
}

var _ Repository = new(gormRepository)

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (b gormRepository) GetAllForPostID(ctx context.Context, postID uint,
	pageRequest *pagorminator.Pagination,
) ([]*models.PostComment, error) {
	var postComments []*models.PostComment
	tx := b.db.WithContext(ctx).Clauses(pageRequest).Order("posted_at asc").
		Where("post_id = ?", postID).Preload("User").Find(&postComments)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return postComments, tx.Error
}

func (b gormRepository) Create(ctx context.Context, postComment *models.PostComment) error {
	return b.db.WithContext(ctx).Create(postComment).Error
}
