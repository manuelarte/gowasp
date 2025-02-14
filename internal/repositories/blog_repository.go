package repositories

import (
	"context"
	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"
	"gowasp/internal/models"
)

type BlogRepository interface {
	GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Blog, error)
}

var _ BlogRepository = new(BlogRepositoryDB)

type BlogRepositoryDB struct {
	DB *gorm.DB
}

func (b BlogRepositoryDB) GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Blog, error) {
	var blogs []*models.Blog
	tx := b.DB.WithContext(ctx).Debug().Clauses(pageRequest).Find(&blogs)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return blogs, tx.Error
}
