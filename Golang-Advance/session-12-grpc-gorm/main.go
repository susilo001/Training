package main

import (
	"log"
	"net"

	"github.com/susilo001/golang-advance/session-12-grpc-gorm/entity"
	grpcHandler "github.com/susilo001/golang-advance/session-12-grpc-gorm/handler/grpc"
	pb "github.com/susilo001/golang-advance/session-12-grpc-gorm/proto/user/v1"
	"github.com/susilo001/golang-advance/session-12-grpc-gorm/repository"
	"github.com/susilo001/golang-advance/session-12-grpc-gorm/service"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// setup gorm connection
	dsn := "postgresql://postgres:password@localhost:5432/postgres"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	gormDB.AutoMigrate(&entity.User{})
	// setup service

	// uncomment to use postgres gorm
	userRepo := repository.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	//userHandler := ginHandler.NewUserHandler(userService)
	userHandler := grpcHandler.NewUserHandler(userService)

	// Run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Running grpc server in port :50051")
	_ = grpcServer.Serve(lis)
}
