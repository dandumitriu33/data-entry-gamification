package model

import (
	"data-entry-gamification/utils/errors"
	"database/sql"
	"time"
)

type Receipt struct {
	ID        int64
	FirstName string
	LastName  string
	Make      string
	ModelYear int
	State     string
	Vin       string
	DateAdded time.Time
	QAScore   int
	QADate    time.Time
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

func MapFromDTOToModel(receiptDTO ReceiptDTO, receipt *Receipt) *errors.RestErr {
	receipt.ID = receiptDTO.ID
	receipt.ModelYear = receiptDTO.ModelYear
	receipt.Make = receiptDTO.Make
	receipt.Vin = receiptDTO.Vin
	receipt.FirstName = receiptDTO.FirstName
	receipt.LastName = receiptDTO.LastName
	receipt.State = receiptDTO.State
	parsedDateAdded, parseErr := time.Parse(time.RFC3339, receiptDTO.DateAdded)
	if parseErr != nil {
		parseErrToDisplay := errors.NewBadRequestError("invalid DateAdded datetime format in DTO")
		return parseErrToDisplay
	}
	receipt.DateAdded = parsedDateAdded
	receipt.QAScore = receiptDTO.QAScore
	parsedQADate := time.Time{}
	if receiptDTO.QADate != "" {
		parsedQADate, parseErr = time.Parse(time.RFC3339, receiptDTO.QADate)
		if parseErr != nil {
			parseErrToDisplay := errors.NewBadRequestError("invalid QADate datetime format in DTO")
			return parseErrToDisplay
		}
	}	
	receipt.QADate = parsedQADate
	return nil
}
