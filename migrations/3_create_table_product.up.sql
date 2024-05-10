CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  sku VARCHAR(255) NOT NULL UNIQUE,  -- Unique identifier for product
  image_url VARCHAR(255),
  notes TEXT,  -- Allows longer text for notes
  price INTEGER NOT NULL,
  stock INTEGER NOT NULL,
  location VARCHAR(255),
  is_available BOOLEAN NOT NULL DEFAULT TRUE  -- Default product to available
);