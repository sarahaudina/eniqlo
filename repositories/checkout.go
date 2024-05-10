package repositories

import (
 "database/sql"
//  "log"

//  "com.eniqlo/models"
)

type CheckoutRepository struct {
 *sql.DB
}

func NewCheckoutRepository(db *sql.DB) CheckoutRepository {
 return CheckoutRepository{DB: db}
}

// func (repo *CheckoutRepository) InsertCheckout(checkout models.CheckoutInput, tx *sql.Tx) (models.Checkout, error) {
// 	// Open connection if no transaction provided
// 	var db *sql.DB
// 	// if tx == nil {
// 	//   db, err := repo.DB.Conn(context.Background())
// 	//   if err != nil {
// 	// 	return checkout, err
// 	//   }
// 	//   defer db.Close()
// 	// } else {
// 	//   db = tx
// 	// }
  
// 	// Prepare insert statement for checkout
// 	stmt, err := repo.DB.Prepare(`
// 	  INSERT INTO checkout (customer_id, paid, change)
// 	  VALUES ($1, $2, $3)
// 	  RETURNING id, created_at, updated_at
// 	`)
// 	if err != nil {
// 	  return checkout, err
// 	}
// 	defer stmt.Close()
  
// 	// Execute insert statement for checkout
// 	result := db.Query(stmt, checkout.CustomerID, checkout.Paid, checkout.Change)
  
// 	// Scan results into checkout struct
// 	err = result.Scan(&checkout.ID, &checkout.CreatedAt, &checkout.UpdatedAt)
// 	if err != nil {
// 	  return checkout, err
// 	}
  
// 	// Insert product details (assuming separate table)
// 	for _, detail := range checkout.ProductDetails {
// 	  err = repo.InsertProductDetail(tx, checkout.ID, detail)
// 	  if err != nil {
// 		return checkout, err
// 	  }
// 	}
  
// 	return checkout, nil
//   }
  
//   func (repo *repository) InsertProductDetail(tx *sql.Tx, checkoutID uint, detail ProductDetail) error {
// 	// Open connection if no transaction provided
// 	var db *sql.DB
// 	if tx == nil {
// 	  db, err := repo.db.Conn(context.Background())
// 	  if err != nil {
// 		return err
// 	  }
// 	  defer db.Close()
// 	} else {
// 	  db = tx
// 	}
  
// 	// Prepare insert statement for product detail
// 	stmt, err := db.PrepareContext(context.Background(), `
// 	  INSERT INTO product_detail (checkout_id, product_id, quantity)
// 	  VALUES ($1, $2, $3)
// 	`)
// 	if err != nil {
// 	  return err
// 	}
// 	defer stmt.Close()
  
// 	// Execute insert statement for product detail
// 	_, err = db.ExecContext(context.Background(), stmt, checkoutID, detail.ProductID, detail.Quantity)
// 	if err != nil {
// 	  return err
// 	}
  
// 	return nil
//   }