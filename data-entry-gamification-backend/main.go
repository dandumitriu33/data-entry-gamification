package main

import (
	"data-entry-gamification/model"
	"data-entry-gamification/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	receiptStore := &service.MySQL{}

	router := gin.Default()
	// router.GET("/receipts", getReceipts)
	router.GET("/receipts", func(c *gin.Context) {
		c.JSON(http.StatusOK, receiptStore.GetAll())
	})
	// router.GET("/receipts/:id", getReceiptByID)
	router.GET("/receipts/:id", func(c *gin.Context) {
		// Get the ID from the path parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		// Get the receipt with the matching ID
		receipt, err := receiptStore.GetByID(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, receipt)
	})
	// router.POST("/receipts", postReceipts)
	router.POST("/receipts", func(c *gin.Context) {
		var newReceipt model.Receipt
		if err := c.BindJSON(&newReceipt); err != nil {
			return
		}

		receiptStore.PostReceipt(newReceipt)
		c.JSON(http.StatusOK, newReceipt)
	})

	router.Run("localhost:8080")
}

// // getReceipts responds with the list of all receipts as JSON.
// func getReceipts(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, receiptStore.GetAll())
// }

// // postReceipts adds a receipt from JSON received in the request body.
// func postReceipts(c *gin.Context) {
// 	var newReceipt model.Receipt

// 	// Call BindJSON to bind the received JSON to
// 	// newReceipt.
// 	if err := c.BindJSON(&newReceipt); err != nil {
// 		return
// 	}

// 	// Add the new receipt to the slice.
// 	storage.Receipts = append(storage.Receipts, newReceipt)
// 	c.IndentedJSON(http.StatusCreated, newReceipt)
// }

// // getReceiptByID locates the receipt whose ID value matches the id
// // parameter sent by the client, then returns that receipt as a response.
// func getReceiptByID(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		log.Panicf("Error converting ID from path: %s\n", err)
// 	}

// 	// Loop over the list of receipts, looking for
// 	// a receipt whose ID value matches the parameter.
// 	for _, r := range storage.Receipts {
// 		if r.ID == id {
// 			c.IndentedJSON(http.StatusOK, r)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "receipt not found"})
// }
