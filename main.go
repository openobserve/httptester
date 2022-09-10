package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "good",
		})
	})

	r.POST("/api/_bulk", func(c *gin.Context) {

		byteBody, _ := c.GetRawData()
		stringBody := string(byteBody)

		fmt.Println(stringBody)
		c.JSON(http.StatusOK, gin.H{"body": stringBody})
	})
	r.Run(":6080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
