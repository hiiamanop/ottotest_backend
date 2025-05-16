package service

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/mocks"
	"gorm.io/gorm"
)

func TestRedemptionService_Redeem_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock repository
	mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
	mockVoucherRepo := mocks.NewMockVoucherRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)

	// Dummy DB (tidak dipakai, tapi butuh argumen)
	var dummyDB *gorm.DB

	// Setup
	customer := &entity.Customer{ID: 1, Name: "Budi", Email: "budi@mail.com", PointBalance: 200000}
	voucher := &entity.Voucher{ID: 2, BrandID: 1, Name: "Voucher", PointCost: 50000}

	// Test case: success
	mockCustomerRepo.EXPECT().
		FindByID(gomock.Any(), uint(1)).
		Return(customer, nil)

	mockVoucherRepo.EXPECT().
		FindByID(gomock.Any(), uint(2)).
		Return(voucher, nil)

	mockCustomerRepo.EXPECT().
		UpdatePointBalance(gomock.Any(), uint(1), 150000).
		Return(nil)

	// TransactionRepo.Create tidak dites disini karena DB transaksional
	redemptionSvc := NewRedemptionService(mockCustomerRepo, mockVoucherRepo, mockTransactionRepo, dummyDB)

	items := []RedemptionItem{{VoucherID: 2, Quantity: 1}}

	trx, err := redemptionSvc.Redeem(context.Background(), 1, items)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), trx.CustomerID)
	assert.Equal(t, 50000, trx.TotalPoint)
	assert.Len(t, trx.Items, 1)
	assert.Equal(t, uint(2), trx.Items[0].VoucherID)
}

func TestRedemptionService_Redeem_InsufficientBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
	mockVoucherRepo := mocks.NewMockVoucherRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)
	var dummyDB *gorm.DB

	customer := &entity.Customer{ID: 1, Name: "Budi", Email: "budi@mail.com", PointBalance: 10000}
	voucher := &entity.Voucher{ID: 2, BrandID: 1, Name: "Voucher", PointCost: 50000}

	mockCustomerRepo.EXPECT().
		FindByID(gomock.Any(), uint(1)).
		Return(customer, nil)

	mockVoucherRepo.EXPECT().
		FindByID(gomock.Any(), uint(2)).
		Return(voucher, nil)

	redemptionSvc := NewRedemptionService(mockCustomerRepo, mockVoucherRepo, mockTransactionRepo, dummyDB)
	items := []RedemptionItem{{VoucherID: 2, Quantity: 1}}

	trx, err := redemptionSvc.Redeem(context.Background(), 1, items)
	assert.Nil(t, trx)
	assert.ErrorIs(t, err, ErrInsufficientBalance)
}

func TestRedemptionService_Redeem_VoucherNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCustomerRepo := mocks.NewMockCustomerRepository(ctrl)
	mockVoucherRepo := mocks.NewMockVoucherRepository(ctrl)
	mockTransactionRepo := mocks.NewMockTransactionRepository(ctrl)
	var dummyDB *gorm.DB

	customer := &entity.Customer{ID: 1, Name: "Budi", Email: "budi@mail.com", PointBalance: 200000}

	mockCustomerRepo.EXPECT().
		FindByID(gomock.Any(), uint(1)).
		Return(customer, nil)

	mockVoucherRepo.EXPECT().
		FindByID(gomock.Any(), uint(2)).
		Return(nil, errors.New("record not found"))

	redemptionSvc := NewRedemptionService(mockCustomerRepo, mockVoucherRepo, mockTransactionRepo, dummyDB)
	items := []RedemptionItem{{VoucherID: 2, Quantity: 1}}

	trx, err := redemptionSvc.Redeem(context.Background(), 1, items)
	assert.Nil(t, trx)
	assert.Error(t, err)
}
