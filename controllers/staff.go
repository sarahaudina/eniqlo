package controllers

import (
	"database/sql"

	"com.eniqlo/models"
	"com.eniqlo/repositories"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"time"
)


type StaffController struct {
	*sql.DB
}
   
func NewStaffController(db *sql.DB) StaffController {
	return StaffController{DB: db}
}

func (m *StaffController) Register(c *gin.Context) {

	var input models.InputCreateStaff

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff := models.InputCreateStaff{
		Name: input.Name,
		PhoneNumber: input.PhoneNumber,
		Password: string(passwordHash),
	}

	db := m.DB
	repo := repositories.NewStaffRepository(db)

	insert := repo.CreateStaff(staff)

	if insert {
		c.JSON(200, gin.H{"status": "success", "msg": "Insert cat successfully"})
	} else {
		c.JSON(500, gin.H{"status": "failed", "msg": "Insert cat failed"})
	}
}

func (m *StaffController) Login(c *gin.Context) {

	var input models.InputLoginStaff

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := m.DB
	repo := repositories.NewStaffRepository(db)
	staff := repo.FindStaff(input.Name)

	if staff == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  staff.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}