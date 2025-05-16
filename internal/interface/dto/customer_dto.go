package dto

type CreateCustomerRequest struct {
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	PointBalance int    `json:"point_balance" binding:"required,gte=0"`
}

type CustomerResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PointBalance int    `json:"point_balance"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
