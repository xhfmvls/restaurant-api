package models 

import (
	"github.com/jinzhu/gorm"
	"github.com/xhfmvls/restaurant-api/pkg/config"
)

type TransactionHeader struct {
	gorm.Model
	UserId int `json:"userid"`
	PurchaseStatus bool `json:"purchasestatus"`
}

type TransactionDetail struct {
	gorm.Model
	TransactionId int `json:"transactionid"`
	FoodId int `json:"foodid"`
	Quantity int `json:"quantity"`
}

type TransactionInformation struct {
	TrHeader TransactionHeader
	TrDetails []TransactionDetail
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&TransactionHeader{})
	db.AutoMigrate(&TransactionDetail{})
}

func (header *TransactionHeader)CreateTransactionHeader() *TransactionHeader {
	db.NewRecord(header)
	db.Create(&header)
	return header
}

func (detail *TransactionDetail)CreateTransactionDetail() *TransactionDetail {
	db.NewRecord(detail)
	db.Create(&detail)
	return detail
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
		TrHeader: transactionHeader,
		TrDetails: transactionDetails,
	}
	return trnsactionInfo
}

func confirmPurchase(transactionId int) /**TransactionHeader*/ {
	// header.PurchaseStatus = true
	println("")
}