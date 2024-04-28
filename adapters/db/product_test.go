package db_test

import (
	"database/sql"
	"github.com/murilloc/go-hexagonal/adapters/db"
	"github.com/murilloc/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
    		"id" string,
    		"name" string,
    		"price" float,
    		"status" string
    		    		);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	exec, err := stmt.Exec()
	if err != nil {
		return
	}
	log.Println(exec)
}

func createProduct(db *sql.DB) {
	insertProduct := `INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)`
	stmt, err := db.Prepare(insertProduct)
	if err != nil {
		log.Fatal(err.Error())
	}
	exec, err := stmt.Exec("1", "Product Test", 19.90, "disabled")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(exec)
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(Db)

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("1")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 19.90, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer func(Db *sql.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(Db)

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.ID = "2"
	product.Name = "Product Test 2"
	product.Price = 19.90
	product.Status = "enabled"

	productResult, err := productDb.Save(product)
	require.Nil(t, err)

	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "disabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

}
