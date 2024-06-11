package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/crud/entity"
	"github.com/susilo001/golang-advance/crud/handler"
	"github.com/susilo001/golang-advance/crud/middleware"
	"github.com/susilo001/golang-advance/crud/repository"
	"github.com/susilo001/golang-advance/crud/router"
	"github.com/susilo001/golang-advance/crud/service"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.AuthMiddleware())
	
	var mockUserDBInSlice []entity.User
	userRepo := repository.NewUserRepository(mockUserDBInSlice)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRouter(r, userHandler)

	log.Println("Running Server On Port 8080")

	r.Run(":8080")
}
