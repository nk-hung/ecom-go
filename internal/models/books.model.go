package models

import (
	"github.com/jinzhu/gorm"
	config "github.com/nk-hung/ecom-go/internal/configs"
)

var db *gorm.DB

type Book struct {
	// gorm.model  `gorm:"json:model"`
	Name        string `gorm:"json:name"`
	Auther      string `json:"auth"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
