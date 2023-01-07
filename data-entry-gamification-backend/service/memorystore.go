package service

import (
	"data-entry-gamification/model"
	"fmt"
)

// ReceiptMemoryStore is an in-memory implementation of ReceiptService
type ReceiptMemoryStore struct {
	Receipts []model.Receipt
}

// // Create a new ReceiptMemoryStore with some sample data
	// receiptStore := &service.ReceiptMemoryStore{
	// 	Receipts: []model.Receipt{
	// 		{ID: 1, FirstName: "Michael", LastName: "Motorist", Make: "Honda", ModelYear: 1999, State: "NY", Vin: "JHMCB7682PC021209"},
	// 		{ID: 2, FirstName: "John", LastName: "Motorist", Make: "Honda", ModelYear: 2012, State: "NY", Vin: "JHMCB7682PC021204"},
	// 		{ID: 3, FirstName: "Jane", LastName: "Motorist", Make: "Honda", ModelYear: 2002, State: "NY", Vin: "JHMCB7682PC021203"},
	// 	},
	// }

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
