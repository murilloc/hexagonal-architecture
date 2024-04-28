package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/murilloc/go-hexagonal/adapters/db"
	"github.com/murilloc/go-hexagonal/application"
)

func main() {

	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Test", 19.90)
	if err != nil {
		panic(err)
	}
	productService.Enable(product)

}
