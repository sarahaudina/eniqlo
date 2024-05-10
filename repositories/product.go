package repositories

import (
 "database/sql"
 "log"

 "com.eniqlo/models"
 "time"
)

type ProductRepository struct {
 *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
 return ProductRepository{DB: db}
}

func (m *ProductRepository) CreateProduct(post models.CreateProduct) bool {
 stmt, err := m.DB.Prepare("INSERT INTO products (name, sku, image_url, notes, price, stock, location, is_available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Name, post.Sku, post.ImageUrl, post.Notes, post.Price, post.Stock, post.Location, post.IsAvailable)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *ProductRepository) FindProducts() []models.Product {
    var result []models.Product

    rows, err := m.DB.Query("SELECT * FROM products")

    if err != nil {
     log.Println(err)
     return nil
    }

    for rows.Next() {
        var (
         id       uint
         name    string
         created_at  time.Time
         updated_at time.Time
		 sku	string
		 image_url	string
		 notes	string
		 price	uint
		 stock	uint
		 location	string
		 is_available	bool
        )
        err := rows.Scan(&id, &created_at, &updated_at, &name, &sku, &image_url, &notes, &price, &stock, &location, &is_available)
        if err != nil {
         log.Println(err)
         return nil
        } else {
         product := models.Product{ID: id, CreatedAt: created_at, UpdatedAt: updated_at, Name: name, Sku: sku, ImageUrl: image_url, Notes: notes, Price: price, Stock: stock, Location: location, IsAvailable: is_available}
         result = append(result, product)
        }
    }

    return result
}

func (db *ProductRepository) UpdateProduct(product models.Product) error {
	_, err := db.Exec("UPDATE products SET name = ?, sku = ?, image_url = ?, notes = ?, price = ?, stock = ?, location = ?, is_available = ? WHERE id = ?",
	  product.Name, product.Sku, product.ImageUrl, product.Notes, product.Price, product.Stock, product.Location, product.IsAvailable, product.ID)
	if err != nil {
	  return err
	}
	return nil
}

func (db *ProductRepository) DeleteProduct(id uint) error {
	result, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
	  return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
	  return err
	}
	if rowsAffected == 0 {
	  return nil
	}
	return nil
}

func (m *ProductRepository) CheckoutProduct(id uint) error {
  return nil
}

func (m *ProductRepository) FindProduct(id uint) *models.Product {
  rows, err := m.DB.Query("SELECT * FROM products WHERE id = $1 LIMIT 1", id)

  if err != nil {
   log.Println(err)
   return nil
  }


  for rows.Next() {
    var (
      id       uint
      name    string
      created_at  time.Time
      updated_at time.Time
      sku	string
      image_url	string
      notes	string
      price	uint
      stock	uint
      location	string
      is_available	bool
    )
    err := rows.Scan(&id, &created_at, &updated_at, &name, &sku, &image_url, &notes, &price, &stock, &location, &is_available)
    if err != nil {
      log.Println(err)
      return nil
    } else {
      product := models.Product{ID: id, CreatedAt: created_at, UpdatedAt: updated_at, Name: name, Sku: sku, ImageUrl: image_url, Notes: notes, Price: price, Stock: stock, Location: location, IsAvailable: is_available}
      return &product
    } 
  }

  return nil
}

// todo: improve this by passing flexible conditions
func (m *ProductRepository) FindAvailableProducts() []models.Product {
  var result []models.Product

  rows, err := m.DB.Query("SELECT * FROM products WHERE is_available='true'")

  if err != nil {
   log.Println(err)
   return nil
  }

  for rows.Next() {
      var (
       id       uint
       name    string
       created_at  time.Time
       updated_at time.Time
   sku	string
   image_url	string
   notes	string
   price	uint
   stock	uint
   location	string
   is_available	bool
      )
      err := rows.Scan(&id, &created_at, &updated_at, &name, &sku, &image_url, &notes, &price, &stock, &location, &is_available)
      if err != nil {
       log.Println(err)
       return nil
      } else {
       product := models.Product{ID: id, CreatedAt: created_at, UpdatedAt: updated_at, Name: name, Sku: sku, ImageUrl: image_url, Notes: notes, Price: price, Stock: stock, Location: location, IsAvailable: is_available}
       result = append(result, product)
      }
  }

  return result
}