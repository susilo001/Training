package config

import (
	"log"

	"github.com/susilo001/golang-advance/session-13-grpc-gateway/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabases() {
	// setup gorm connection
	dsn := "postgresql://postgres:password@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entity.User{})
}