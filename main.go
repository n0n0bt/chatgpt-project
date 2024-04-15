package main

import (
	"clients"
	"log"
	"handlers"
	"github.com/gin-gonic/gin"
)

var APIkey = "api..."
var client = clients.CreateClient(APIkey, "https://api.openai.com/v1/chat/completions")

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	setupRoutes(router)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(router *gin.Engine) {
	router.GET("/ask", handlers.askGEThandler)
	router.POST("/ask", handlers.askPOSThandler)
}
