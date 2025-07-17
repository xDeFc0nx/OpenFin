package types

import (
	"time"
)

type WebSocket struct {
	ID           string    `db:"id" json:"ID"`
	ConnectionID string    `db:"connection_id" json:"ConnectionID"`
	UserID       string    `db:"user_id" json:"UserID"`
	IsActive     bool      `db:"is_active" json:"IsActive"`
	LastPing     time.Time `db:"last_ping" json:"LastPing"`
	CreatedAt    time.Time `db:"created_at" json:"CreatedAt"`
}

type User struct {
	ID        string    `db:"id" json:"ID"`
	FirstName string    `db:"first_name" json:"FirstName"`
	LastName  string    `db:"last_name" json:"LastName"`
	Email     string    `db:"email" json:"Email"`
	Password  string    `db:"password" json:"-"`
	Currency  string    `db:"currency" json:"Currency"`
	CreatedAt time.Time `db:"created_at" json:"CreatedAt"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
}

type Accounts struct {
	Income    float64    `db:"income" json:"Income"`
	Expense   float64    `db:"expense" json:"Expense"`
	Balance   float64    `db:"balance" json:"Balance"`
	ID        string     `db:"id" json:"ID"`
	UserID    string     `db:"user_id" json:"UserID"`
	Type      string     `db:"type" json:"Type"`
	CreatedAt *time.Time `db:"created_at" json:"CreatedAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"UpdatedAt"`
}

type Transaction struct {
	Amount      float64   `db:"amount" json:"Amount"`
	ID          string    `db:"id" json:"ID"`
	UserID      string    `db:"user_id" json:"UserID"`
	AccountID   string    `db:"account_id" json:"AccountID"`
	Type        string    `db:"type" json:"Type"`
	Description string    `db:"description" json:"Description"`
	IsRecurring bool      `db:"is_recurring" json:"IsRecurring"`
	CreatedAt   time.Time `db:"created_at" json:"CreatedAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
}

type Recurring struct {
	Amount        float64    `db:"amount" json:"Amount"`
	ID            string     `db:"id" json:"ID"`
	TransactionID string     `db:"transaction_id" json:"TransactionID"`
	Frequency     string     `db:"frequency" json:"Frequency"`
	StartDate     time.Time  `db:"start_date" json:"StartDate"`
	NextDate      time.Time  `db:"next_date" json:"NextDate"`
	EndDate       *time.Time `db:"end_date" json:"EndDate,omitempty"`
	CreatedAt     time.Time  `db:"created_at" json:"CreatedAt"`
	UpdatedAt     time.Time  `db:"updated_at" json:"UpdatedAt"`
}

type Budget struct {
	TotalSpent  float64   `db:"total_spent" json:"TotalSpent"`
	LimitAmount float64   `db:"limit_amount" json:"LimitAmount"`
	ID          string    `db:"id" json:"ID"`
	UserID      string    `db:"user_id" json:"UserID"`
	AccountID   string    `db:"account_id" json:"AccountID"`
	Description string    `db:"description" json:"Description"`
	CreatedAt   time.Time `db:"created_at" json:"CreatedAt"`
	UpdatedAt   time.Time `db:"updated_at" json:"-"`
}

type Goal struct {
	Amount        float64 `db:"amount" json:"Amount"`
	GoalAmount    float64 `db:"goal_amount" json:"GoalAmount"`
	CurrentAmount float64 `db:"current_amount" json:"CurrentAmount"`
	ID            string  `db:"id" json:"ID"`
	UserID        string  `db:"user_id" json:"UserID"`
	AccountID     string  `db:"account_id" json:"AccountID"`
	Description   string  `db:"description" json:"Description"`
}
