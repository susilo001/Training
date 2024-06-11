package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/session-3-unit-test/router"
)

func main() {
	// Inisialisasi router Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Definisikan routes
	router.SetupRouter(r)

	// Jalankan server pada port 8080
	log.Println("Running server on port 8080")
	r.Run(":8080")
}