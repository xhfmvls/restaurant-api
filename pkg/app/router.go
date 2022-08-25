package app

import (
	"github.com/gorilla/mux"
	"github.com/xhfmvls/restaurant-api/pkg/controllers"
)

var NewRouter = func(router *mux.Router) {
	// POST new Food to Menu
	router.HandleFunc("/menu", controllers.PostFood).Methods("POST")
	// GET All foods from Menu
	router.HandleFunc("/menu", controllers.GetMenu).Methods("GET")
	// GET food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.GetFood).Methods("GET")
	// DELETE food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.DeleteFood).Methods("DELETE")
	// PUT (Update) food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.UpdateFood).Methods("PUT")
}
