package receipts

import (
	"data-entry-gamification/model"
	"data-entry-gamification/service"
	"data-entry-gamification/utils/errors"
	"data-entry-gamification/utils/string_utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	SecretKey = "abc123"
)

func AddReceipt(c *gin.Context) {
	var receipt model.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	// Add receipt and points to user
	// Get User ID from JWT token (separate service and DB in future)
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	// token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")
		c.JSON(restErr.Status, restErr)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, restErr := service.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	// Add Receipt and other transaction parts (points, level)
	result, saveErr := service.CreateReceipt(c, receipt, *user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

// func AddReceipt2(c *gin.Context) {
// 	var receipt model.Receipt

// 	if err := c.ShouldBindJSON(&receipt); err != nil {
// 		err := errors.NewBadRequestError("invalid json body")
// 		c.JSON(err.Status, err)
// 		return
// 	}

// 	result, saveErr := service.CreateReceipt(receipt)
// 	if saveErr != nil {
// 		c.JSON(saveErr.Status, saveErr)
// 		return
// 	}

// 	c.JSON(http.StatusOK, result)
// }

func GetAllCount(c *gin.Context) {
	result, err := service.GetAllCount()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetAllCountToday(c *gin.Context) {
	result, err := service.GetAllCountToday()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetUnverifiedReceipt(c *gin.Context) {
	result, err := service.GetUnverifiedReceipt()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func UpdateReceipt(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("could not retrieve cookie")
		c.JSON(getErr.Status, getErr)
		return
	}

	// token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")
		c.JSON(restErr.Status, restErr)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	log.Println("getting user roles by user ID", issuer)
	// Confirm user has QA role
	userRoles, restErr := service.UserRoles(c, issuer)
	if err != nil {
		err := errors.NewBadRequestError("invalid user data")
		c.JSON(err.Status, restErr)
		return
	}
	log.Println("checking user roles", userRoles)
	if !string_utils.Contains(userRoles, "qa") {
		err := errors.NewBadRequestError("invalid user athorization")
		c.JSON(err.Status, err)
		return
	}

	// PUT request
	var receipt model.Receipt

	if err := c.ShouldBind(&receipt); err != nil {
		err := errors.NewBadRequestError("invalid receipt data")
		c.JSON(err.Status, err)
		return
	}

	result, restErr := service.UpdateReceipt(receipt)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
