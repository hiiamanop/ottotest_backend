package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hiiamanop/ottotest_backend/internal/interface/dto"
	"github.com/hiiamanop/ottotest_backend/internal/usecase/service"
)

type RedemptionHandler struct {
	redemptionService *service.RedemptionService
}

func NewRedemptionHandler(redemptionService *service.RedemptionService) *RedemptionHandler {
	return &RedemptionHandler{redemptionService: redemptionService}
}

// POST /transaction/redemption
func (h *RedemptionHandler) CreateRedemption(c *gin.Context) {
	var req dto.CreateRedemptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var items []service.RedemptionItem
	for _, v := range req.Vouchers {
		items = append(items, service.RedemptionItem{
			VoucherID: v.VoucherID,
			Quantity:  v.Quantity,
		})
	}
	trx, err := h.redemptionService.Redeem(c.Request.Context(), req.CustomerID, items)
	if err != nil {
		if err == service.ErrInsufficientBalance {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient point balance"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var respItems []dto.RedemptionItemResponse
	for _, item := range trx.Items {
		respItems = append(respItems, dto.RedemptionItemResponse{
			VoucherID: item.VoucherID,
			Quantity:  item.Quantity,
			PointCost: item.PointCost,
		})
	}
	resp := dto.RedemptionResponse{
		TransactionID: trx.ID,
		CustomerID:    trx.CustomerID,
		TotalPoint:    trx.TotalPoint,
		Status:        trx.Status,
		CreatedAt:     trx.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Items:         respItems,
	}
	c.JSON(http.StatusCreated, resp)
}

// GET /transaction/redemption?transactionId={id}
func (h *RedemptionHandler) GetRedemptionDetail(c *gin.Context) {
	trxIDStr := c.Query("transactionId")
	trxID, err := strconv.ParseUint(trxIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction id"})
		return
	}
	trx, err := h.redemptionService.GetTransactionByID(c.Request.Context(), uint(trxID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}
	var respItems []dto.RedemptionItemResponse
	for _, item := range trx.Items {
		respItems = append(respItems, dto.RedemptionItemResponse{
			VoucherID: item.VoucherID,
			Quantity:  item.Quantity,
			PointCost: item.PointCost,
		})
	}
	resp := dto.RedemptionResponse{
		TransactionID: trx.ID,
		CustomerID:    trx.CustomerID,
		TotalPoint:    trx.TotalPoint,
		Status:        trx.Status,
		CreatedAt:     trx.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Items:         respItems,
	}
	c.JSON(http.StatusOK, resp)
}
