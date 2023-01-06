package storage

import (
	"data-entry-gamification/model"
)

// Receipts slice to seed data.
var Receipts = []model.Receipt{
	{ID: 1, FirstName: "Michael", LastName: "Motorist", Make: "Honda", ModelYear: 1999, State: "NY", Vin: "JHMCB7682PC021209"},
	{ID: 2, FirstName: "John", LastName: "Motorist", Make: "Honda", ModelYear: 2012, State: "NY", Vin: "JHMCB7682PC021204"},
	{ID: 3, FirstName: "Jane", LastName: "Motorist", Make: "Honda", ModelYear: 2002, State: "NY", Vin: "JHMCB7682PC021203"},
}