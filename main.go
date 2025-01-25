package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("hello I am AI app")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8089")
}
