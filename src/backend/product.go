package backend

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}

func getProducts(db *sql.DB) ([]product, error) {
	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Inventory, &p.Name, &p.Price, &p.ProductCode, &p.Status); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, err
}

func (p *product) getProduct(db *sql.DB) error {
	row := db.QueryRow("SELECT productCode, inventory, name, price, status FROM products WHERE id = ?", p.ID)
	return row.Scan(&p.Inventory, &p.Name, &p.Price, &p.ProductCode, &p.Status)
}
