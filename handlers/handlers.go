package main

import (
	"fmt"
	"net/http"
	"clients"
	"github.com/gin-gonic/gin"
)

// askGetHandler displays field to ask ChatGPT questions.
func askGEThandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// askPostHandler asks ChatGPT question from user and serves response.
func askPOSThandler(c *gin.Context) {
	website := c.PostForm("website")
	answ, err := clients.AskGPTansw("Generate a meta title and description for the website: " + website)
	if err != nil {
		fmt.Println("error asking Chat GPT:", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "Sorry, something went wrong"})
		return
	}
	c.HTML(http.StatusAccepted, "index.html", gin.H{"message": answ})
}
