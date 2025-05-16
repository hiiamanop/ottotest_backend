package repository

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
)

type TransactionItemRepository interface {
	Create(ctx context.Context, item *entity.TransactionItem) error
	FindByTransactionID(ctx context.Context, transactionID uint) ([]entity.TransactionItem, error)
}
