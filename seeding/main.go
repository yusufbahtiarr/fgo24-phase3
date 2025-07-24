package main

import (
	"context"
	"fmt"
	"go-test/utils"
	"log"
)

func main() {
	conn, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect: ")
	}
	defer conn.Close()

	_, err = conn.Exec(context.Background(), `INSERT INTO category_products (name, description) VALUES
		('Elektronik', 'Produk seperti HP, laptop, TV, dll'),
		('Pakaian', 'Baju, celana, jaket, dan sejenisnya'),
		('Aksesoris', 'Barang pelengkap seperti jam tangan, tas, topi')
	ON CONFLICT DO NOTHING; `)
	if err != nil {
		log.Fatalf("failed insert category_products: %v", err)
	}

	_, err = conn.Exec(context.Background(), `
		INSERT INTO products (name, purchase_price, selling_price, stock, image_url, category_id) VALUES
			('Smartphone XYZ', 2500000, 3500000, 15, 'smartphone.jpg', 1),
			('Laptop ABC', 6000000, 7500000, 10, 'laptop.jpg', 1),
			('Celana Jeans', 100000, 150000, 25, 'jeans.jpg', 2),
			('Jaket Hoodie', 120000, 175000, 20, 'hoodie.jpg', 2),
			('Jam Tangan', 50000, 85000, 30, 'jam.jpg', 3),
			('Tas Selempang', 40000, 70000, 20, 'tas.jpg', 3)
		ON CONFLICT DO NOTHING;
	`)
	if err != nil {
		log.Fatalf("failed insert products: %v", err)
	}

	adminHash := utils.HashPassword("admin123")
	yusufHash := utils.HashPassword("yusuf123")

	_, err = conn.Exec(context.Background(), `
		INSERT INTO users (username, password) VALUES
			('admin', $1),
			('yusuf', $2)
		ON CONFLICT DO NOTHING;
	`, adminHash, yusufHash)

	if err != nil {
		log.Fatalf("Gagal insert users: %v", err)
	}

	fmt.Println("Seeding berhasil!")

}
