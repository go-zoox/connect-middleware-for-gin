package connect

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetToken get user from context.
func GetToken(ctx *gin.Context) (token string, err error) {
	token = ctx.GetHeader("x-connect-token")
	if token == "" {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}

// MustGetToken get user from context.
func MustGetToken(ctx *gin.Context) (token string) {
	token, err := GetToken(ctx)
	if err != nil {
		panic(err)
	}

	return token
}
