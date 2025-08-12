package main

import (
	"situsforum/internal/handlers/memberships"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	membershipHandler := memberships.NewHandler(router)
	membershipHandler.RegisterRoutes()

	// router.Run() // listen and serve on 0.0.0.0:8080
	router.Run(":9999") // listen and serve on 9999 port
}
