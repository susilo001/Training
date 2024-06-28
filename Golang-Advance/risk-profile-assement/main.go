package main

import (
	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
	"github.com/susilo001/golang-advance/risk-profile-assement/handler"
	"github.com/susilo001/golang-advance/risk-profile-assement/repository"

	"github.com/susilo001/golang-advance/risk-profile-assement/router"
	"github.com/susilo001/golang-advance/risk-profile-assement/service"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// setup gorm connectoin
	dsn := "postgresql://postgres:password@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	gormDB.AutoMigrate(&entity.User{}, &entity.Submission{})
	// setup service

	// uncomment to use postgres gorm
	userRepo := repository.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	submissionRepo := repository.NewSubmissionRepository(gormDB)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)

	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}
