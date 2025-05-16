package service

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/domain/repository"
)

type VoucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) *VoucherService {
	return &VoucherService{repo: repo}
}

func (s *VoucherService) CreateVoucher(ctx context.Context, voucher *entity.Voucher) error {
	return s.repo.Create(ctx, voucher)
}

func (s *VoucherService) GetVoucherByID(ctx context.Context, id uint) (*entity.Voucher, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *VoucherService) GetVouchersByBrand(ctx context.Context, brandID uint) ([]entity.Voucher, error) {
	return s.repo.FindByBrandID(ctx, brandID)
}
