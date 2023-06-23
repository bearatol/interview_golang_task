package service

import (
	"github.com/bearatol/interview_golang_task/sevices/price_generator/internal/mapping"
)

type PDF interface {
	Set(barcode, title string, cost int32) error
	Get(barcodes []string) ([]*mapping.PriceFile, error)
	Delete(barcodes []string) error
}

func (s *Service) SetFile(barcode, title string, cost int32) error {
	return s.pdf.Set(barcode, title, cost)
}
func (s *Service) GetFilesByBarcodes(barcodes []string) ([]*mapping.PriceFile, error) {
	return s.pdf.Get(barcodes)
}
func (s *Service) DeleteFilesByBarcodes(barcodes []string) error {
	return s.pdf.Delete(barcodes)
}
