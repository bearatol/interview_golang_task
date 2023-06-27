package repository

import (
	"context"

	"github.com/bearatol/interview_golang_task/sevices/core/internal/mapping"
)

func (r *Repository) ProductGet(ctx context.Context, login string) ([]*mapping.Product, error) {
	query := `
	SELECT * FROM ` + productsTableName + `
	WHERE user_id IN (
		SELECT id FROM ` + usersTableName + `
		WHERE login = $1
	)
	`
	productList := []*mapping.Product{}
	err := r.db.SelectContext(ctx, &productList, query, login)

	return productList, err
}

func (r *Repository) ProductCreate(ctx context.Context, login string, product *mapping.ProductAvailableFileds) error {
	productAll := &mapping.Product{
		Barcode: product.Barcode,
		Name:    product.Name,
		Desc:    product.Desc,
		Cost:    product.Cost,
	}

	if err := r.db.GetContext(
		ctx,
		&productAll.UserID,
		"SELECT id FROM "+usersTableName+" WHERE login = $1",
		login,
	); err != nil {
		return err
	}

	query := `
	INSERT INTO ` + productsTableName + `
		(barcode, name, description, cost, user_id)
	VALUES
		(:barcode, :name, :description, :cost, :user_id)
	`
	if _, err := r.db.NamedExecContext(
		ctx,
		query,
		productAll,
	); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ProductUpdate(ctx context.Context, product *mapping.ProductAvailableFileds) error {
	query := `
	UPDATE ` + productsTableName + `
	SET
		name = :name, description = :description, cost = :cost, updated_at = DEFAULT
	WHERE barcode = :barcode
	`
	_, err := r.db.NamedExecContext(ctx, query, product)
	return err
}
func (r *Repository) ProductDelete(ctx context.Context, barcode string) error {
	query := `
	DELETE FROM ` + productsTableName + `
	WHERE barcode = $1
	`
	_, err := r.db.ExecContext(ctx, query, barcode)
	return err
}
