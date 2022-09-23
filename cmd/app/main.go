package main

import (
	"github.com/seyedmo30/email-forget-password-golang-grpc-example/internal/router"
)

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
