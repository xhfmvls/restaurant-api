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

func (food *Food) AddFood() *Food {
	db.NewRecord(food)
	db.Create(&food)
	return food
}

func GetMenu(sortType string, page int, limit int, searchName string) []Food {
	offset := (page - 1) * limit
	name := searchName
	var Foods []Food
	db.Where(Food{Name: name}).Order(sortType).Offset(offset).Limit(limit).Find(&Foods)
	return Foods
}

func GetFoodById(id int64) *Food {
	var searchedFood Food
	db.Where("ID=?", id).Find(&searchedFood)
	return &searchedFood
}

func DeleteFoodById(id int64) *Food {
	var deletedFood Food
	db.Where("ID=?", id).Delete(&deletedFood)
	return &deletedFood
}

func UpdateFoodById(food *Food, id int64) *Food {
	foodDetails := GetFoodById(id)
	if food.Name != "" {
		foodDetails.Name = food.Name
	}
	if food.Price != 0 {
		foodDetails.Price = food.Price
	}
	db.Save(&foodDetails)
	return foodDetails
}