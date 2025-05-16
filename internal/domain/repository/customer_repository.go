package repository

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
)

type CustomerRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.Customer, error)
	UpdatePointBalance(ctx context.Context, id uint, newBalance int) error
	Create(ctx context.Context, c *entity.Customer) error // Tambahkan Create
}
