package repositories

import (
 "database/sql"
 "log"

 "com.eniqlo/models"
 "time"
)

type CustomerRepository struct {
 *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
 return CustomerRepository{DB: db}
}

func (m *CustomerRepository) CreateCustomer(post models.CreateCustomer) bool {
 stmt, err := m.DB.Prepare("INSERT INTO customers (name, phone_number) VALUES ($1, $2)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Name, post.PhoneNumber)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *CustomerRepository) FindCustomers() []models.Customer {
    var result []models.Customer

    rows, err := m.DB.Query("SELECT * FROM customers")

    if err != nil {
     log.Println(err)
     return nil
    }

    for rows.Next() {
        var (
         id       uint
         name    string
         phone_number    string
         created_at  time.Time
         updated_at time.Time
        )
        err := rows.Scan(&id, &name, &phone_number, &created_at, &updated_at)
        if err != nil {
         log.Println(err)
         return nil
        } else {
         cust := models.Customer{ID: id, Name: name, PhoneNumber: phone_number, CreatedAt: created_at, UpdatedAt: updated_at}
         result = append(result, cust)
        }
    }

    return result
}