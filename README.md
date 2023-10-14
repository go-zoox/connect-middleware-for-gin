# Connect Middleware for Gin

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/connect-middleware-for-gin)](https://pkg.go.dev/github.com/go-zoox/connect-middleware-for-gin)
[![Build Status](https://github.com/go-zoox/connect-middleware-for-gin/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/connect-middleware-for-gin/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/connect-middleware-for-gin)](https://goreportcard.com/report/github.com/go-zoox/connect-middleware-for-gin)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/connect-middleware-for-gin/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/connect-middleware-for-gin?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/connect-middleware-for-gin.svg)](https://github.com/go-zoox/connect-middleware-for-gin/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/connect-middleware-for-gin.svg?label=Release)](https://github.com/go-zoox/connect-middleware-for-gin/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/connect-middleware-for-gin
```

## Getting Started

```go
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
```

## Related Projects
* [go-zoox/connect](https://github.com/go-zoox/connect) - The Auth Connect.
* [go-zoox/connect-middleware-for-zoox](https://github.com/go-zoox/connect-middleware-for-zoox) - The Auth Connect Middleware for Zoox.

## License
GoZoox is released under the [MIT License](./LICENSE).
