package persistence

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"gorm.io/gorm"
)

type CustomerGormRepository struct {
	db *gorm.DB
}

func NewCustomerGormRepository(db *gorm.DB) *CustomerGormRepository {
	return &CustomerGormRepository{db: db}
}

func (r *CustomerGormRepository) FindByID(ctx context.Context, id uint) (*entity.Customer, error) {
	var customer entity.Customer
	if err := r.db.WithContext(ctx).First(&customer, id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *CustomerGormRepository) UpdatePointBalance(ctx context.Context, id uint, newBalance int) error {
	return r.db.WithContext(ctx).Model(&entity.Customer{}).Where("id = ?", id).Update("point_balance", newBalance).Error
}

func (r *CustomerGormRepository) Create(ctx context.Context, c *entity.Customer) error {
	return r.db.WithContext(ctx).Create(c).Error
}
