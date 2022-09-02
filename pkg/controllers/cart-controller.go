package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
	"github.com/xhfmvls/restaurant-api/pkg/models"
	"github.com/xhfmvls/restaurant-api/pkg/utils"
)

type CartInput struct {
	FoodId int `json:"foodid"`
	Qty    int `json:"quantity"`
}

func AddFood(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	addedFoodInput := CartInput{}
	utils.ParseBody(r, &addedFoodInput)
	foodId := addedFoodInput.FoodId
	qty := addedFoodInput.Qty
	cart, err := models.AddFoodToCart(userId, foodId, qty)
	if err == 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, _ := json.Marshal(cart)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	foodList := models.GetFoodFromCart(userId)
	res, _ := json.Marshal(foodList)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	deleteCartInput := CartInput{}
	utils.ParseBody(r, &deleteCartInput)
	foodId := deleteCartInput.FoodId
	deletedCart := models.DeleteFoodFromCart(userId, foodId)
	res, _ := json.Marshal(deletedCart)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCartFood(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	updatedFoodInput := CartInput{}
	utils.ParseBody(r, &updatedFoodInput)
	qty := updatedFoodInput.Qty
	foodId := updatedFoodInput.FoodId
	if qty == 0 || foodId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedCart := models.UpdateFoodQuantity(userId, foodId, qty)
	res, _ := json.Marshal(updatedCart)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
