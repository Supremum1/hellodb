/*
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

*/

// HTTP

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Product struct {
	id    int64
	name  string
	price int32
}

func init() {
	var err error
	db, err = sql.Open(
		"postgres",
		"postgres://postgres:2006@localhost:5432/products?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

var db *sql.DB

func main() {
	http.HandleFunc("/products", productsIndex)
	http.ListenAndServe(":3000", nil)

}
func productsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)

	}
	rows, err := db.Query("SELECT * from products")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)

	}
	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.id, &product.name, &product.price)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
	for _, pr := range products {
		fmt.Printf("%d, %s, %d\n", pr.id, pr.name, pr.price)
	}

}
