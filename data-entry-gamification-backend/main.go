package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type receipt struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Make      string `json:"make"`
	ModelYear int    `json:"model_year"`
	State     string `json:"state"`
	Vin       string `json:"vin"`
}

// albums slice to seed record album data.
var receipts = []receipt{
	{ID: 1, FirstName: "Michael", LastName: "Motorist", Make: "Honda", ModelYear: 1999, State: "NY", Vin: "JHMCB7682PC021209"},
	{ID: 2, FirstName: "John", LastName: "Motorist", Make: "Honda", ModelYear: 2012, State: "NY", Vin: "JHMCB7682PC021204"},
	{ID: 3, FirstName: "Jane", LastName: "Motorist", Make: "Honda", ModelYear: 2002, State: "NY", Vin: "JHMCB7682PC021203"},
}

func main() {
	varName := "MYSQL_DEV_USERNAME"
	value, exists := os.LookupEnv(varName)

	if exists {
		log.Printf(">>>>>>> Got ENV: %s\n", value)
	} else {
		log.Printf(">>>>>>> %s does not exist.\n", varName)
	}

	router := gin.Default()
	router.GET("/receipts", getReceipts)
	router.GET("/receipts/:id", getReceiptByID)
	router.POST("/receipts", postReceipts)

	router.Run("localhost:8080")
}

// getReceipts responds with the list of all receipts as JSON.
func getReceipts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, receipts)
}

// postReceipts adds a receipt from JSON received in the request body.
func postReceipts(c *gin.Context) {
	var newReceipt receipt

	// Call BindJSON to bind the received JSON to
	// newReceipt.
	if err := c.BindJSON(&newReceipt); err != nil {
		return
	}

	// Add the new receipt to the slice.
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusCreated, newReceipt)
}

// getReceiptByID locates the receipt whose ID value matches the id
// parameter sent by the client, then returns that receipt as a response.
func getReceiptByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicf("Error converting ID from path: %s\n", err)
	}

	// Loop over the list of receipts, looking for
	// a receipt whose ID value matches the parameter.
	for _, r := range receipts {
		if r.ID == id {
			c.IndentedJSON(http.StatusOK, r)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
}
