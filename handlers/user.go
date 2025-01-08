package handlers

import (
	"filestore/config"
	"filestore/models"
	"filestore/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handler
// @Summary Register new user
// @Description Create a new user account
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration details"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	//Validate phone number using regex

	if !utils.IsValidPhone(user.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	// Generate API Key and Secret
	user.APIKey = utils.GenerateAPIKey()
	user.APISecret = utils.GenerateAPISecret()

	user.CreatedAt = time.Now()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	//save user to database

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})

}

// GetUserInfo handler
// @Summary Get user information
// @Description Get current user's profile information
// @Tags user
// @Produce json
// @Security TokenAuth
// @Success 200 {object} UserResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/user [get]
func GetUserInfo(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User
	if err := config.DB.Select("id, firstname, lastname, phone, email,api_key,api_secret, created_at").First(&user, userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// RegenerateAPISecret handler
// @Summary Regenerate API secret
// @Description Generate a new API secret for the current user
// @Tags user
// @Produce json
// @Security TokenAuth
// @Success 200 {object} SecretResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/regenerate-secret [post]
func RegenerateAPISecret(c *gin.Context) {
	userID := c.GetUint("user_id")
	newSecret := utils.GenerateAPISecret()

	if err := config.DB.Model(&models.User{}).Where("id = ?", userID).Update("api_secret", newSecret).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to regenerate API secret"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"api_secret": newSecret})
}

type RegisterResponse struct {
	Message string       `json:"message" example:"User created successfully"`
	User    UserResponse `json:"user"`
}

type RegisterRequest struct {
	FirstName string `json:"firstname" binding:"required" example:"John"`
	LastName  string `json:"lastname" binding:"required" example:"Doe"`
	Email     string `json:"email" binding:"required,email" example:"john@example.com"`
	Password  string `json:"password" binding:"required" example:"password123"`
	Phone     string `json:"phone" binding:"required" example:"+250783544364"`
}

type UserResponse struct {
	ID        uint      `json:"id" example:"1"`
	FirstName string    `json:"firstname" example:"John"`
	LastName  string    `json:"lastname" example:"Doe"`
	Email     string    `json:"email" example:"john@example.com"`
	Phone     string    `json:"phone" example:"+250783544364"`
	APIKey    string    `json:"api_key" example:"ak_123456789"`
	APISecret string    `json:"api_secret" example:"as_987654321"`
	CreatedAt time.Time `json:"created_at" example:"2025-01-08T12:00:00Z"`
}

type SecretResponse struct {
	APISecret string `json:"api_secret" example:"as_987654321"`
}
