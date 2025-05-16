package dto

type CreateVoucherRequest struct {
	BrandID     uint   `json:"brand_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	PointCost   int    `json:"point_cost" binding:"required"`
}

type VoucherResponse struct {
	ID          uint   `json:"id"`
	BrandID     uint   `json:"brand_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PointCost   int    `json:"point_cost"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
