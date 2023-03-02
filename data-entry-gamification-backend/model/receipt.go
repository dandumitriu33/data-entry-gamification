package model

import (
	"time"
)

type Receipt struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Make      string    `json:"make"`
	ModelYear int       `json:"model_year"`
	State     string    `json:"state"`
	Vin       string    `json:"vin"`
	DateAdded time.Time `json:"date_added"`
	QAScore   int       `json:"qa_score"`
	QADate    time.Time `json:"qa_date"`
}

type ReceiptDTO struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Make      string `json:"make"`
	ModelYear int    `json:"model_year"`
	State     string `json:"state"`
	Vin       string `json:"vin"`
	DateAdded string `json:"date_added"`
	QAScore   int    `json:"qa_score"`
	QADate    string `json:"qa_date"`
}
