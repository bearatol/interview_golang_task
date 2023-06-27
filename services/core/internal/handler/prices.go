package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
	"github.com/gin-gonic/gin"
)

type Prices interface {
	PricesGet(ctx context.Context, barcode string) ([]string, error)
	PriceCreate(ctx context.Context, fileData *mapping.FileData) error
	PricesGetFile(ctx context.Context, fileName string) ([]byte, error)
	PriceDelete(ctx context.Context, fileName string) error
}

// @Summary      Get all prices files by one product
// @Description  get all prices files by one product
// @Tags         prices
// @Param        barcode    query     string  true  "barcode of product"
// @Success      200 {array} string
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products/prices [get]
func (h *Handler) PricesGet(c *gin.Context) {
	barcode := c.Query("barcode")
	if barcode == "" {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid barcode")
		return
	}
	res, err := h.price.PricesGet(h.ctx, barcode)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot get file list")
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary      Get file of price by name
// @Description  get file of price by name
// @Tags         prices
// @Param 		 filename path string true "name of file"
// @Produce  application/pdf
// @Success 	 200 {file} binary
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products/prices/{filename} [get]
func (h *Handler) PricesGetFile(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid file name")
		return
	}

	res, err := h.price.PricesGetFile(h.ctx, fileName)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot get file")
		return
	}

	c.Header("Content-Description", "price file")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)

	c.Data(http.StatusOK, "application/pdf", res)
}

// @Summary      Create price file
// @Description  create price file
// @Tags         prices
// @Param        filedata  body      mapping.FileData true "file data"
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products/prices [post]
func (h *Handler) PriceCreate(c *gin.Context) {
	jsonBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	fileData := &mapping.FileData{}
	if err := json.Unmarshal(jsonBody, fileData); err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "invalid data")
		return
	}
	if len(fileData.Barcode) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid barcode")
		return
	}

	if len(fileData.Title) == 0 {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid title")
		return
	}

	err = h.price.PriceCreate(h.ctx, fileData)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot create price file")
		return
	}

	newSuccessResponse(c)
}

// @Summary      Delete price files
// @Description  delete price files
// @Tags         prices
// @Param        filename    query     string  true  "file name of a price"
// @Success      200 {object} successResponse
// @Failure 	 400 {object} errorResponse
// @Security     BearerAuth
// @Router       /products/prices [delete]
func (h *Handler) PriceDelete(c *gin.Context) {
	fileName := c.Query("filename")
	if fileName == "" {
		newErrorResponse(c, h.log, http.StatusBadRequest, nil, "invalid file name")
		return
	}

	err := h.price.PriceDelete(h.ctx, fileName)
	if err != nil {
		newErrorResponse(c, h.log, http.StatusBadRequest, err, "cannot delete price file")
		return
	}

	newSuccessResponse(c)
}
