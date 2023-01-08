package service

import (
	"database/sql"
	"log"
	"os"

	"data-entry-gamification/model"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL is a struct that represents a MySQL connection.
type MySQL struct {
	DB *sql.DB
}

// Connect will open a connection to a MySQL database.
func (m *MySQL) Connect() error {
	dbUsername := os.Getenv("MYSQL_DEV_USERNAME")
	dbPassword := os.Getenv("MYSQL_DEV_PASSWORD")

	log.Printf("USR: %s\n", dbUsername)
	log.Printf("PWD: %s\n", dbPassword)
	dsn := dbUsername + ":" + dbPassword + "@tcp(127.0.0.1:3306)/vehicleregistration"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	m.DB = db
	log.Println("Connected")
	return nil
}

// Disconnect will close the connection to the MySQL database.
func (m *MySQL) Disconnect() error {
	log.Println("DB Disconnected.")
	return m.DB.Close()
}

// GetAll returns all receipts
func (m *MySQL) GetAll() []model.Receipt {
	m.Connect()
	defer m.Disconnect()

	rows, err := m.DB.Query("SELECT id, model_year, make, vin, first_name, last_name, state FROM receipts")
	// rows, err := m.DB.Query("SELECT id FROM receipts")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	
	result := []model.Receipt{}	
	for rows.Next() {
		var receipt model.Receipt		
		err = rows.Scan(&receipt.ID, &receipt.ModelYear, &receipt.Make, &receipt.Vin, &receipt.FirstName, &receipt.LastName, &receipt.State)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, receipt)
	}
	return result
}

// PostReceipt adds a new receipt to the store
func (m *MySQL) PostReceipt(receipt model.Receipt) {
	m.Connect()
	defer m.Disconnect()

	stmt, err := m.DB.Prepare("INSERT INTO receipts(model_year, make, vin, first_name, last_name, state) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	model_year := receipt.ModelYear
	make := receipt.Make
	vin := receipt.Vin
	first_name := receipt.FirstName
	last_name := receipt.LastName
	state := receipt.State
	_, err = stmt.Exec(model_year, make, vin, first_name, last_name, state)
	if err != nil {
		panic(err.Error())
	}
	log.Println("Receipt added to database.")
}

// GetByID returns the receipts with the given ID, or an error if no such receipts exists
func (m *MySQL) GetByID(id int) (model.Receipt, error) {
	var receipt model.Receipt
	return receipt, nil
}
