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
	r := mux.NewRouter()
	godotenv.Load(".env")
	port, portEnvErr := strconv.Atoi(os.Getenv("PORT"))
	if portEnvErr != nil {
		panic("PORT not available")
	}
	fmt.Println(port)
	app.NewRouter(r)
	http.Handle("/api/v1/", r)
	address := fmt.Sprintf("localhost:%d", port)
	log.Fatal(http.ListenAndServe(address, r))
}
