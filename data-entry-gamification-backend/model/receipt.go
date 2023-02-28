package model

import (
	"database/sql"
	"time"
)

type Receipt struct {
	ID        int64          `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Make      string         `json:"make"`
	ModelYear int            `json:"model_year"`
	State     string         `json:"state"`
	Vin       string         `json:"vin"`
	DateAdded time.Time      `json:"date_added"`
	QAScore   sql.NullInt64  `json:"qa_score"`
	QADate    sql.NullString `json:"qa_date"`
}

type ReceiptDTO struct {
	ID        int64          `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Make      string         `json:"make"`
	ModelYear int            `json:"model_year"`
	State     string         `json:"state"`
	Vin       string         `json:"vin"`
	DateAdded string         `json:"date_added"`
	QAScore   sql.NullInt64  `json:"qa_score"`
	QADate    sql.NullString `json:"qa_date"`
}
