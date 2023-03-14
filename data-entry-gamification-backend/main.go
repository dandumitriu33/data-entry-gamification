package main

import (
	"data-entry-gamification/controller/receipts"
	"data-entry-gamification/controller/users"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

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

	router.POST("/api/receipts", receipts.AddReceipt)
	router.GET("/api/receipts/allcount", receipts.GetAllCount)
	router.GET("/api/receipts/allcounttoday", receipts.GetAllCountToday)
	router.GET("/api/receipts/unverified", receipts.GetUnverifiedReceipt)
	router.PUT("/api/receipts", receipts.UpdateReceipt)
	router.PUT("/api/receipts/verified", receipts.UpdateReceipt)

	router.POST("/api/register", users.Register)
	router.POST("/api/login", users.Login)
	router.GET("/api/user", users.Get)
	router.GET("/api/user/info", users.GetUserInfo)
	router.GET("/api/user/roles", users.GetUserRoles)
	router.GET("/api/user/avatar", users.GetUserAvatar)
	router.PUT("/api/user/avatar", users.PutUserAvatar)
	router.GET("/api/logout", users.Logout)

	router.Run("localhost:8080")
}
