package repositories

import (
	"context"

	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

type PostCommentRepository interface {
	Create(ctx context.Context, postComment *models.PostComment) error
	GetAllForPostID(ctx context.Context, postID uint64,
		pageRequest *pagorminator.Pagination) ([]*models.PostComment, error)
}

var _ PostCommentRepository = new(PostCommentRepositoryDB)

type PostCommentRepositoryDB struct {
	DB *gorm.DB
}

func (b PostCommentRepositoryDB) GetAllForPostID(ctx context.Context, postID uint64,
	pageRequest *pagorminator.Pagination,
) ([]*models.PostComment, error) {
	var postComments []*models.PostComment
	tx := b.DB.WithContext(ctx).Clauses(pageRequest).Order("posted_at asc").
		Where("post_id = ?", postID).Preload("User").Find(&postComments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return postComments, tx.Error
}

func (b PostCommentRepositoryDB) Create(ctx context.Context, postComment *models.PostComment) error {
	return b.DB.WithContext(ctx).Create(postComment).Error
}
