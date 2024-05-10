CREATE TABLE IF NOT EXISTS checkout_product_details (
  id SERIAL PRIMARY KEY,
  checkout_id INTEGER REFERENCES checkout(id) NOT NULL,
  product_id VARCHAR(255) NOT NULL,
  quantity INTEGER CHECK (quantity > 0) NOT NULL
);