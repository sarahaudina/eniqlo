package controllers

import (
	"database/sql"

	"com.eniqlo/models"
	"com.eniqlo/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)


type CustomerController struct {
	*sql.DB
}
   
func NewCustomerController(db *sql.DB) CustomerController {
	return CustomerController{DB: db}
}

func (m *CustomerController) FindCustomer(g *gin.Context) {
	db := m.DB
	repo_cust := repositories.NewCustomerRepository(db)
	find_cust := repo_cust.FindCustomer("sarah")
	if find_cust != nil {
	 g.JSON(200, gin.H{"status": "success", "data": find_cust, "msg": "get manga successfully"})
	} else {
	 g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}

func (m *CustomerController) CreateCustomer(c *gin.Context) {

	var input models.CreateCustomer

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.CreateCustomer{
		Name: input.Name,
		PhoneNumber: input.PhoneNumber,
	}

	db := m.DB
	repo_cust := repositories.NewCustomerRepository(db)

	insert := repo_cust.CreateCustomer(customer)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "Insert cat successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "Insert cat failed"})
	}
}