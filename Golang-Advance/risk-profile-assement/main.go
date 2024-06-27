package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/susilo001/golang-advance/risk-profile-assement/entity"
	"github.com/susilo001/golang-advance/risk-profile-assement/handler"
	postgres_pgx "github.com/susilo001/golang-advance/risk-profile-assement/repository/pgx"
	postgres_gorm "github.com/susilo001/golang-advance/risk-profile-assement/repository/postgre_gorm"

	"github.com/susilo001/golang-advance/risk-profile-assement/router"
	"github.com/susilo001/golang-advance/risk-profile-assement/service"

	"log"

	"github.com/jackc/pgx/v5/pgxpool"
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

	// slice db is disabled. uncomment to enabled
	// var mockUserDBInSlice []entity.User
	// _ = slice.NewUserRepository(mockUserDBInSlice)

	// uncomment to use postgres pgx
	// userRepo := postgres_pgx.NewUserRepository(pgxPool)

	// uncomment to use postgres gorm
	userRepo := postgres_gorm.NewUserRepository(gormDB)

	// service and handler declaration
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	submissionRepo := postgres_gorm.NewSubmissionRepository(gormDB)
	submissionService := service.NewSubmissionService(submissionRepo)
	submissionHandler := handler.NewSubmissionHandler(submissionService)


	// Routes
	router.SetupRouter(r, userHandler, submissionHandler)

	// Run the server
	log.Println("Running server on port 8080")
	r.Run(":8080")
}

func connectDB(dbURL string) (postgres_pgx.PgxPoolIface, error) {
	return pgxpool.New(context.Background(), dbURL)
}