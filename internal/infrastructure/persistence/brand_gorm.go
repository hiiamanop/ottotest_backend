package persistence

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"gorm.io/gorm"
)

type BrandGormRepository struct {
	db *gorm.DB
}

func NewBrandGormRepository(db *gorm.DB) *BrandGormRepository {
	return &BrandGormRepository{db: db}
}

func (r *BrandGormRepository) Create(ctx context.Context, brand *entity.Brand) error {
	return r.db.WithContext(ctx).Create(brand).Error
}
