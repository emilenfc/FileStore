package main

import (
	"filestore/config"
	"filestore/handlers"
	"filestore/middleware"
	"filestore/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "filestore/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var uploadDir string
var baseURL string

// UploadResponse structure for the JSON response
type UploadResponse struct {
	Message  string `json:"message"`
	FileName string `json:"file_name"`
	Dir      string `json:"dir"`
	FileURL  string `json:"file_url"`
	FullPath string `json:"full_path"`
}

// UploadHandler
// @Summary Upload file
// @Description Upload a file to a specific folder
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param X-API-Key header string true "API Key for authentication"
// @Param X-API-Secret header string true "API Secret for authentication"
// @Param folder formData string true "Folder name" required
// @Param file formData file true "File to upload" required
// @Success 200 {object} UploadResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /upload [post]
func UploadHandler(c *gin.Context) {
	// Get API credentials from header
	apiKey := c.GetHeader("X-API-Key")
	apiSecret := c.GetHeader("X-API-Secret")

	if apiKey == "" || apiSecret == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "API key and secret key are required to upload files"})
		return
	}

	// Find user by API key
	var user models.User
	if err := config.DB.Where("api_key = ?", apiKey).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API credentials"})
		return
	}

	// Validate API secret
	if apiSecret != user.APISecret {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API credentials"})
		return
	}

	// Parse the multipart form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to process file upload"})
		return
	}

	// Get and validate folder name
	folderName := c.PostForm("folder")
	if folderName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Folder name is required"})
		return
	}

	// Clean folder name and convert to uppercase
	folderName = strings.ToUpper(filepath.Clean(folderName))

	// Check if folder exists or create new one
	var folder models.Folder
	result := config.DB.Where("name = ? AND user_id = ?", folderName, user.ID).First(&folder)
	if result.Error != nil {
		// Create new folder
		folder = models.Folder{
			Name:   folderName,
			UserID: user.ID,
		}
		if err := config.DB.Create(&folder).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating folder"})
			return
		}
	}

	fmt.Println("c.Request", c.Request)
	// Get file from form
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file"})
		return
	}
	defer file.Close()

	// Create user's directory path
	userPath := filepath.Join(uploadDir, user.APIKey, folderName)
	if err := os.MkdirAll(userPath, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create directory"})
		return
	}

	// Generate unique filename
	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), header.Filename)
	filePath := filepath.Join(userPath, fileName)

	// Save file
	destFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}

	// Save file record in database
	// fileRecord := models.File{
	// 	Name:     fileName,
	// 	Path:     filePath,
	// 	FolderID: folder.ID,
	// 	UserID:   user.ID,
	// }
	// if err := config.DB.Create(&fileRecord).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error recording file"})
	// 	return
	// }

	// Generate response
	fileURL := fmt.Sprintf("%s/uploads/%s/%s/%s", baseURL, user.APIKey, url.PathEscape(folderName), url.PathEscape(fileName))
	response := UploadResponse{
		Message:  "File uploaded successfully",
		FileName: fileName,
		Dir:      folderName, // Return original folder name without API key
		FileURL:  fileURL,
		FullPath: filePath,
	}

	c.JSON(http.StatusOK, response)
}

// FileStore handler
// @Summary Get file
// @Description Download a file
// @Tags files
// @Produce octet-stream
// @Param path path string true "File path"
// @Success 200 {file} binary "File contents"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /uploads/{path} [get]
func filestore(c *gin.Context) {
	// Extract path components
	path := c.Param("path")
	fmt.Println("path", path)

	// Decode the URL-encoded path
	decodedPath, err := url.QueryUnescape(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}

	fmt.Println("decodedPath", decodedPath)

	// Extract components after decoding
	components := strings.Split(decodedPath, "/")[1:] // spliting and Remove empty string at the beginning
	fmt.Println("components", components)

	if len(components) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file path"})
		return
	}
	apiKey := components[0]
	folderName := components[1]
	fileName := components[2]

	fmt.Println("API Key:", apiKey)
	// Verify user exists with this API key
	var user models.User
	if err := config.DB.Where("api_key = ?", apiKey).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid API key"})
		return
	}

	// Build file path
	filePath := filepath.Join(uploadDir, apiKey, folderName, fileName)
	fmt.Println("File Path:", filePath)
	// Verify file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Serve the file
	c.File(filePath)
}

// main.go handlers
// @title File Store API
// @version 1.0
// @description A secure file storage API with user authentication and folder management

// @contact.name API Support
// @contact.email iyaemile4@gmail.com
// @contact.phone +250783544364

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8085
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey TokenAuth
// @in header
// @name Authorization
// @description JWT token for authentication
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	config.InitDB()

	// Auto-migrate the database
	config.DB.AutoMigrate(&models.User{}, &models.Folder{}) //&models.File{}

	// Set configuration values from the environment
	uploadDir = os.Getenv("UPLOAD_DIR")
	baseURL = os.Getenv("BASE_URL")

	// Initialize Gin router
	r := gin.Default()

	url := ginSwagger.URL(baseURL + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// You can also add a direct redirect route
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key, X-API-Secret")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/user", handlers.GetUserInfo)
		api.POST("/regenerate-secret", handlers.RegenerateAPISecret)
		api.GET("/folders", handlers.GetUserFolders)
		api.GET("/folders/:folder", handlers.GetFolderContents)

	}
	// File serving route (public but requires valid path)
	r.POST("/upload", UploadHandler)
	r.GET("/uploads/*path", filestore)

	// Start server
	fmt.Printf("Server started at %s\n", baseURL)
	log.Fatal(r.Run(":8085"))

}

type ErrorResponse struct {
	Error string `json:"error" example:"Invalid credentials"`
}
