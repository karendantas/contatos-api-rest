package main

import (
	"github.com/gin-gonic/gin"
)
func main () {
	server := gin.Default()

	// ponteiro para um contexto gin 
	server.GET("/ping", func (ctx *gin.Context) {
		//passava para o body a mensagem pong
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8000")
}