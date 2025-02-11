package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObjectParams struct {
	Object string `uri:"object" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.GET("/large-content-length", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Length", "214748364800")
		c.JSON(200, gin.H{
			"message": "abc",
		})
	})

	r.HEAD("/large-content-length", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Length", "214748364800")
		c.JSON(200, gin.H{
			"message": "abc",
		})
	})

	r.GET("/bucket/*object", func(c *gin.Context) {
		var params ObjectParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
		fmt.Println("object: ", params.Object)

		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	r.Run("0.0.0.0:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
