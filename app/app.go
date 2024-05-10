package app

import (
 "database/sql"
 _ "github.com/lib/pq"
 "fmt"
 "log"

 "github.com/gin-gonic/gin"

 "com.eniqlo/controllers"
 "com.eniqlo/middlewares"
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

 // Staff
 staffController := controllers.NewStaffController(a.DB)
 routes.POST("/staff/register", staffController.Register)
 routes.POST("/staff/login", staffController.Login)
 
 // Customer, pass staff token here
 customerController := controllers.NewCustomerController(a.DB)
 routes.POST("/customer/register", middlewares.CheckAuth, customerController.Register)
 routes.GET("/customer", middlewares.CheckAuth, customerController.FindCustomers)

 // Products, pass staff token here
 productController := controllers.NewProductController(a.DB)
 routes.POST("/product", middlewares.CheckAuth, productController.AddProduct)
 routes.GET("/product", middlewares.CheckAuth, productController.FindProducts)
 routes.PUT("/product/:id", middlewares.CheckAuth, productController.UpdateProduct)
 routes.DELETE("/product/:id", middlewares.CheckAuth, productController.DeleteProduct)

 // Customer facing product endpoint
 routes.GET("/product/customer", productController.FindProducts)


 // checkouts. pass staff token here. todo: create checkout and get checkouts
 // routes.POST("/product/checkout", productController.CheckoutProduct)

 a.Routes = routes
}

func (a *App) Run(){
 a.Routes.Run(":8080")
}