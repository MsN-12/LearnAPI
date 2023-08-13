package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func getProducts(db *sql.DB) ([]Product, error) {
	query := "SELECT id, name , quantity, price from products"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
func (p *Product) getProduct(db *sql.DB) error {
	query := fmt.Sprintf("SELECT name, quantity, price, product From Products where id=%v", p.ID)
	row := db.QueryRow(query)
	err := row.Scan(&p.Name, &p.Quantity, &p.Price)
	if err != nil {
		return err
	}
	return nil
}
