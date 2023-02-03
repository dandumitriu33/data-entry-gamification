package receipts

import (
	"data-entry-gamification/model"
	"data-entry-gamification/service"
	"data-entry-gamification/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddReceipt(c *gin.Context) {
	var receipt model.Receipt

	if err := c.ShouldBindJSON(&receipt); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := service.CreateReceipt(receipt)
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
