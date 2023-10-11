package domain

import "time"

type Transaction struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"-" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Amount      int       `json:"amount" db:"amount"`
	Date        time.Time `json:"date" db:"date"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
	TotalData   int       `json:"-" db:"total_data"`
}

type TransactionSummary struct {
	IncomeInMonth     int           `json:"income_in_month"`
	ExpenseInMonth    int           `json:"expense_in_month"`
	LatestTransaction []Transaction `json:"latest_transaction"`
}

type TransactionDTO struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description"`
	Amount      int    `json:"amount" validate:"required,min=1"`
	Date        string `json:"date" validate:"required,ISO8601date"`
}
