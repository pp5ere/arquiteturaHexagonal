package db_test

import (
	"database/sql"
	"log"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pp5ere/hexagonal/adapters/db"
	"github.com/pp5ere/hexagonal/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp()  {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTables(Db)
	createProduct(Db)
}

func createTables(db *sql.DB)  {
	table := `Create table products ("id" string, "name" string, "price" float, "status" string);`
	stmt, err := db.Prepare(table);if err != nil {
		log.Fatalln(err.Error())
	}
	stmt.Exec()
}

func createProduct(db * sql.DB)  {
	insert := `insert into products values("asdjkl", "sniker", 23.6, "disabled")`
	stmt, err := db.Prepare(insert);if err != nil {
		log.Fatalln(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T)  {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("asdjkl")
	require.Nil(t, err)
	require.Equal(t, "sniker", product.GetName())
	require.Equal(t, 23.6, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProducDb_Save(t *testing.T)  {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.Product{Id: "asd", Name: "new", Price: 12.6, Status: "disabled"}

	//insert test
	productResult, err := productDb.Save(&product)
	require.Nil(t, err)
	require.Equal(t, product.GetId(), productResult.GetId())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())

	product.Name = "other"
	product.Price = 25.4
	product.Status = "enabled"

	//update test
	productResult, err = productDb.Save(&product)
	require.Nil(t, err)
	require.Equal(t, product.GetId(), productResult.GetId())
	require.Equal(t, product.GetName(), productResult.GetName())
	require.Equal(t, product.GetPrice(), productResult.GetPrice())
	require.Equal(t, product.GetStatus(), productResult.GetStatus())
}