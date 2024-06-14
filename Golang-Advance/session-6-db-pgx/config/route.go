package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/session-6-db-pgx/entity"
	"github.com/susilo001/golang-advance/session-6-db-pgx/handler"
	"github.com/susilo001/golang-advance/session-6-db-pgx/middleware"
	"github.com/susilo001/golang-advance/session-6-db-pgx/repository"
	"github.com/susilo001/golang-advance/session-6-db-pgx/router"
	"github.com/susilo001/golang-advance/session-6-db-pgx/service"
)

func InitRoute() {
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