package repositories

import (
	"context"

	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

type PostRepository interface {
	GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Post, error)
	GetByID(ctx context.Context, id uint64) (models.Post, error)
}

var _ PostRepository = new(PostRepositoryDB)

type PostRepositoryDB struct {
	DB *gorm.DB
}

func (b PostRepositoryDB) GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Post, error) {
	var posts []*models.Post
	tx := b.DB.WithContext(ctx).Clauses(pageRequest).Order("posted_at asc").Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return posts, tx.Error
}

func (b PostRepositoryDB) GetByID(ctx context.Context, id uint64) (models.Post, error) {
	var post models.Post
	tx := b.DB.WithContext(ctx).First(&post, id)
	if tx.Error != nil {
		return models.Post{}, tx.Error
	}
	return post, nil
}
