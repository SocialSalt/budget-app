package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	Date            time.Time  `json:"date"`
	Company         string     `json:"company"`
	Category        string     `json:"category"`
	Amount          float64    `json:"amount"`
	AccountNumber   string     `json:"account_number"`
	Institution     string     `json:"institution"`
	Month           time.Month `json:"month"`
	Week            int        `json:"week"`
	CheckNumber     string     `json:"check_number"`
	FullDescription string     `json:"full_description"`
	DateAdded       time.Time  `json:"date_added"`
}

type Budget struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Category string     `json:"category"`
	Year     uint       `json:"year"`
	Month    time.Month `json:"month"`
	Amount   float64    `json:"amount"`
}

type Balance struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Date          time.Time  `json:"date"`
	Account       string     `json:"account"`
	AccountNumber string     `json:"account_number"`
	AccountID     string     `json:"account_id"`
	BalanceID     string     `json:"balance_id"`
	Institution   string     `json:"institution"`
	Month         time.Month `json:"month"`
	Week          int        `json:"week"`
	AccountType   string     `json:"account_type"`
	Class         string     `json:"class"`
	DateAdded     time.Time  `json:"date_added"`
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("budget-app.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Transaction{})
	database.AutoMigrate(&Budget{})
	database.AutoMigrate(&Balance{})

	DB = database
}
