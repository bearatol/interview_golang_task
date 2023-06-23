package service

import (
	"context"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type PriceGenerator interface {
	Create(ctx context.Context, barcode, name string, cost int32) error
	Get(ctx context.Context, barcodes []string) ([]*mapping.PriceFile, error)
	Delete(ctx context.Context, barcodes []string) error
}

func (s *Service) PriceGeneratorCreate(ctx context.Context, barcode, name string, cost int32) error {
	return s.priceGen.Create(ctx, barcode, name, cost)
}
func (s *Service) PriceGeneratorGet(ctx context.Context, barcodes []string) ([]*mapping.PriceFile, error) {
	return s.priceGen.Get(ctx, barcodes)
}
func (s *Service) PriceGeneratorDelete(ctx context.Context, barcodes []string) error {
	return s.priceGen.Delete(ctx, barcodes)
}
