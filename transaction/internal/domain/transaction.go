package domain

import "time"

type Transaction struct {
	ID          string    `json:"id"`
	UserID      string    `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type TransactionSummary struct {
	IncomeInMonth     int           `json:"income_in_month"`
	ExpenseInMonth    int           `json:"expense_in_month"`
	LatestTransaction []Transaction `json:"latest_transaction"`
}

type TransactionDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Date        time.Time `json:"date"`
}
