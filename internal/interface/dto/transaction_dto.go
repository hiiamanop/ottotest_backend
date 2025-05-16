package dto

type RedemptionVoucher struct {
	VoucherID uint `json:"voucher_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CreateRedemptionRequest struct {
	CustomerID uint                `json:"customer_id" binding:"required"`
	Vouchers   []RedemptionVoucher `json:"vouchers" binding:"required,dive"`
}

type RedemptionItemResponse struct {
	VoucherID uint `json:"voucher_id"`
	Quantity  int  `json:"quantity"`
	PointCost int  `json:"point_cost"`
}

type RedemptionResponse struct {
	TransactionID uint                     `json:"transaction_id"`
	CustomerID    uint                     `json:"customer_id"`
	TotalPoint    int                      `json:"total_point"`
	Status        string                   `json:"status"`
	CreatedAt     string                   `json:"created_at"`
	Items         []RedemptionItemResponse `json:"items"`
}
