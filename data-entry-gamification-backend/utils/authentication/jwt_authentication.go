package authentication

import (
	"data-entry-gamification/utils/errors"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"strconv"
)

const (
	SecretKey = "abc123"
)

func AuthenticateFromJWT(c *gin.Context) (int64, *errors.RestErr){

	// Add receipt and points to user
	// Get User ID from JWT token (separate service and DB in future)
	cookie, err := c.Cookie("jwt")
	if err != nil {
		getErr := errors.NewInternalServerError("error retrieving cookie")
		return 0, getErr
	}

	// token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(*jwt.Token) (interface{}, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		restErr := errors.NewInternalServerError("error parsing cookie")		
		return 0, restErr
	}

	claims := token.Claims.(*jwt.StandardClaims)
	issuer, err := strconv.ParseInt(claims.Issuer, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		return 0, restErr
	}

	return issuer, nil
}

