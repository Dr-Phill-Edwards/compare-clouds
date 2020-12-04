package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var messages []string

func GetMessages(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func PostMessage(c *gin.Context) {
	text := c.PostForm("posttext")
	messages = append(messages, text)
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func PutMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	text := c.PostForm("puttext")
	messages[id] = text
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func OptionMessage(c *gin.Context) {
	for key, value := range c.Request.Header {
		fmt.Println(key, ": ", value)
	}
	c.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
}

func main() {
	messages = append(messages, "Hello Cloud!")
	r := gin.Default()
	r.GET("/api/messages", GetMessages)
	r.POST("/api/messages", PostMessage)
	r.PUT("/api/messages/:id", PutMessage)
	r.OPTIONS("/api/messages/:id", OptionMessage)
	r.Run(":8000")
}
