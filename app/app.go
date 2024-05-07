package app

import (
 "database/sql"
 _ "github.com/lib/pq"
 "fmt"
 "log"

 "github.com/gin-gonic/gin"

 "com.eniqlo/controllers"
)

type App struct {
 DB     *sql.DB
 Routes *gin.Engine
}

func (a *App) CreateConnection(){
 connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", UNAMEDB, PASSDB, HOSTDB, DBNAME)
 db, err := sql.Open("postgres", connStr)
 if err != nil {
  log.Fatal(err)
 }
 a.DB = db
}

func (a *App) CreateRoutes() {
 routes := gin.Default()

 // auth
 customerController := controllers.NewCustomerController(a.DB)
 routes.GET("/customer", customerController.FindCustomer)
 routes.POST("/customer/register", customerController.CreateCustomer)

 a.Routes = routes
}

func (a *App) Run(){
 a.Routes.Run(":8080")
}