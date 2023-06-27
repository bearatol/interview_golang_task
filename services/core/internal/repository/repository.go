package repository

import "github.com/jmoiron/sqlx"

const (
	usersTableName    = "users"
	productsTableName = "products"
	priceTableName    = "product_price"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db}
}
