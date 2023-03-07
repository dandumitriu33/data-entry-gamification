package model

import (
	"data-entry-gamification/utils/errors"
	"database/sql"
	"log"
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
	log.Println("Mapping from DTO to Model")
	receipt.ID = receiptDTO.ID
	receipt.ModelYear = receiptDTO.ModelYear
	receipt.Make = receiptDTO.Make
	receipt.Vin = receiptDTO.Vin
	receipt.FirstName = receiptDTO.FirstName
	receipt.LastName = receiptDTO.LastName
	receipt.State = receiptDTO.State
	log.Println("DateAdded:",receiptDTO.DateAdded)
	parsedDateAdded, parseErr := time.Parse("2006-01-02 03:04:00 +0000 UTC", receiptDTO.DateAdded)
	log.Println("ParsedDateAdded:",parsedDateAdded)
	if parseErr != nil {
		parseErrToDisplay := errors.NewBadRequestError("invalid DateAdded datetime format in DTO")
		log.Println(parseErrToDisplay)
		return parseErrToDisplay
	}
	receipt.DateAdded = parsedDateAdded
	log.Println("QAScore pre DTO map to Model:", receiptDTO.QAScore)
	receipt.QAScore = receiptDTO.QAScore
	parsedQADate := time.Time{}
	if receiptDTO.QADate != "" {
		parsedQADate, parseErr = time.Parse("2006-01-02 03:04:00 +0000 UTC", receiptDTO.QADate)
		if parseErr != nil {
			parseErrToDisplay := errors.NewBadRequestError("invalid QADate datetime format in DTO")
			log.Println(parseErrToDisplay)
			return parseErrToDisplay
		}
	}	
	receipt.QADate = parsedQADate
	return nil
}

func MapFromModelToDTO (receipt Receipt, receiptDTO *ReceiptDTO) {
	receiptDTO.ID = receipt.ID
	receiptDTO.ModelYear = receipt.ModelYear
	receiptDTO.Make = receipt.Make
	receiptDTO.Vin = receipt.Vin
	receiptDTO.FirstName = receipt.FirstName
	receiptDTO.LastName = receipt.LastName
	receiptDTO.State = receipt.State
	receiptDTO.DateAdded = receipt.DateAdded.String()
	receiptDTO.QAScore = receipt.QAScore
	receiptDTO.QADate = receipt.QADate.String()
}

func MapFromModelToDAO (receipt Receipt, receiptDAO *ReceiptDAO) {
	receiptDAO.ID = receipt.ID
	receiptDAO.ModelYear = receipt.ModelYear
	receiptDAO.Make = receipt.Make
	receiptDAO.Vin = receipt.Vin
	receiptDAO.FirstName = receipt.FirstName
	receiptDAO.LastName = receipt.LastName
	receiptDAO.State = receipt.State
	receiptDAO.DateAdded = receipt.DateAdded
	receiptDAO.QAScore = sql.NullInt64{ Int64: int64(receipt.QAScore), Valid: true}  
	receiptDAO.QADate = sql.NullTime{ Time: receipt.QADate, Valid: true} 
}
