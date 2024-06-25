package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// p := NewProduct("Notebook 2", 1999.00)
	// err = insertProduct(db, *p)
	// if err != nil {
	// 	log.Fatal("Erro ao inserir produto", err)
	// }

	// p.Price = 3500.00
	// err = updateProduct(db, *p)
	// if err != nil {
	// 	log.Fatal("Erro ao alterar produto", err)
	// }

	// p1, err := getProductById2(db, "14636a01-b13c-4ee4-8fba-c275ce12a582")
	// if err != nil {
	// 	log.Fatal("Erro ao buscar produto", err)
	// }

	// resp, err := deleteProduct(db, "14636a01-b13c-4ee4-8fba-c275ce12a582")
	// if err != nil {
	// 	log.Fatal("Erro ao deletar produto", err)
	// }

	// if resp {
	// 	fmt.Println("Sucesso ao excluir registro")
	// }

	products, err := getAllProducts(db)
	if err != nil {
		log.Fatal("Erro ao buscar produtos", err)
	}

	for _, p := range products {
		fmt.Printf("Produto %s Preço %.2f  \n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, product Product) error {
	stm, err := db.Prepare("INSERT INTO products(id, name, price) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduct(db *sql.DB, product Product) error {
	stm, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func getProductById(db *sql.DB, id string) (*Product, error) {
	var p Product
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func getProductById2(db *sql.DB, id string) (*Product, error) {
	var product Product
	stm, err := db.Prepare("SELECT * FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	err = stm.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func getAllProducts(db *sql.DB) ([]Product, error) {
	stm, err := db.Prepare("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return nil, err
	}

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) (bool, error) {
	stm, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer stm.Close()
	_, err = stm.Exec(id)
	if err != nil {
		return false, err
	}

	return true, nil
}
