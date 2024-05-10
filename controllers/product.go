package controllers

import (
	"database/sql"

	"com.eniqlo/models"
	"com.eniqlo/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
	"strconv"
)

type ProductController struct {
	*sql.DB
}
   
func NewProductController(db *sql.DB) ProductController {
	return ProductController{DB: db}
}

func (m *ProductController) AddProduct(c *gin.Context) {

	var input models.CreateProduct

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := m.DB
	repo := repositories.NewProductRepository(db)

	insert := repo.CreateProduct(input)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "Insert cat successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "Insert cat failed"})
	}
}

func (m *ProductController) FindProducts(g *gin.Context) {
	db := m.DB
	repo := repositories.NewProductRepository(db)
	products := repo.FindProducts()
	if products != nil {
	 g.JSON(200, gin.H{"status": "success", "data": products, "msg": "get manga successfully"})
	} else {
	 g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}

func (m *ProductController) UpdateProduct(c *gin.Context) {
	var input models.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := m.DB
	repo := repositories.NewProductRepository(db)
	result := repo.UpdateProduct(input)

	if result != nil {
	 c.JSON(200, gin.H{"status": "success", "msg": "get manga successfully"})
	} else {
	 c.JSON(200, gin.H{"status": "success", "msg": "get manga successfully"})
	}
}

func (m *ProductController) DeleteProduct(g *gin.Context) {
	// Get the id
	path := g.Request.URL.Path
	parts := strings.Split(path, "/")
	_id := parts[3]
	__id, err := strconv.ParseUint(_id, 10, 32)
	if err != nil {
		g.JSON(200, gin.H{"status": "success", "msg": "delete manga successfully"})
	}

	id := uint(__id)
	db := m.DB
	repo := repositories.NewProductRepository(db)
	result := repo.DeleteProduct(id)

	if result != nil {
	 g.JSON(200, gin.H{"status": "success", "msg": "delete manga successfully"})
	} else {
	 g.JSON(200, gin.H{"status": "success", "msg": "deletev manga successfully"})
	}
}

// func (m *ProductController) CheckoutProduct(c *gin.Context) {
// 	var input models.CheckoutInput

// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	db := c.DB
// 	repo := repositories.NewCheckoutRepository(db)

// 	insert := repo.InsertCheckout(input)

// 	if insert {
// 		c.JSON(200, gin.H{"status": "success", "msg": "Insert cat successfully"})
// 	} else {
// 		c.JSON(500, gin.H{"status": "failed", "msg": "Insert cat failed"})
// 	}
// }


func (m *ProductController) FindAvailableProducts(g *gin.Context) {
	db := m.DB
	repo := repositories.NewProductRepository(db)
	products := repo.FindAvailableProducts()
	if products != nil {
	 g.JSON(200, gin.H{"status": "success", "data": products, "msg": "get manga successfully"})
	} else {
	 g.JSON(200, gin.H{"status": "success", "data": nil, "msg": "get manga successfully"})
	}
}