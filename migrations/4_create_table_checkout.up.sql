CREATE TABLE IF NOT EXISTS checkouts (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  product_id INTEGER REFERENCES products(id) ON DELETE CASCADE,
  customer_id INTEGER REFERENCES customers(id) ON DELETE CASCADE,
  quantity INTEGER NOT NULL,
  paid INTEGER NOT NULL,
  "change" INTEGER NOT NULL -- Double quotes are used to avoid conflict with a reserved keyword  
);