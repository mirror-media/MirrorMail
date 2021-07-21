package main

import (
	"MirrorMail/configs"
	"MirrorMail/mail"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := configs.LoadConfig(); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s !", err))
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"response": "OK"})
	})
	router.POST("/mail", mail.SendMail)
	router.Run()

}
