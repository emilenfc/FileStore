package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string   `json:"firstname" binding:"required"`
	Lastname  string   `json:"lastname" binding:"required"`
	Email     string   `json:"email" gorm:"unique" binding:"required,email"`
	Password  string   `json:"password" binding:"required,min=5"`
	Phone     string   `json:"phone" binding:"required"`
	APIKey    string   `json:"api_key" gorm:"unique"`
	APISecret string   `json:"api_secret" gorm:"unique"`
	Folders   []Folder `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}

type Folder struct {
	gorm.Model
	Name   string `json:"name"`
	UserID uint   `json:"user_id" `
	// Files []File `gorm:"foreignKey:FolderID"`
	CreatedAt time.Time
}

// type File struct {
// 	gorm.Model
// 	Name string `json:"name"`
// 	Path string `json:"path"`
// 	FolderID uint `json:"folder_id" `
// 	UserID uint `json:"user_id" `
// 	CreatedAt time.Time
// }
