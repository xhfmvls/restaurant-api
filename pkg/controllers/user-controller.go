package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	// "strings"

	// "github.com/gorilla/mux"
	// "github.com/xhfmvls/restaurant-api/pkg/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/xhfmvls/restaurant-api/pkg/models"
	"github.com/xhfmvls/restaurant-api/pkg/utils"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JwtResponse struct {
	Key string `json:"key"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	credentials := &models.Credentials{}
	utils.ParseBody(r, credentials)
	username, password := credentials.Username, credentials.Password
	password = utils.Sha256(password)

	loginUser := models.GetUserByNameAndPassword(username, password)

	if loginUser.ID == 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	exparationTime := time.Now().Add(time.Minute * 5)

	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exparationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: exparationTime,
	})
	resp, _ := json.Marshal(JwtResponse{
		Key: tokenString,
	})
	
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func Register(w http.ResponseWriter, r *http.Request) {
	credentials := models.Credentials{}

	utils.ParseBody(r, &credentials)
	username, password := credentials.Username, credentials.Password
	password = utils.Sha256(password)

	if username == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return 
	}

	exsistUser := models.GetUserByName(username)

	println(exsistUser.Username)

	if exsistUser.ID != 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	newUser := models.User{
		Username: username,
		PasswordHash: password,
	}

	createdUser := newUser.CreateUser()
	res, _ := json.Marshal(createdUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
