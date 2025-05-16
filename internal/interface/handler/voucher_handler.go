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

type VoucherHandler struct {
	voucherService *service.VoucherService
}

func NewVoucherHandler(voucherService *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{voucherService: voucherService}
}

// CreateVoucher godoc
// @Summary      Create voucher
// @Description  Add a new voucher
// @Tags         voucher
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateVoucherRequest true "Voucher Data"
// @Success      201  {object} dto.VoucherResponse
// @Failure      400  {object} map[string]string
// @Router       /voucher [post]
func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var req dto.CreateVoucherRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voucher := &entity.Voucher{
		BrandID:     req.BrandID,
		Name:        req.Name,
		Description: req.Description,
		PointCost:   req.PointCost,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := h.voucherService.CreateVoucher(c.Request.Context(), voucher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.VoucherResponse{
		ID:          voucher.ID,
		BrandID:     voucher.BrandID,
		Name:        voucher.Name,
		Description: voucher.Description,
		PointCost:   voucher.PointCost,
		CreatedAt:   voucher.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   voucher.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusCreated, resp)
}

// GetVoucherByID godoc
// @Summary      Get voucher by ID
// @Description  Get voucher details by ID
// @Tags         voucher
// @Accept       json
// @Produce      json
// @Param        id query uint true "Voucher ID"
// @Success      200  {object} dto.VoucherResponse
// @Failure      400  {object} map[string]string
// @Failure      404  {object} map[string]string
// @Router       /voucher [get]
func (h *VoucherHandler) GetVoucherByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	voucher, err := h.voucherService.GetVoucherByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}

	resp := dto.VoucherResponse{
		ID:          voucher.ID,
		BrandID:     voucher.BrandID,
		Name:        voucher.Name,
		Description: voucher.Description,
		PointCost:   voucher.PointCost,
		CreatedAt:   voucher.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   voucher.UpdatedAt.Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, resp)
}

// GetVouchersByBrand godoc
// @Summary      Get vouchers by brand ID
// @Description  Get all vouchers for a specific brand
// @Tags         voucher
// @Accept       json
// @Produce      json
// @Param        id query uint true "Brand ID"
// @Success      200  {array} dto.VoucherResponse
// @Failure      400  {object} map[string]string
// @Failure      404  {object} map[string]string
// @Router       /voucher/brand [get]
func (h *VoucherHandler) GetVouchersByBrand(c *gin.Context) {
	brandIDStr := c.Query("id")
	brandID, err := strconv.ParseUint(brandIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid brand id"})
		return
	}

	vouchers, err := h.voucherService.GetVouchersByBrand(c.Request.Context(), uint(brandID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []dto.VoucherResponse
	for _, v := range vouchers {
		resp = append(resp, dto.VoucherResponse{
			ID:          v.ID,
			BrandID:     v.BrandID,
			Name:        v.Name,
			Description: v.Description,
			PointCost:   v.PointCost,
			CreatedAt:   v.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   v.UpdatedAt.Format(time.RFC3339),
		})
	}
	c.JSON(http.StatusOK, resp)
}
