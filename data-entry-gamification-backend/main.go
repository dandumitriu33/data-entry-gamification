package main

import (
	"data-entry-gamification/controller/receipts"
	"data-entry-gamification/controller/users"
	"data-entry-gamification/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	receiptStore := &service.MySQL{}

	router := gin.Default()
	// router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200"
		},
		MaxAge: 12 * time.Hour,
	}))
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
	// router.POST("/receipts", func(c *gin.Context) {
	// 	var newReceipt model.Receipt
	// 	if err := c.BindJSON(&newReceipt); err != nil {
	// 		return
	// 	}

	// 	receiptStore.PostReceipt(newReceipt)
	// 	c.JSON(http.StatusOK, newReceipt)
	// })

	router.POST("/receipts", receipts.AddReceipt)

	router.GET("/api/receipts/allcount", receipts.GetAllCount)
	router.GET("/api/receipts/allcounttoday", receipts.GetAllCountToday)

	router.POST("/api/register", users.Register)
	router.POST("/api/login", users.Login)
	router.GET("/api/user", users.Get)
	router.GET("/api/user/info", users.GetUserInfo)
	router.PUT("/api/user/avatar", users.PutUserAvatar)
	router.GET("/api/logout", users.Logout)

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
