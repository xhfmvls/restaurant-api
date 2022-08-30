package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	// "github.com/gorilla/sessions"

	// "strings"

	// "github.com/gorilla/mux"
	// "github.com/xhfmvls/restaurant-api/pkg/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/xhfmvls/restaurant-api/pkg/middlewares"
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
		Id: int(loginUser.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exparationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	tokenString = fmt.Sprintf("Bearer %s", tokenString)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: exparationTime,
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	resp, _ := json.Marshal(JwtResponse{
		Key: tokenString,
	})

	w.Header().Set("Access-Control-Allow-Credentials", "true")
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

	if exsistUser.ID != 0 {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	newUser := models.User{
		Username:     username,
		PasswordHash: password,
	}

	createdUser := newUser.CreateUser()
	res, _ := json.Marshal(createdUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.IdKey).(int)
	searchedUser := models.GetUserById(id)
	res, _ := json.Marshal(searchedUser)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.IdKey).(int)

	userCredentials := models.Credentials{}
	utils.ParseBody(r, &userCredentials)

	var passwordHash string
	if userCredentials.Password == "" {
		passwordHash = ""
	} else {
		passwordHash = utils.Sha256(userCredentials.Password) 
	}

	updatedUser := models.User{
		Username: userCredentials.Username,
		PasswordHash: passwordHash,
	}
	updatedUserDetails := models.UpdateUser(&updatedUser, id)
	res, _ := json.Marshal(updatedUserDetails)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}