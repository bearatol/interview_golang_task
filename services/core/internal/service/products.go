package service

import (
	"context"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

type RepoProducts interface {
	ProductGet(ctx context.Context, login string) ([]*mapping.Product, error)
	ProductCreate(ctx context.Context, login string, product *mapping.ProductAvailableFileds) error
	ProductUpdate(ctx context.Context, product *mapping.ProductAvailableFileds) error
	ProductDelete(ctx context.Context, barcode string) error
}

func (s *Service) ProductGet(ctx context.Context, token string) ([]*mapping.Product, error) {
	login, err := s.login(token)
	if err != nil {
		return nil, err
	}
	return s.repoProducts.ProductGet(ctx, login)
}

func (s *Service) ProductCreate(ctx context.Context, token string, product *mapping.ProductAvailableFileds) error {
	login, err := s.login(token)
	if err != nil {
		return err
	}

	return s.repoProducts.ProductCreate(ctx, login, product)
}

func (s *Service) ProductUpdate(ctx context.Context, product *mapping.ProductAvailableFileds) error {
	return s.repoProducts.ProductUpdate(ctx, product)
}

func (s *Service) ProductDelete(ctx context.Context, barcode string) error {
	return s.repoProducts.ProductDelete(ctx, barcode)
}
