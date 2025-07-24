package models

import (
	"context"
	"fmt"
	"go-test/utils"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	ID              int       `json:"id"`
	ProductID       int       `json:"product_id"`
	TransactionType string    `json:"transaction_type"`
	Quantity        int       `json:"quantity"`
	UserID          int       `json:"user_id"`
	CreatedAt       time.Time `json:"created_at"`
}
type TransactionRequest struct {
	ProductID       int    `json:"product_id"`
	TransactionType string `json:"transaction_type"`
	Quantity        int    `json:"quantity"`
}

type TransactionHistoryResponse struct {
	ID            string  `json:"id"`
	ProductID     string  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	CategoryName  string  `json:"category_name"`
	TotalIn       int     `json:"total_in"`
	TotalOut      int     `json:"total_out"`
	PurchasePrice float64 `json:"purchase_price"`
	SellingPrice  float64 `json:"selling_price"`
	TotalPurchase float64 `json:"total_purchase"`
	TotalSales    float64 `json:"total_sales"`
	StockLeft     int     `json:"stock_left"`
}

func CreateTransaction(t TransactionRequest, userID int) (*Transaction, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return nil, err
	}
	fmt.Println("transaksi:", t)
	fmt.Println("userid:", userID)

	query := `
		INSERT INTO transactions (product_id, transaction_type, quantity, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, product_id, transaction_type, quantity, user_id, created_at
	`
	var newTx Transaction
	err = conn.QueryRow(context.Background(), query,
		t.ProductID, t.TransactionType, t.Quantity, userID).
		Scan(&newTx.ID, &newTx.ProductID, &newTx.TransactionType, &newTx.Quantity, &newTx.UserID, &newTx.CreatedAt)

	if err != nil {
		return nil, err
	}

	stockChange := t.Quantity
	if t.TransactionType == "out" {
		stockChange = -t.Quantity
	}

	stockQuery := `
		UPDATE products
		SET stock = stock + $1
		WHERE id = $2
	`

	_, err = conn.Exec(context.Background(), stockQuery, stockChange, t.ProductID)
	if err != nil {
		return nil, err
	}

	return &newTx, nil
}

func GetTransactionHistories() ([]TransactionHistoryResponse, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		log.Println("DB Connection Error:", err)
		return []TransactionHistoryResponse{}, err
	}
	defer conn.Close()

	query := `
	SELECT 
		t.id,                          
		p.id AS product_id,                          
		p.name AS product_name,
		c.name AS category_name,
		p.purchase_price,
		p.selling_price,
		SUM(CASE WHEN t.transaction_type = 'in' THEN t.quantity ELSE 0 END) AS total_in,
		SUM(CASE WHEN t.transaction_type = 'out' THEN t.quantity ELSE 0 END) AS total_out,
		p.stock AS stock_left,
		SUM(CASE WHEN t.transaction_type = 'in' THEN t.quantity * p.purchase_price ELSE 0 END) AS total_purchase,
		SUM(CASE WHEN t.transaction_type = 'out' THEN t.quantity * p.selling_price ELSE 0 END) AS total_sales
	FROM
		transactions t
	JOIN
		products p ON t.product_id = p.id
	JOIN
		category_products c ON p.category_id = c.id
	GROUP BY t.id, p.id, p.name, c.name, p.purchase_price, p.selling_price, p.stock
	ORDER BY MAX(t.created_at) DESC;
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Println("Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	histories, err := pgx.CollectRows[TransactionHistoryResponse](rows, pgx.RowToStructByName)
	if err != nil {
		log.Println("Mapping Error:", err)
		return []TransactionHistoryResponse{}, err
	}

	return histories, nil
}

func CheckStock(productID int) (int, error) {
	conn, err := utils.ConnectDB()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var stock int
	query := `SELECT stock FROM products WHERE id = $1`

	err = conn.QueryRow(context.Background(), query, productID).Scan(&stock)
	if err != nil {
		return 0, fmt.Errorf("failed to get product stock: %v", err)
	}

	log.Printf("Stock for product %d is %d", productID, stock)

	return stock, nil
}
