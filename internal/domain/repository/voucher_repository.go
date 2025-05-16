package repository

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
)

type VoucherRepository interface {
	Create(ctx context.Context, voucher *entity.Voucher) error
	FindByID(ctx context.Context, id uint) (*entity.Voucher, error)
	FindByBrandID(ctx context.Context, brandID uint) ([]entity.Voucher, error)
}
