package controller

import (
	"net/http"

	"github.com/seyedmo30/email-forget-password-golang-grpc-example/internal/service"

	"github.com/gin-gonic/gin"
)

func Generate(c *gin.Context) {
	email_string := c.Param("email")

	code := service.GenerateCode()
	c.JSON(http.StatusOK, gin.H{"user": email_string, "value": code})
}
