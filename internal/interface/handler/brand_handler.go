package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hiiamanop/ottotest_backend/internal/domain/entity"
	"github.com/hiiamanop/ottotest_backend/internal/interface/dto"
	"github.com/hiiamanop/ottotest_backend/internal/usecase/service"
)

type BrandHandler struct {
	brandService *service.BrandService
}

func NewBrandHandler(brandService *service.BrandService) *BrandHandler {
	return &BrandHandler{brandService: brandService}
}

// CreateBrand godoc
// @Summary Create a new brand
// @Description Create a new brand
// @Tags brand
// @Accept json
// @Produce json
// @Param request body dto.CreateBrandRequest true "Create brand request"
// @Success 201 {object} dto.BrandResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /brand [post]
func (h *BrandHandler) CreateBrand(c *gin.Context) {
	var req dto.CreateBrandRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	brand := &entity.Brand{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := h.brandService.CreateBrand(c.Request.Context(), brand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.BrandResponse{
		ID:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
		CreatedAt:   brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   brand.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, resp)
}
