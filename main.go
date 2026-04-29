package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	id    int64
	name  string
	price int32
}

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://postgres:2006@localhost:5432/products?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * from products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.id, &product.name, &product.price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for _, pr := range products {
		fmt.Printf("%d, %s, %d\n", pr.id, pr.name, pr.price)
	}

}
