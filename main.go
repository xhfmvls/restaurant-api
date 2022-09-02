package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/xhfmvls/restaurant-api/pkg/app"
)

func main() {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	godotenv.Load(".env")
	port, portEnvErr := strconv.Atoi(os.Getenv("PORT"))
	if portEnvErr != nil {
		panic("PORT not available")
	}
	app.NewRouter(r)
	http.Handle("/", r)
	fmt.Printf("Server Running on Port %d\n", port)
	address := fmt.Sprintf("localhost:%d", port)
	log.Fatal(http.ListenAndServe(address, r))
}
