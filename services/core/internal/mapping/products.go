package mapping

import "time"

type Product struct {
	Barcode   string    `db:"barcode" json:"barcode"`
	Name      string    `db:"name" json:"name"`
	Desc      string    `db:"description" json:"description"`
	Cost      int32     `db:"cost" json:"cost"`
	UserID    uint64    `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type ProductAvailableFileds struct {
	Barcode string `db:"barcode" json:"barcode"`
	Name    string `db:"name" json:"name"`
	Desc    string `db:"description" json:"description"`
	Cost    int32  `db:"cost" json:"cost"`
}
