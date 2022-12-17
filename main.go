package myLogger

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitFunction() {
	log.Println("initial function from myLogger")
}

func InitRoute() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"status":  true,
			"message": "Transaction success",
			"detail":  "",
		})
	})
	router.Run(":8088")
}
