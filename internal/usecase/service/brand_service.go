package service

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/domain/repository"
)

type BrandService struct {
	repo repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) CreateBrand(ctx context.Context, brand *entity.Brand) error {
	return s.repo.Create(ctx, brand)
}
