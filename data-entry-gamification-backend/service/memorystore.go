package service

import (
	"data-entry-gamification/model"
	"fmt"
)

// ReceiptMemoryStore is an in-memory implementation of ReceiptService
type ReceiptMemoryStore struct {
	Receipts []model.Receipt
}

// GetAll returns all receipts
func (s *ReceiptMemoryStore) GetAll() []model.Receipt {
	return s.Receipts
}

// GetByID returns the receipts with the given ID, or an error if no such receipts exists
func (s *ReceiptMemoryStore) GetByID(id int) (model.Receipt, error) {
	for _, u := range s.Receipts {
		if u.ID == id {
			return u, nil
		}
	}

	return model.Receipt{}, fmt.Errorf("receipt not found")
}

// PostReceipt adds a new receipt to the store
func (s *ReceiptMemoryStore) PostReceipt(receipt model.Receipt) {
	s.Receipts = append(s.Receipts, receipt)
}
