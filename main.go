package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		log.Println("main logic")
		c.JSON(200, gin.H{"message": "Hello!"})
	})
	r.Run()
}
