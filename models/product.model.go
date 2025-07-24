package models

import (
	"context"
	"go-test/utils"
	"log"

	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Stock         int    `json:"stock"`
	PurchasePrice int    `json:"purchase_price"`
	SellingPrice  int    `json:"selling_price"`
	ImageURL      string `json:"image_url"`
	CategoryName  string `json:"category_name"`
}

type CategoryProduct struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetProducts() ([]Product, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		log.Println("DB Connection Error:", err)
		return []Product{}, err
	}
	defer conn.Close()

	query := `
	SELECT 
 		 p.id, 
 		 p.name, 
 		 p.stock, 
 		 p.purchase_price, 
 		 p.selling_price, 
 		 p.image_url, 
 		 cp.name AS category_name
	FROM products p
	JOIN category_products cp ON cp.id = p.category_id;`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println("Query Error:", err)
		return []Product{}, err
	}

	products, err := pgx.CollectRows[Product](rows, pgx.RowToStructByName)
	if err != nil {
		log.Println("Mapping Error:", err)
		return []Product{}, err
	}

	return products, err

}

func GetProductByID(id int) (Product, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		log.Println("DB Connection Error:", err)
		return Product{}, err
	}
	defer conn.Close()

	query := `
	SELECT 
 		 p.id, 
 		 p.name, 
 		 p.stock, 
 		 p.purchase_price, 
 		 p.selling_price, 
 		 p.image_url, 
 		 cp.name AS category_name
	FROM products p
	JOIN category_products cp ON cp.id = p.category_id
	WHERE p.id = $1;`
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		log.Println("Query Error:", err)
		return Product{}, err
	}

	product, err := pgx.CollectOneRow[Product](rows, pgx.RowToStructByName)
	if err != nil {
		log.Println("Mapping Error:", err)
		return Product{}, err
	}

	return product, err
}

func GetCategoryProducts() ([]CategoryProduct, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		log.Println("DB Connection Error:", err)
		return []CategoryProduct{}, err
	}
	defer conn.Close()

	query := `
	SELECT id, name, description FROM category_products;`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println("Query Error:", err)
		return []CategoryProduct{}, err
	}

	categoryProducts, err := pgx.CollectRows[CategoryProduct](rows, pgx.RowToStructByName)
	if err != nil {
		log.Println("Mapping Error:", err)
		return []CategoryProduct{}, err
	}

	return categoryProducts, err

}
