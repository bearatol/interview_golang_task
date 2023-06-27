package repository

import (
	"context"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

func (r *Repository) PricesGet(ctx context.Context, barcode string) ([]string, error) {
	query := `
	SELECT name FROM ` + priceTableName + `
	WHERE product_barcode = $1
	`
	prices := []*mapping.ProductPrice{}
	if err := r.db.SelectContext(ctx, &prices, query, barcode); err != nil {
		return nil, err
	}

	pricesList := make([]string, len(prices))
	for k, p := range prices {
		pricesList[k] = p.Name
	}
	return pricesList, nil
}

func (r *Repository) PriceCreate(ctx context.Context, fileName, barcode string) error {
	query := `
	INSERT INTO ` + priceTableName + ` (name, product_barcode)
	VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, fileName, barcode)
	return err
}
func (r *Repository) PriceDelete(ctx context.Context, fileName string) error {
	query := `
	DELETE FROM ` + priceTableName + ` WHERE name = $1 AND product_barcode = $2
	`
	_, err := r.db.ExecContext(ctx, query, fileName)
	return err
}
