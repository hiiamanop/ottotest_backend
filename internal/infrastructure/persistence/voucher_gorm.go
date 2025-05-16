package persistence

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"gorm.io/gorm"
)

type VoucherGormRepository struct {
	db *gorm.DB
}

func NewVoucherGormRepository(db *gorm.DB) *VoucherGormRepository {
	return &VoucherGormRepository{db: db}
}

func (r *VoucherGormRepository) Create(ctx context.Context, voucher *entity.Voucher) error {
	return r.db.WithContext(ctx).Create(voucher).Error
}

func (r *VoucherGormRepository) FindByID(ctx context.Context, id uint) (*entity.Voucher, error) {
	var voucher entity.Voucher
	if err := r.db.WithContext(ctx).First(&voucher, id).Error; err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (r *VoucherGormRepository) FindByBrandID(ctx context.Context, brandID uint) ([]entity.Voucher, error) {
	var vouchers []entity.Voucher
	if err := r.db.WithContext(ctx).Where("brand_id = ?", brandID).Find(&vouchers).Error; err != nil {
		return nil, err
	}
	return vouchers, nil
}
