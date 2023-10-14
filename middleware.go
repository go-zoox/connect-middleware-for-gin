package connect

import (
	"github.com/gin-gonic/gin"
	user "github.com/go-zoox/connect/user"
	"github.com/go-zoox/jwt"
)

// CreateOptions is the options for the Create middleware.
type CreateOptions struct {
	RequireAuth bool
}

// Create creates a connect middleware for gin.
func Create(secretKey string, opts ...*CreateOptions) gin.HandlerFunc {
	var signer jwt.Jwt
	var optsX *CreateOptions
	if len(opts) > 0 && opts[0] != nil {
		optsX = opts[0]
	}

	return func(ctx *gin.Context) {
		if signer == nil {
			signer = jwt.New(secretKey)
		}

		token := ctx.GetHeader("x-connect-token")
		if token != "" {
			user := &user.User{}
			if err := user.Decode(signer, token); err != nil {
				// if ctx.AcceptJSON() {
				// 	ctx.JSON(http.StatusUnauthorized, gin.H{
				// 		"code":    401001,
				// 		"message": "Unauthorized",
				// 	})
				// } else {
				// 	ctx.Status(401)
				// }

				ctx.Status(401)
				ctx.Abort()
				return
			}

			ctx.Set(ContextUserKey, user)
		}

		if optsX != nil && optsX.RequireAuth {
			if _, ok := ctx.Get(ContextUserKey); !ok {
				// if ctx.AcceptJSON() {
				// 	ctx.JSON(http.StatusUnauthorized, gin.H{
				// 		"code":    401001,
				// 		"message": "Unauthorized",
				// 	})
				// } else {
				// 	ctx.Status(401)
				// }

				ctx.Status(401)
				ctx.Abort()
				return
			}
		}

		ctx.Next()
	}
}
