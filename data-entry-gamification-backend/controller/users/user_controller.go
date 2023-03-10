package users

import (
	"data-entry-gamification/model"
	"data-entry-gamification/service"
	"data-entry-gamification/utils/authentication"
	"data-entry-gamification/utils/errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	SecretKey = "abc123"
)

func Register(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json body")
		c.JSON(err.Status, err)
		return
	}

	result, saveErr := service.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Login(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("invalid json")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := service.GetUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := errors.NewInternalServerError("login failed")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
}

func Get(c *gin.Context) {
	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	result, restErr := service.GetUserByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func GetUserInfo(c *gin.Context) {
	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	result, restErr := service.GetUserInfoByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUserAvatar(c *gin.Context) {
	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	userInfo, restErr := service.GetUserInfoByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	// Serving the avatar file
	c.File(userInfo.ImageURI)
}

func PutUserAvatar(c *gin.Context) {
	// Authenticate From JWT
	issuer, err := authentication.AuthenticateFromJWT(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	userInfo, restErr := service.GetUserInfoByID(issuer)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	// PUT avatar based on user ID
	var requestUserAvatar model.UserAvatar

	if err := c.ShouldBind(&requestUserAvatar); err != nil {
		err := errors.NewBadRequestError("invalid avatar body")
		c.JSON(err.Status, err)
		return
	}

	result, restErr := service.PutUserAvatar(c, *userInfo, requestUserAvatar)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
