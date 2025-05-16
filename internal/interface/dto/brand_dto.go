package dto

type CreateBrandRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type BrandResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
