package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xhfmvls/restaurant-api/pkg/config"
)

var db *gorm.DB

type Food struct {
	gorm.Model
	Name  string  `gorm:"" json:"name"`
	Price float64 `json:"price"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Food{})
}

func (f *Food) AddFood() *Food {
	db.NewRecord(f)
	db.Create(&f)
	return f
}

func GetMenu() []Food {
	var Foods []Food
	db.Find(&Foods)
	return Foods
}
