package service

import (
	"context"
	"errors"
	"time"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/domain/repository"
	"gorm.io/gorm"
)

var ErrInsufficientBalance = errors.New("insufficient point balance")

type RedemptionItem struct {
	VoucherID uint
	Quantity  int
	PointCost int
}

type RedemptionService struct {
	customerRepo    repository.CustomerRepository
	voucherRepo     repository.VoucherRepository
	transactionRepo repository.TransactionRepository
	db              *gorm.DB
}

func NewRedemptionService(
	customerRepo repository.CustomerRepository,
	voucherRepo repository.VoucherRepository,
	transactionRepo repository.TransactionRepository,
	db *gorm.DB,
) *RedemptionService {
	return &RedemptionService{
		customerRepo:    customerRepo,
		voucherRepo:     voucherRepo,
		transactionRepo: transactionRepo,
		db:              db,
	}
}

func (s *RedemptionService) Redeem(
	ctx context.Context,
	customerID uint,
	items []RedemptionItem,
) (*entity.Transaction, error) {
	var trx *entity.Transaction

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		customer, err := s.customerRepo.FindByID(ctx, customerID)
		if err != nil {
			return err
		}

		totalPoint := 0
		var transactionItems []entity.TransactionItem
		for _, item := range items {
			voucher, err := s.voucherRepo.FindByID(ctx, item.VoucherID)
			if err != nil {
				return err
			}
			subtotal := voucher.PointCost * item.Quantity
			totalPoint += subtotal
			transactionItems = append(transactionItems, entity.TransactionItem{
				VoucherID: item.VoucherID,
				Quantity:  item.Quantity,
				PointCost: voucher.PointCost,
				CreatedAt: time.Now(),
			})
		}
		if customer.PointBalance < totalPoint {
			return ErrInsufficientBalance
		}
		// Update customer balance
		if err = s.customerRepo.UpdatePointBalance(ctx, customerID, customer.PointBalance-totalPoint); err != nil {
			return err
		}
		trx = &entity.Transaction{
			CustomerID: customerID,
			TotalPoint: totalPoint,
			Status:     "SUCCESS",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Items:      transactionItems,
		}
		if err := tx.Create(trx).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return trx, nil
}

func (s *RedemptionService) GetTransactionByID(ctx context.Context, id uint) (*entity.Transaction, error) {
	return s.transactionRepo.FindByID(ctx, id)
}
