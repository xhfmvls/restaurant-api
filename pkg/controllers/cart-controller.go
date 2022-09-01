package controllers

import (
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
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
	cart := models.AddFoodToCart(userId, foodId, qty)
	res, _ := json.Marshal(cart)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

// func UpdateFood(w http.ResponseWriter, r *http.Request) {
// 	userId :=  r.Context().Value(middlewares.IdKey).(int)
// 	updatedFoodInput := CartInput{}
// 	utils.ParseBody(r, &addedFoodInput)
// }
