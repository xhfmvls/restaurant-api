package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/xhfmvls/restaurant-api/pkg/config"
)

type User struct {
	gorm.Model
	Username     string `gorm:"" json:"username"`
	PasswordHash string `json:"passwordHash"`
}

type Credentials struct {
	Username string `gorm:"" json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Id int `gorm:"" json:"id"`
	jwt.StandardClaims
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (user *User) CreateUser() *User {
	db.NewRecord(user)
	db.Create(&user)
	return user
}

func GetUserById(id int) User {
	var user User
	db.Where("ID=?", id).Find(&user)
	return user
}

func GetUserByName(name string) User {
	var user User
	db.Where("Username=?", name).Find(&user)
	return user
}

func GetUserByNameAndPassword(username string, passwordHash string) User {
	var user User
	db.Where("Username=?", username).Where("Password_Hash=?", passwordHash).Find(&user)
	return user
}

func DeleteUser(id int) User {
	var deletedUser User
	db.Where("ID=?", id).Delete(&deletedUser)
	return deletedUser
}

func UpdateUser(newUser *User, id int) User {
	userDetails := GetUserById(id)
	if newUser.Username != "" {
		userDetails.Username = newUser.Username
	}
	if newUser.PasswordHash != "" {
		userDetails.PasswordHash = newUser.PasswordHash
	}
	db.Save(&userDetails)
	return userDetails
}
