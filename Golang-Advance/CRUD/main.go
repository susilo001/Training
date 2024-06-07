package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/crud/middleware"
	"github.com/susilo001/golang-advance/crud/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	router.SetupRouter(r)

	log.Println("Running Server On Port 8080")

	r.Run(":8080")
}
