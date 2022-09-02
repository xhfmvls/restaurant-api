package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
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
	sortType := r.Context().Value(middlewares.SortKey).(string)
	limit := r.Context().Value(middlewares.LimitKey).(int)
	page := r.Context().Value(middlewares.PageKey).(int)
	searchParam := r.Context().Value(middlewares.SearchKey).(string)
	filterQuery := r.Context().Value(middlewares.PriceFilterKey).(string)
	menu := models.GetMenu(sortType, page, limit, searchParam, filterQuery)
	res, _ := json.Marshal(menu)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	foodId := vars["foodId"]
	id, err := strconv.ParseInt(foodId, 0, 0)
	if err != nil {
		panic("ID not valid")
	}
	foodDetails := models.GetFoodById(id)
	res, _ := json.Marshal(foodDetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	foodId := vars["foodId"]
	id, err := strconv.ParseInt(foodId, 0, 0)
	if err != nil {
		panic("ID not valid")
	}
	deletedFood := models.DeleteFoodById(id)
	res, _ := json.Marshal(deletedFood)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	updateFood := &models.Food{}
	utils.ParseBody(r, updateFood)

	vars := mux.Vars(r)
	foodId := vars["foodId"]
	id, err := strconv.ParseInt(foodId, 0, 0)
	if err != nil {
		panic("ID not valid")
	}
	newFoodDetail := models.UpdateFoodById(updateFood, id)
	res, _ := json.Marshal(newFoodDetail)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
