package service

import (
	"data-entry-gamification/model"
)

// ReceiptService is an interface for interacting with receipt data
type ReceiptService interface {
	GetAll() []model.Receipt
	PostReceipt(receipt model.Receipt)
	GetByID(id int) (model.Receipt, error)
}