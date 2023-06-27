package service

import (
	"context"
	"fmt"
	"time"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type RepoPrices interface {
	PricesGet(ctx context.Context, barcode string) ([]string, error)
	PriceCreate(ctx context.Context, fileName, barcode string) error
	PriceDelete(ctx context.Context, fileName string) error
}

type PriceGenerator interface {
	Ping(ctx context.Context) error
	Create(ctx context.Context, fileName, barcode, title string, cost int32) error
	Get(ctx context.Context, fileName string) ([]byte, error)
	Delete(ctx context.Context, fileName string) error
}

func (s *Service) PriceCreate(ctx context.Context, fileData *mapping.FileData) error {
	fileName := fmt.Sprintf("doc_%s_%s.pdf", fileData.Barcode, time.Now().Format("01-02-2006"))

	err := s.repoPrices.PriceCreate(ctx, fileName, fileData.Barcode)
	if err != nil {
		return err
	}

	err = s.priceGen.Create(ctx, fileName, fileData.Barcode, fileData.Title, fileData.Cost)
	if err != nil {
		if err := s.repoPrices.PriceDelete(ctx, fileName); err != nil {
			return err
		}
		return err
	}
	return nil
}

func (s *Service) PricesGet(ctx context.Context, barcode string) ([]string, error) {
	return s.repoPrices.PricesGet(ctx, barcode)
}

func (s *Service) PricesGetFile(ctx context.Context, fileName string) ([]byte, error) {
	return s.priceGen.Get(ctx, fileName)
}

func (s *Service) PriceDelete(ctx context.Context, fileName string) error {
	err := s.priceGen.Delete(ctx, fileName)
	if err != nil {
		return err
	}
	return s.repoPrices.PriceDelete(ctx, fileName)
}
