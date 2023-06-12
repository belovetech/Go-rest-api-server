package backend

import "database/sql"

type order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customerName"`
	Total        int    `json:"total"`
	Status       string `json:"status"`
}

func (o *order) createOrder(db *sql.DB) error {
	res, err := db.Exec("INSERT INTO orders(customerName, total, status) VALUES(?, ?, ?)", o.CustomerName, o.Total, o.Status)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	o.ID = int(id)
	return nil
}

func getOrders(db *sql.DB) ([]order, error) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orders := []order{}
	for rows.Next() {
		var o order
		if err := rows.Scan(&o.ID, &o.CustomerName, &o.Total, &o.Status); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, err
}

func (o *order) getOrder(db *sql.DB) error {
	row := db.QueryRow("SELECT id, customerName, total, status FROM orders WHERE id = ?", o.ID)
	return row.Scan(&o.ID, &o.CustomerName, &o.Total, &o.Status)
}
