package repositories

import (
 "database/sql"
 "log"

 "com.eniqlo/models"
 "time"
)

type StaffRepository struct {
 *sql.DB
}

func NewStaffRepository(db *sql.DB) StaffRepository {
 return StaffRepository{DB: db}
}

func (m *StaffRepository) CreateStaff(post models.InputCreateStaff) bool {
 stmt, err := m.DB.Prepare("INSERT INTO staffs (name, password, phone_number) VALUES ($1, $2, $3)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Name, post.Password, post.PhoneNumber)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *StaffRepository) FindStaff(name string) *models.Staff {
    rows, err := m.DB.Query("SELECT * FROM staffs WHERE name = $1 LIMIT 1", name)

    if err != nil {
     log.Println(err)
     return nil
    }

    for rows.Next() {
        var (
         id       uint
         name    string
		 password	string
         phone_number    string
         created_at  time.Time
         updated_at time.Time
        )
        err := rows.Scan(&id, &name, &password, &phone_number, &created_at, &updated_at)
        if err != nil {
         log.Println(err)
         return nil
        } else {
         staff := models.Staff{ID: id, Name: name, PhoneNumber: phone_number, CreatedAt: created_at, UpdatedAt: updated_at}
         return &staff
        }
    }

    return nil
}