package service

import (
	// "context"
	// "data-entry-gamification/model"
	"testing"
)

// func TestCreateReceipt(t *testing.T) {
// 	// Create a test context and receipt object
// 	ctx := context.Background()
// 	receipt := model.Receipt{
// 		ID:    1,
// 	}

// 	// Call the CreateReceipt function
// 	user := model.User{ID: 1}
// 	createdReceipt, err := CreateReceipt(ctx, receipt, user)

// 	// Check if an error was returned
// 	if err != nil {
// 		t.Errorf("CreateReceipt() returned an error: %v", err)
// 	}

// 	// Check if the receipt ID is not zero
// 	if createdReceipt.ID == 0 {
// 		t.Errorf("CreateReceipt() did not set the receipt ID")
// 	}
// }

func TestGetAllCount(t *testing.T) {
	// Call the GetAllCount function
	count, err := GetAllCount()

	// Check if an error was returned
	if err != nil {
		t.Errorf("GetAllCount() returned an error: %v", err)
	}

	// Check if the count is zero or greater
	if count < 0 {
		t.Errorf("GetAllCount() returned an invalid count")
	}
}

func TestGetAllCountToday(t *testing.T) {
	// Call the GetAllCountToday function
	count, err := GetAllCountToday()

	// Check if an error was returned
	if err != nil {
		t.Errorf("GetAllCountToday() returned an error: %v", err)
	}

	// Check if the count is zero or greater
	if count < 0 {
		t.Errorf("GetAllCountToday() returned an invalid count")
	}
}

func TestGetUnverifiedReceipt(t *testing.T) {
	// Call the GetUnverifiedReceipt function
	receipt, err := GetUnverifiedReceipt()

	// Check if an error was returned
	if err != nil {
		t.Errorf("GetUnverifiedReceipt() returned an error: %v", err)
	}

	// Check if the receipt ID is not zero
	if receipt.ID == 0 {
		t.Errorf("GetUnverifiedReceipt() returned an invalid receipt")
	}
}

// func TestUpdateReceipt(t *testing.T) {
// 	// Create a test receipt DTO object
// 	receiptDTO := model.ReceiptDTO{
// 		ID:          1,
// 	}

// 	// Call the UpdateReceipt function
// 	updatedReceiptDTO, err := UpdateReceipt(receiptDTO)

// 	// Check if an error was returned
// 	if err != nil {
// 		t.Errorf("UpdateReceipt() returned an error: %v", err)
// 	}

// 	// Check if the receipt DTO was updated
// 	if updatedReceiptDTO.ID != receiptDTO.ID {
// 		t.Errorf("UpdateReceipt() did not update the receipt DTO")
// 	}
// }
