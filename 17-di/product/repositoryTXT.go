package product

import "database/sql"

type ProductRepositoryTXT struct {
	db *sql.DB
}

func NewProductRepositoryTXT(db *sql.DB) *ProductRepositoryTXT {
	return &ProductRepositoryTXT{db}
}

func (r *ProductRepositoryTXT) GetProduct(id int) (Product, error) {
	return Product{ID: id, Name: "Product Name Example"}, nil
}
