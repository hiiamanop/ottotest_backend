package persistence

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"gorm.io/gorm"
)

type TransactionGormRepository struct {
	db *gorm.DB
}

func NewTransactionGormRepository(db *gorm.DB) *TransactionGormRepository {
	return &TransactionGormRepository{db: db}
}

func (r *TransactionGormRepository) Create(ctx context.Context, trx *entity.Transaction) error {
	return r.db.WithContext(ctx).Create(trx).Error
}

func (r *TransactionGormRepository) FindByID(ctx context.Context, id uint) (*entity.Transaction, error) {
	var trx entity.Transaction
	if err := r.db.WithContext(ctx).Preload("Items").First(&trx, id).Error; err != nil {
		return nil, err
	}
	return &trx, nil
}
