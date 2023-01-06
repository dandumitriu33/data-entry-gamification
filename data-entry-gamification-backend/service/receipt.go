package service

import (
	"data-entry-gamification/model"
)

// ReceiptService is an interface for interacting with receipt data
type ReceiptService interface {
	GetAll() []model.Receipt
	// PostReceipt()
	GetByID(id int) (model.Receipt, error)
}