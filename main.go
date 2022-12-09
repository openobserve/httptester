package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zinclabs/httptester/pkg/base"
)

func main() {
	r := gin.Default()

	r.GET("/", base.GET)
	r.POST("/api/_bulk", base.Post)

	// Opentelemetry
	r.POST("/api/v1/traces", base.Post)

	r.Run(":6080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
