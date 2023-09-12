package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	
	 "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/helper"
	"github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/utils/response"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

func AdminAuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	var tokenString string
	if strings.HasPrefix(s, "Bearer") {
		tokenString = strings.TrimPrefix(s, "Bearer ")
	} else {
		tokenString = s
	}

	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		errRes := response.Responses(http.StatusUnauthorized, "not authorised", nil, err.Error())
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	claims,ok:=token.Claims.(*helper.CustomAdminClaim)
	
	if !ok{
		errRes := response.Responses(http.StatusUnauthorized, "not authorised", nil, fmt.Errorf("claim not retrieved"))
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	
	
	if claims.Role=="user"{
		errRes := response.Responses(http.StatusUnauthorized, "no admin token", nil, errors.New("access restricted to only admin"))
		c.JSON(http.StatusUnauthorized, errRes)
		c.Abort()
		return
	}
	id := claims.ID
	c.Set("adminId", id)
	c.Next()
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString,&helper.CustomAdminClaim{} ,func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	return token, err
}


