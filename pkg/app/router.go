package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xhfmvls/restaurant-api/pkg/controllers"
	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
)

var NewRouter = func(router *mux.Router) {

	v1Router := router.PathPrefix("/v1").Subrouter()

	// POST new Food to Menu
	v1Router.HandleFunc("/menu", controllers.PostFood).Methods("POST")
	// GET All foods from Menu
	v1Router.Handle("/menu", middlewares.PriceFilter(middlewares.Search(middlewares.Sorting(middlewares.Pagination(http.HandlerFunc(controllers.GetMenu)))))).Methods("GET")
	// GET food from Menu
	v1Router.HandleFunc("/menu/{foodId}", controllers.GetFood).Methods("GET")
	// DELETE food from Menu
	v1Router.HandleFunc("/menu/{foodId}", controllers.DeleteFood).Methods("DELETE")
	// PUT (Update) food from Menu
	v1Router.HandleFunc("/menu/{foodId}", controllers.UpdateFood).Methods("PUT")

	// User Login
	v1Router.HandleFunc("/auth/login", controllers.Login).Methods("POST")
	// Register
	v1Router.HandleFunc("/auth/register", controllers.Register).Methods("POST")

	// Get User Information
	v1Router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetProfile))).Methods("GET")
	// Update User Information
	v1Router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateProfile))).Methods("PUT")
	// Delete Account
	v1Router.Handle("/user", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteAccount))).Methods("DELETE")

	// POST Food To User's Cart
	v1Router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.AddFood))).Methods("POST")
	// GET User's Cart
	v1Router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetCart))).Methods("GET")
	// DELETE User's Cart
	v1Router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.DeleteCart))).Methods("DELETE")
	// PUT User's Cart (Quantity only)
	v1Router.Handle("/cart", middlewares.AuthMiddleware(http.HandlerFunc(controllers.UpdateCartFood))).Methods("PUT")

	// POST (Create) Transaction
	v1Router.Handle("/transaction", middlewares.AuthMiddleware(http.HandlerFunc(controllers.CreateTransaction))).Methods("POST")
	// GET Transactions List
	v1Router.Handle("/transaction", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUserTransactions))).Methods("GET")
	// GET Transaction Detail
	v1Router.Handle("/transaction/{transactionId}", middlewares.AuthMiddleware(http.HandlerFunc(controllers.GetUserTransactionDetail))).Methods("GET")
}
