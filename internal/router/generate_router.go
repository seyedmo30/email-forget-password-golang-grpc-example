package router

import (
	"github.com/gin-gonic/gin"
	"github.com/seyedmo30/email-forget-password-golang-grpc-example/internal/controller"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/generate/:email", controller.Generate)

	return r
}
