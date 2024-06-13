package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/unit-test-crud-user/entity"
	"github.com/susilo001/golang-advance/unit-test-crud-user/handler"
	"github.com/susilo001/golang-advance/unit-test-crud-user/middleware"
	"github.com/susilo001/golang-advance/unit-test-crud-user/repository"
	"github.com/susilo001/golang-advance/unit-test-crud-user/router"
	"github.com/susilo001/golang-advance/unit-test-crud-user/service"
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
