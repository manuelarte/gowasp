package repositories

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"
	"gowasp/internal/models"
)

type BlogCommentRepository interface {
	Create(ctx context.Context, blogComment *models.BlogComment) error
	GetAllForBlog(ctx context.Context, blogID uint, pageRequest *pagorminator.Pagination) ([]*models.BlogComment, error)
}

var _ BlogCommentRepository = new(BlogCommentRepositoryDB)

type BlogCommentRepositoryDB struct {
	DB *gorm.DB
}

func (b BlogCommentRepositoryDB) GetAllForBlog(ctx context.Context, blogID uint, pageRequest *pagorminator.Pagination) ([]*models.BlogComment, error) {
	var blogComments []*models.BlogComment
	tx := b.DB.WithContext(ctx).Clauses(pageRequest).Order("posted_at asc").Where("blog_id = ?", blogID).Find(&blogComments)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return blogComments, tx.Error
}

func (b BlogCommentRepositoryDB) Create(ctx context.Context, blogComment *models.BlogComment) error {
	return b.DB.WithContext(ctx).Create(blogComment).Error
}
