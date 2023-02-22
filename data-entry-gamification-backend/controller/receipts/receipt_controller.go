package receipts

import (
	"data-entry-gamification/model"
	"data-entry-gamification/service"
	"data-entry-gamification/utils/errors"
	"data-entry-gamification/utils/string_utils"
	"data-entry-gamification/utils/authentication"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReceipt(c *gin.Context) {
	// Bind JSON from form
	var receipt model.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c);
	if err != nil {
		c.JSON(err.Status, err)
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
	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c);
	if err != nil {
		c.JSON(err.Status, err)
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
