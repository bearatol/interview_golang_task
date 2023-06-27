package mapping

type FileData struct {
	Barcode string `json:"barcode"`
	Title   string `json:"title"`
	Cost    int32  `json:"cost"`
}

type ProductPrice struct {
	Name           string `db:"name"`
	ProductBarcode string `db:"product_barcode"`
}
