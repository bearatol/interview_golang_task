package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/gin-gonic/gin"
)

type Products interface {
	ProductGet(ctx context.Context, token string) ([]*mapping.Product, error)
	ProductCreate(ctx context.Context, token string, product *mapping.ProductAvailableFileds) error
	ProductUpdate(ctx context.Context, product *mapping.ProductAvailableFileds) error
	ProductDelete(ctx context.Context, barcode string) error
}

// @Summary      Get products information
// @Description  get products information
// @Tags         products
// @Success      200 {array} mapping.Product
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products [get]
func (h *Handler) ProductGet(c *gin.Context) {
	// token set in middleware
	token := c.Param("token")
	res, err := h.product.ProductGet(h.ctx, token)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot get products")
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary      Create product
// @Description  create product
// @Tags         products
// @Param        product  body      mapping.ProductAvailableFileds true "product data"
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products [post]
func (h *Handler) ProductCreate(c *gin.Context) {
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	product := &mapping.ProductAvailableFileds{}
	if err := json.Unmarshal(jsonBody, product); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	if len(product.Barcode) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid barcode")
		return
	}
	if len(product.Name) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid name")
		return
	}
	if len(product.Desc) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid description")
		return
	}

	// token set in middleware
	token := c.Param("token")

	if err := h.product.ProductCreate(h.ctx, token, product); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot create product")
		return
	}

	newSuccessResponse(c)
}

// @Summary      Update product
// @Description  update product
// @Tags         products
// @Param        product  body      mapping.ProductAvailableFileds true "product data"
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products [put]
func (h *Handler) ProductUpdate(c *gin.Context) {
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	product := &mapping.ProductAvailableFileds{}
	if err := json.Unmarshal(jsonBody, product); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	if len(product.Barcode) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid barcode")
		return
	}

	if err := h.product.ProductUpdate(h.ctx, product); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot update product")
		return
	}

	newSuccessResponse(c)
}

// @Summary      Delete product
// @Description  delete product
// @Tags         products
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Param        barcode    query     string  true  "barcode"
// @Security     BearerAuth
// @Router       /products [delete]
func (h *Handler) ProductDelete(c *gin.Context) {
	barcode := c.Query("barcode")
	if len(barcode) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid barcode")
		return
	}

	err := h.product.ProductDelete(h.ctx, barcode)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot delete product")
		return
	}

	newSuccessResponse(c)
}
