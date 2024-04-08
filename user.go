package connect

import (
	"fmt"

	"github.com/gin-gonic/gin"
	user "github.com/go-zoox/connect/user"
)

// GetUser gets the user from the context.
func GetUser(ctx *gin.Context) (u *user.User, err error) {
	v, exist := ctx.Get(ContextUserKey)
	if !exist {
		return nil, fmt.Errorf("user not found")
	}

	u, ok := v.(*user.User)
	if !ok {
		return nil, fmt.Errorf("user invalid")
	}

	return u, nil
}

// MustGetUser get user from context.
func MustGetUser(ctx *gin.Context) (u *user.User) {
	u, err := GetUser(ctx)
	if err != nil {
		panic(err)
	}

	return u
}
