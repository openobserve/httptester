package base

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	var v interface{}
	byteBody, _ := c.GetRawData()

	// get indented json for logging
	json.Unmarshal(byteBody, &v)
	indented, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(indented))

	c.JSON(http.StatusOK, gin.H{"body": v})
}
