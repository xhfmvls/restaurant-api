package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xhfmvls/restaurant-api/pkg/controllers"
	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
)

var NewRouter = func(router *mux.Router) {
	// POST new Food to Menu
	router.HandleFunc("/menu", controllers.PostFood).Methods("POST")
	// GET All foods from Menu
	router.Handle("/menu", middlewares.PriceFilter(middlewares.Search(middlewares.Sorting(middlewares.Pagination(http.HandlerFunc(controllers.GetMenu)))))).Methods("GET")
	// GET food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.GetFood).Methods("GET")
	// DELETE food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.DeleteFood).Methods("DELETE")
	// PUT (Update) food from Menu
	router.HandleFunc("/menu/{foodId}", controllers.UpdateFood).Methods("PUT")

	// User Login
	router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	// Register
	router.HandleFunc("/auth/register", controllers.Register).Methods("POST")

	// Get User Information
	router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
	// Update User Information
	router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateProfile))).Methods("PUT")
	// Delete Account
	router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteAccount))).Methods("DELETE")

	// POST Food To User's Cart
	router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddFood))).Methods("POST")
	// GET User's Cart
	router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCart))).Methods("GET")
	// Delete User's Cart
	router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteCart))).Methods("DELETE")
}
