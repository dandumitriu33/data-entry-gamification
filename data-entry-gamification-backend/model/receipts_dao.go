package model

import (
	"data-entry-gamification/storage/receipts_db"
	"data-entry-gamification/utils/errors"
	"time"
)

var (
	queryInsertReceipt = "INSERT INTO receipts (model_year, make, vin, first_name, last_name, state, date_added) VALUES (?, ?, ?, ?, ?, ?, ?);"
	queryGetAllCount = "SELECT COUNT(*) FROM receipts;"
)

func (receipt *Receipt) Save() *errors.RestErr {
	stmt, err := receipts_db.Client.Prepare(queryInsertReceipt)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	currentTime := time.Now()
	date_added := currentTime.Format("20060102150405")
	insertResult, saveErr := stmt.Exec(receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, date_added)
	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	receiptID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	receipt.ID = receiptID
	return nil
}

func (receipt *Receipt) GetAllCount() (int64, *errors.RestErr) {
	var count int64	

	row := receipts_db.Client.QueryRow(queryGetAllCount)
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.NewInternalServerError("database error")
	}

	return count, nil
}