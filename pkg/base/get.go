package base

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	fmt.Println("Path: ", c.FullPath())
	c.JSON(http.StatusOK, gin.H{
		"status": "good",
	})
}
