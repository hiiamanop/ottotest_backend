package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/infrastructure/persistence"
	"github.com/hiiamanop/ottotest_backend/internal/interface/handler"
	"github.com/hiiamanop/ottotest_backend/internal/usecase/service"
	"github.com/hiiamanop/ottotest_backend/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	db := config.InitDB(cfg)
	fmt.Println("Database connected!")

	// Auto-migrate Brand model (opsional, pastikan schema match)
	db.AutoMigrate(&entity.Brand{}, &entity.Voucher{}, &entity.Customer{}, &entity.Transaction{}, &entity.TransactionItem{})

	// Setup repository & service
	brandRepo := persistence.NewBrandGormRepository(db)
	brandService := service.NewBrandService(brandRepo)
	brandHandler := handler.NewBrandHandler(brandService)

	voucherRepo := persistence.NewVoucherGormRepository(db)
	voucherService := service.NewVoucherService(voucherRepo)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	customerRepo := persistence.NewCustomerGormRepository(db)
	transactionRepo := persistence.NewTransactionGormRepository(db)
	redemptionService := service.NewRedemptionService(customerRepo, voucherRepo, transactionRepo, db)
	redemptionHandler := handler.NewRedemptionHandler(redemptionService)

	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	r := gin.Default()
	r.POST("/brand", brandHandler.CreateBrand)
	r.POST("/voucher", voucherHandler.CreateVoucher)
	r.GET("/voucher", voucherHandler.GetVoucherByID)
	r.GET("/voucher/brand", voucherHandler.GetVouchersByBrand)
	r.POST("/transaction/redemption", redemptionHandler.CreateRedemption)
	r.GET("/transaction/redemption", redemptionHandler.GetRedemptionDetail)
	r.POST("/customer", customerHandler.CreateCustomer)
	r.GET("/customer", customerHandler.GetCustomer)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
