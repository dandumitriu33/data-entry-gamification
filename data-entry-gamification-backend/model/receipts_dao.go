package model

import (
	"context"
	"data-entry-gamification/storage/receipts_db"
	"data-entry-gamification/utils/errors"
	"log"
	"time"
)

var (
	queryInsertReceipt         = "INSERT INTO receipts (model_year, make, vin, first_name, last_name, state, date_added) VALUES (?, ?, ?, ?, ?, ?, ?);"
	queryGetAllCount           = "SELECT COUNT(*) FROM receipts;"
	queryGetAllCountToday      = "SELECT COUNT(*) FROM receipts WHERE DATE(date_added) = DATE(NOW());"
	queryInsertUserIDReceiptID = "INSERT INTO user_receipts (user_id, receipt_id) VALUES (?, ?);"
	queryGetUserPointsByUserID = "SELECT points FROM user_info WHERE user_id = ?;"
	queryUpdateUserPoints      = "UPDATE user_info SET points = ?, level = ? WHERE user_id = ?;"
	// queryGetLatestUnverifiedReceipt = "SELECT id, model_year, make, vin, first_name, last_name, state, date_added, qa_score, qa_date FROM receipts WHERE qa_score IS NULL ORDER BY date_added DESC LIMIT 1;"
	queryGetLatestUnverifiedReceipt = "SELECT id, model_year, make, vin, first_name, last_name, state, date_added, qa_score FROM receipts WHERE qa_score IS NULL ORDER BY date_added DESC LIMIT 1;"
	queryUpdateReceipt = "UPDATE receipts SET model_year = ?, make = ?, vin = ?, first_name = ?, last_name = ?, state = ?, qa_score = ?, qa_date = ? WHERE id = ?;"
)

func (receipt *Receipt) Save(ctx context.Context, userID int64) *errors.RestErr {
	tx, err := receipts_db.Client.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewInternalServerError("database transaction error")
	}
	defer tx.Rollback()

	// Add receipt and get the ID
	addResult, addErr := tx.ExecContext(ctx, queryInsertReceipt, receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, receipt.DateAdded)
	if addErr != nil {
		return errors.NewInternalServerError("database transaction add error")
	}
	receiptID, err := addResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database transaction receipt ID retrieval error")
	}
	receipt.ID = receiptID

	// Add userID and receiptID to user_receipts table
	_, pairErr := tx.ExecContext(ctx, queryInsertUserIDReceiptID, userID, receiptID)
	if pairErr != nil {
		return errors.NewInternalServerError("database transaction pair error")
	}

	// Get User points
	var userPoints int64
	getPointsErr := tx.QueryRowContext(ctx, queryGetUserPointsByUserID, userID).Scan(&userPoints)
	if getPointsErr != nil {
		return errors.NewInternalServerError("database transaction get points error")
	}

	// Update User Points
	_, pointsErr := tx.ExecContext(ctx, queryUpdateUserPoints, userPoints+1, 0, userID)
	if pointsErr != nil {
		return errors.NewInternalServerError("database transaction add points error")
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return errors.NewInternalServerError("database transaction commit error")
	}
	return nil
}

// func (receipt *Receipt) Save2(ctx context.Context, userID int64) *errors.RestErr {
// 	tx, err := receipts_db.Client.BeginTx(ctx, nil)
// 	if err != nil {
// 		return errors.NewInternalServerError("database transaction error")
// 	}
// 	defer tx.Rollback()

// 	// Add receipt and get the ID
// 	currentTime := time.Now()
// 	date_added := currentTime.Format("20060102150405")
// 	addResult, addErr := tx.ExecContext(ctx, queryInsertReceipt, receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, date_added)
// 	if addErr != nil {
// 		return errors.NewInternalServerError("database transaction add error")
// 	}
// 	receiptID, err := addResult.LastInsertId()
// 	if err != nil {
// 		return errors.NewInternalServerError("database transaction receipt ID retrieval error")
// 	}
// 	receipt.ID = receiptID

// 	// Add userID and receiptID to user_receipts table
// 	_, pairErr := tx.ExecContext(ctx, queryInsertUserIDReceiptID, userID, receiptID)
// 	if pairErr != nil {
// 		return errors.NewInternalServerError("database transaction pair error")
// 	}

// 	// Get User points
// 	var userPoints int64
// 	getPointsErr := tx.QueryRowContext(ctx, queryGetUserPointsByUserID, userID).Scan(&userPoints)
// 	if getPointsErr != nil {
// 		return errors.NewInternalServerError("database transaction get points error")
// 	}

// 	// Update User Points
// 	_, pointsErr := tx.ExecContext(ctx, queryUpdateUserPoints, userPoints+1, 0, userID)
// 	if pointsErr != nil {
// 		return errors.NewInternalServerError("database transaction add points error")
// 	}

// 	// Commit the transaction.
// 	if err = tx.Commit(); err != nil {
// 		return errors.NewInternalServerError("database transaction commit error")
// 	}
// 	return nil
// }

func (receipt *Receipt) Save2() *errors.RestErr {
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

func (receipt *Receipt) GetAllCountToday() (int64, *errors.RestErr) {
	var count int64

	row := receipts_db.Client.QueryRow(queryGetAllCountToday)
	err := row.Scan(&count)
	if err != nil {
		return 0, errors.NewInternalServerError("database error")
	}

	return count, nil
}

func (receiptDAO *ReceiptDAO) GetUnverifiedReceipt() (Receipt, *errors.RestErr) {
	// TODO - on retrieval, add a score of 0 so that other QA requests don't get the same receipt to score
	// 0 score should be updated to something once the evaluation is done
	// if evaluation fails or is not done - user that submitted the receipt should inform the 0 score - DB fix to null
	receipt := Receipt{}
	stmt, err := receipts_db.Client.Prepare(queryGetLatestUnverifiedReceipt)
	if err != nil {
		return receipt, errors.NewInternalServerError("error preparing unverified receipt")
	}
	defer stmt.Close()

	log.Println("datetime before:", receiptDAO.DateAdded, receiptDAO.QAScore, receiptDAO.QADate, receiptDAO.ID)
	result := stmt.QueryRow()
	// if getErr := result.Scan(&receipt.ID, &receipt.ModelYear, &receipt.Make, &receipt.Vin, &receipt.FirstName, &receipt.LastName, &receipt.State, &receipt.DateAdded, &receipt.QAScore, &receipt.QADate); getErr != nil {
	if getErr := result.Scan(&receiptDAO.ID, &receiptDAO.ModelYear, &receiptDAO.Make, &receiptDAO.Vin, &receiptDAO.FirstName, &receiptDAO.LastName, &receiptDAO.State, &receiptDAO.DateAdded, &receiptDAO.QAScore); getErr != nil {
		log.Println(receiptDAO)
		return receipt, errors.NewInternalServerError("error retrieving unverified receipt")
	}
	// return nil

	
	receipt.ID = receiptDAO.ID
	receipt.ModelYear = receiptDAO.ModelYear
	receipt.Make = receiptDAO.Make
	receipt.Vin = receiptDAO.Vin
	receipt.FirstName = receiptDAO.FirstName
	receipt.LastName = receiptDAO.LastName
	receipt.State = receiptDAO.State
	// parsedDateAdded, parseErr := time.Parse(time.RFC3339, receiptDAO.DateAdded)
	// if parseErr != nil {
	// 	parseErrToDisplay := errors.NewBadRequestError("invalid DateAdded datetime format in DAO")
	// 	log.Println(parseErrToDisplay)
	// }
	// receipt.DateAdded = parsedDateAdded
	receipt.DateAdded = receiptDAO.DateAdded
	// var parseErr error
	receipt.QAScore = int(receiptDAO.QAScore.Int64)
	// parsedQADate := time.Time{}
	// if receiptDAO.QADate != "" {
	// 	parsedQADate, parseErr = time.Parse(time.RFC3339, receiptDAO.QADate)
	// 	if parseErr != nil {
	// 		parseErrToDisplay := errors.NewBadRequestError("invalid QADate datetime format in DTO")
	// 		log.Println(parseErrToDisplay)
	// 	}
	// }
	// receipt.QADate = parsedQADate
	receipt.QADate = receiptDAO.QADate
	log.Println("receipt:", receipt)
	if receipt.ID == int64(0) {
		// No valid receipts were found - all verified
		return receipt, errors.NewBadRequestError("could not find unverified receipts")
	}

	return receipt, nil
}

func (receipt *Receipt) UpdateReceipt() *errors.RestErr {
	log.Println(receipt)
	stmt, err := receipts_db.Client.Prepare(queryUpdateReceipt)
	if err != nil {
		return errors.NewInternalServerError("database error update receipt stmt")
	}
	defer stmt.Close()
	currentTime := time.Now()
	qaDate := currentTime.Format("20060102150405")
	// "UPDATE receipts SET model_year = ?, make = ?, vin = ?, first_name = ?, last_name = ?, state = ?, qa_score = ?, qa_date = ? WHERE id = ?;"
	updateResult, saveErr := stmt.Exec(receipt.ModelYear, receipt.Make, receipt.Vin, receipt.FirstName, receipt.LastName, receipt.State, receipt.QAScore, qaDate, receipt.ID)
	if saveErr != nil {
		return errors.NewInternalServerError("database error updating receipt")
	}

	receiptID, err := updateResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("database error getting update ID")
	}
	receipt.ID = receiptID
	return nil
}
