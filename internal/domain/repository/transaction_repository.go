package repository

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
)

type TransactionRepository interface {
	Create(ctx context.Context, trx *entity.Transaction) error
	FindByID(ctx context.Context, id uint) (*entity.Transaction, error)
}
