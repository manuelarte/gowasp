package repositories

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"
	"gowasp/internal/models"
)

type BlogRepository interface {
	GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Blog, error)
	GetById(ctx context.Context, id int) (models.Blog, error)
}

var _ BlogRepository = new(BlogRepositoryDB)

type BlogRepositoryDB struct {
	DB *gorm.DB
}

func (b BlogRepositoryDB) GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Blog, error) {
	var blogs []*models.Blog
	tx := b.DB.WithContext(ctx).Clauses(pageRequest).Order("posted_at asc").Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return blogs, tx.Error
}

func (b BlogRepositoryDB) GetById(ctx context.Context, id int) (models.Blog, error) {
	var blog models.Blog
	tx := b.DB.WithContext(ctx).First(&blog, id)
	if tx.Error != nil {
		return models.Blog{}, tx.Error
	}
	return blog, nil
}
