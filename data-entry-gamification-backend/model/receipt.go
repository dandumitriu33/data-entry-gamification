package model

import (
	"database/sql"
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

type ReceiptDAO struct {
	ID        int64
	FirstName string
	LastName  string
	Make      string
	ModelYear int
	State     string
	Vin       string
	DateAdded time.Time
	QAScore   sql.NullInt64
	QADate    sql.NullTime
}

func MapFromDAOToModel(receiptDAO ReceiptDAO, receipt *Receipt) {
	receipt.ID = receiptDAO.ID
	receipt.ModelYear = receiptDAO.ModelYear
	receipt.Make = receiptDAO.Make
	receipt.Vin = receiptDAO.Vin
	receipt.FirstName = receiptDAO.FirstName
	receipt.LastName = receiptDAO.LastName
	receipt.State = receiptDAO.State
	receipt.DateAdded = receiptDAO.DateAdded
	receipt.QAScore = int(receiptDAO.QAScore.Int64)	
	receipt.QADate = receiptDAO.QADate.Time
}
