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

	// get indented json for body
	json.Unmarshal(byteBody, &v)

	res := make(map[string]interface{})

	res["headers"] = c.Request.Header
	res["body"] = v

	indented, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(indented))

	c.JSON(http.StatusOK, res)
}
