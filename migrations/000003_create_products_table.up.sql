CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE,
  purchase_price INT NOT NULL,
  selling_price INT NOT NULL,
  stock INT NOT NULL,
  image_url VARCHAR(255) NOT NULL,
  category_id INT REFERENCES category_products(id) ON DELETE CASCADE ON UPDATE CASCADE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);