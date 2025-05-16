package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/interface/dto"
	"github.com/hiiamanop/ottotest_backend/internal/usecase/service"
)

type CustomerHandler struct {
	customerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

// POST /customer
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer := &entity.Customer{
		Name:         req.Name,
		Email:        req.Email,
		PointBalance: req.PointBalance,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := h.customerService.CreateCustomer(c.Request.Context(), customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := dto.CustomerResponse{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		PointBalance: customer.PointBalance,
		CreatedAt:    customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    customer.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, resp)
}

// GET /customer?id={id}
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	customer, err := h.customerService.GetCustomerByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	resp := dto.CustomerResponse{
		ID:           customer.ID,
		Name:         customer.Name,
		Email:        customer.Email,
		PointBalance: customer.PointBalance,
		CreatedAt:    customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    customer.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, resp)
}
