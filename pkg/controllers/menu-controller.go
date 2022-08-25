package controllers

import(
	"net/http"
	"encoding/json"
	"github.com/xhfmvls/restaurant-api/pkg/models"
	"github.com/xhfmvls/restaurant-api/pkg/utils"
)

var NewFood models.Food

func PostFood(w http.ResponseWriter, r *http.Request) {
	newFood := &models.Food{}
	utils.ParseBody(r, newFood)
	food := newFood.AddFood()
	res, _ := json.Marshal(food)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	menu := models.GetMenu()
	res, _ := json.Marshal(menu)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
	
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	
}