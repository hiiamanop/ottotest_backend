package service

import (
	"context"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/domain/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, c *entity.Customer) error {
	return s.repo.Create(ctx, c)
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id uint) (*entity.Customer, error) {
	return s.repo.FindByID(ctx, id)
}
