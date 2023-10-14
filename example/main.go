package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-zoox/connect-middleware-for-gin"
)

func main() {
	r := gin.Default()

	r.Use(connect.Create("YOUR_SECRET_KEY"))

	r.GET("/user", func(c *gin.Context) {
		user, err := connect.GetUser(c)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.JSON(200, gin.H{
			"user": user,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "helloworld",
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
