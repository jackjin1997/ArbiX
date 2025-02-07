package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/arbitrage", func(c *gin.Context) {
		// 这里将来放套利逻辑
		c.JSON(http.StatusOK, gin.H{
			"message": "ArbiX is working",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
