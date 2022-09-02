package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
	"github.com/xhfmvls/restaurant-api/pkg/models"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	transaction := models.CreateTransaction(userId)
	res, _ := json.Marshal(transaction)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetUserTransactions(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value(middlewares.IdKey).(int)
	transactionsList := models.GetTransactionsHeaderList(userId)
	res, _ := json.Marshal(transactionsList)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserTransactionDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionId, err := strconv.Atoi(vars["transactionId"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	transactionInfo := models.GetTransactionDetails(transactionId)
	res, _ := json.Marshal(transactionInfo)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
