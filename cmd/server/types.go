package server

import (
	"time"
)

// transaction types
type Transaction struct {
	ID              int64     `json:"id"`
	Date            time.Time `json:"date"`
	Company         string    `json:"company"`
	Category        string    `json:"category"`
	Amount          int64     `json:"amount"`
	AccountNumber   string    `json:"account_number"`
	Institution     string    `json:"institution"`
	FullDescription string    `json:"full_description"`
	DateAdded       time.Time `json:"date_added"`
}

type TransactionFilter struct {
	Category      *string
	StartDate     *time.Time
	EndDate       *time.Time
	AccountNumber *string
	MinAmount     *int64
	MaxAmount     *int64
	Group         *string
}

// budget types
type Budget struct {
	ID       int64      `json:"id"`
	Category string     `json:"category"`
	Group    string     `json:"group"`
	Year     uint       `json:"year"`
	Month    time.Month `json:"month"`
	Amount   int64      `json:"amount"`
}

type BudgetFilter struct {
	Category *string
	Group    *string
	Year     *uint
	Month    *time.Month
}

// balance types
type Balance struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Date          time.Time `json:"date"`
	Account       string    `json:"account"`
	AccountNumber string    `json:"account_number"`
	AccountID     string    `json:"account_id"`
	BalanceID     string    `json:"balance_id"`
	Institution   string    `json:"institution"`
	AccountType   string    `json:"account_type"`
	Class         string    `json:"class"`
	DateAdded     time.Time `json:"date_added"`
}

type BalanceFilter struct {
	StartDate     *time.Time
	EndDate       *time.Time
	AccountNumber *string
	Class         *string
}
