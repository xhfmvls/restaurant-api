package models

import (
	"github.com/jinzhu/gorm"
	"github.com/xhfmvls/restaurant-api/pkg/config"
)

type TransactionHeader struct {
	gorm.Model
	UserId         int  `json:"userid"`
	PurchaseStatus bool `json:"purchasestatus"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionId int `json:"transactionid"`
	FoodId        int `json:"foodid"`
	Quantity      int `json:"quantity"`
}

type TransactionInformation struct {
	TrHeader  TransactionHeader
	TrDetails []TransactionDetail
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&TransactionHeader{})
	db.AutoMigrate(&TransactionDetail{})
}

func (header *TransactionHeader) CreateTransactionHeader() *TransactionHeader {
	db.NewRecord(header)
	db.Create(&header)
	return header
}

func (detail *TransactionDetail) CreateTransactionDetail() *TransactionDetail {
	db.NewRecord(detail)
	db.Create(&detail)
	return detail
}

func CreateTransaction(userId int) TransactionInformation {
	FoodInCartList := GetFoodFromCart(userId)

	if len(FoodInCartList) == 0 {
		panic("No Item(s) in Cart")
	}

	header := TransactionHeader{}

	header.PurchaseStatus = true
	header.UserId = userId
	header.CreateTransactionHeader()

	var transactionDetails []TransactionDetail

	for _, food := range FoodInCartList {
		detail := TransactionDetail{}

		detail.TransactionId = int(header.ID)
		detail.FoodId = food.FoodId
		detail.Quantity = food.Quantity
		detail.CreateTransactionDetail()

		DeleteFoodFromCart(userId, detail.FoodId)

		transactionDetails = append(transactionDetails, detail)
	}

	transactionInfo := TransactionInformation{}
	transactionInfo.TrHeader = header
	transactionInfo.TrDetails = transactionDetails

	return transactionInfo
}

func GetTransactionsHeaderList(userId int) []TransactionHeader {
	var transactionList []TransactionHeader
	db.Where("User_Id=?", userId).Find(&transactionList)
	return transactionList
}

func GetTransactionHeader(transactionId int) TransactionHeader {
	transactionHeader := TransactionHeader{}
	db.Where("ID=?", transactionId).Find(&transactionHeader)
	return transactionHeader
}

func GetTransactionDetails(transactionId int) TransactionInformation {
	transactionHeader := GetTransactionHeader(transactionId)
	var transactionDetails []TransactionDetail
	db.Where("Transaction_Id=?", transactionId).Find(&transactionDetails)
	trnsactionInfo := TransactionInformation{
		TrHeader:  transactionHeader,
		TrDetails: transactionDetails,
	}
	return trnsactionInfo
}
