# Go Backend Inventory

## How to Run this Project

1. Create a new empty directory for the project and navigate into it
2. Clone this project into the empty current directory:

```
git clone https://github.com/yusufbahtiarr/fgo24-phase3.git .
```

3. Install dependencies

```
go mod tidy
```

4. Run the project

```
go run main.go
```

## User Requirment

| No  | Feature                    | Description                                                                                                                                                                                                                                |
| --- | -------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| 1   | User Login                 | Users can log in using their credentials to access and manage inventory features. Authentication is required to protect access to stock and transaction data.                                                                              |
| 2   | View Product Categories    | Users can view a list of product categories, including the category name and description. seeding.                                                                                                                                         |
| 3   | View Product List          | Users can view all products including: product name, image (URL), category, purchase price, selling price, and available stock.                                                                                                            |
| 4   | Product Stock Transactions | Admin can record stock transactions, both incoming (stock in) and outgoing (stock out). Incoming transactions increase product stock, while outgoing transactions decrease it — ensuring the quantity does not exceed the available stock. |
| 5   | Transaction History        | Users can view all transaction history with details:transaction id, product id, product name, category, quantity in, quantity out, purchase price, selling price, total purchase value, total sales value, and current stock availability. |

## Flowchart

```mermaid
flowchart TD
    Mulai@{ shape: circle, label: "Mulai" }
    Login@{ shape: rect, label: "Tampilkan Menu Login" }
    ValidasiLogin@{ shape: diamond, label: "Validasi User?" }
    TampilMenu@{ shape: rect, label: "Tampilkan Menu" }
    PilihMenu@{ shape: diamond, label: "Pilih Menu?" }
    BarangKeluar@{ shape: rect, label: "Barang Keluar"}
    BarangMasuk@{ shape: rect, label: "Barang Masuk"}
    Category@{ shape: rect, label: "Category"}
    Product@{ shape: rect, label: "Product"}
    HistoryTransaction@{ shape: rect, label: "History Transaction"}
    InputBarangKuantitas@{ shape: lean-r, label: "Produk, Quantity" }
    InputBarangKuantitas2@{ shape: lean-r, label: "Produk, Quantity" }
    CekStok@{ shape: diamond, label: "Stok Produk?" }
    OutputBarangMasuk@{ shape: lean-r, label: '"Data Barang Masuk"'}
    OutputBarangKeluar@{ shape: lean-r, label: '"Data Barang Keluar"'}
    OutputCategory@{ shape: lean-r, label: '"Data Category"'}
    OutputProduct@{ shape: lean-r, label: '"Data Product"'}
    OutputHistoryTransaction@{ shape: lean-r, label: '"Data History Transaction"' }

    Selesai@{ shape: dbl-circ, label: "Selesai" }

    Mulai --> Login --> ValidasiLogin
    ValidasiLogin --Tidak--> Login
    ValidasiLogin --Ya--> TampilMenu
    TampilMenu --> PilihMenu
    PilihMenu --> BarangMasuk
    PilihMenu --> BarangKeluar
    PilihMenu --> Category
    PilihMenu --> Product
    PilihMenu --> HistoryTransaction
    BarangMasuk --> InputBarangKuantitas --> OutputBarangMasuk
    BarangKeluar --> InputBarangKuantitas2 --> CekStok
    Category --> OutputCategory
    Product --> OutputProduct
    HistoryTransaction --> OutputHistoryTransaction
    OutputBarangMasuk --> Selesai
    CekStok --Ya--> OutputBarangKeluar --> Selesai
    OutputCategory --> Selesai
    OutputProduct --> Selesai
    OutputHistoryTransaction --> Selesai
    CekStok --Tidak--> PilihMenu
    PilihMenu --> Selesai


```

### Entity Relationship Diagram

```mermaid
erDiagram
    users {
        int id PK
        string username
        string password
        date created_at
        date updated_at
    }
    products {
        int id PK
        string name
        int purchase_price
        int selling_price
        int stok
        string image_url
        int category_id FK
        date created_at
        date updated_at
    }
    category_products {
        int id PK
        string name
        string description
        date created_at
        date updated_at
    }
    transactions {
        int id PK
        int product_id FK
        string transaction_type
        int quantity
        int user_id FK
        date created_at
        date updated_at
    }
    users ||--o{ transactions : create
    products ||--o{ transactions : has
    products }o--|| category_products : has

```

## Endpoints Overview

| Method | Endpoint                | Description                                      | Auth |
| ------ | ----------------------- | ------------------------------------------------ | ---- |
| POST   | /users/login            | Login                                            | No   |
| GET    | /products               | Retrieves list of products                       | Yes  |
| GET    | /products/category      | Retrieves list of product categories             | Yes  |
| POST   | /inventory/transactions | Create a new transaction (stock-in or stock-out) | Yes  |
| GET    | /inventory/transactions | Retrieves transactions history                   | Yes  |

## Dependencies

- Gin Gonic : Used to build REST APIs, manage routes, handle HTTP requests/responses, and apply middleware.
- JWT (JSON Web Token) : Used for creating and verifying tokens for user authentication and authorization.
- Pgx (PostgreSQL driver) : Used to connect to a PostgreSQL database and run SQL queries.
- Godotenv : Loads environment variables from a .env file into the app.
- Argon2 : Used for securely hashing and verifying passwords.
