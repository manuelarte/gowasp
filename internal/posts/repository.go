package posts

import (
	"context"

	"github.com/manuelarte/pagorminator"
	"gorm.io/gorm"

	"github.com/manuelarte/gowasp/internal/models"
)

//nolint:iface // separate repository from service
type Repository interface {
	GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Post, error)
	GetByID(ctx context.Context, id uint64) (models.Post, error)
}

var _ Repository = new(gormRepository)

type gormRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{DB: db}
}

func (b gormRepository) GetAll(ctx context.Context, pageRequest *pagorminator.Pagination) ([]*models.Post, error) {
	var posts []*models.Post
	tx := b.DB.WithContext(ctx).Clauses(pageRequest).Find(&posts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return posts, tx.Error
}

func (b gormRepository) GetByID(ctx context.Context, id uint64) (models.Post, error) {
	var post models.Post
	tx := b.DB.WithContext(ctx).First(&post, id)
	if tx.Error != nil {
		return models.Post{}, tx.Error
	}

	return post, nil
}
