package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xhfmvls/restaurant-api/pkg/config"
)

type Cart struct {
	gorm.Model
	UserId   int `json:"userid"`
	FoodId   int `json:"foodid"`
	Quantity int `json:"quantity"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Cart{})
}

func AddFoodToCart(userId int, foodId int, qty int) Cart {
	if foodId == 0 {
		panic("Food ID not provided")
	}
	if qty == 0 {
		qty = 1
	}
	cart := Cart{}
	cart.UserId = userId
	cart.FoodId = foodId
	cart.Quantity = qty
	db.NewRecord(cart)
	db.Create(&cart)
	return cart
}

func GetFoodFromCart(userId int) []Cart {
	var cart []Cart
	db.Where("User_Id=?", userId).Find(&cart)
	return cart
}