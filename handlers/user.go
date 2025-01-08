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

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func GetUserInfo(c *gin.Context) {
	userID := c.GetUint("user_id")
	var user models.User
	if err := config.DB.Select("id, firstname, lastname, phone, email,api_key,api_secret, created_at").First(&user, userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func RegenerateAPISecret(c *gin.Context) {
	userID := c.GetUint("user_id")
	newSecret := utils.GenerateAPISecret()

	if err := config.DB.Model(&models.User{}).Where("id = ?", userID).Update("api_secret", newSecret).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to regenerate API secret"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"api_secret": newSecret})
}
