package repository

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
)

type BrandRepository interface {
	Create(ctx context.Context, brand *entity.Brand) error
}
