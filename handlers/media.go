package handlers

import (
	"filestore/config"
	"filestore/models"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserFolders(c *gin.Context) {
	userID := c.GetUint("user_id")
	var folders []models.Folder
	if err := config.DB.Where("user_id = ?", userID).Find(&folders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch folders"})
		return
	}
	c.JSON(http.StatusOK, folders)
}

// GetFolderContents for which is file saved in the folders
func GetFolderContents(c *gin.Context) {
	userID := c.GetUint("user_id")
	folderName := c.Param("folder")
	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// make sure folder name is in uppercase
	folderName = strings.ToUpper(filepath.Clean(folderName))

	// Build the directory path
	uploadDir := os.Getenv("UPLOAD_DIR")

	dirPath := filepath.Join(uploadDir, user.APIKey, folderName)

	// Check if directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Folder not found"})
		return
	}

	// Read directory contents
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading directory contents"})
		return
	}

	type FileResponse struct {
		Name      string    `json:"name"`
		URL       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
		Size      int64     `json:"size"`
	}

	var response []FileResponse
	baseURL := os.Getenv("BASE_URL")

	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip subdirectories
		}

		info, err := entry.Info()
		if err != nil {
			continue // Skip files with errors
		}

		// Extract original filename (remove timestamp prefix)
		originalName := strings.Join(strings.Split(entry.Name(), "_")[1:], "_")

		fileURL := fmt.Sprintf("%s/uploads/%s/%s/%s",
			baseURL,
			user.APIKey,
			url.PathEscape(folderName),
			url.PathEscape(entry.Name()))

		response = append(response, FileResponse{
			Name:      originalName,
			URL:       fileURL,
			CreatedAt: info.ModTime(),
			Size:      info.Size(),
		})
	}

	c.JSON(http.StatusOK, response)
}