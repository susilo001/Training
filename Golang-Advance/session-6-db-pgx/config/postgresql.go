package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitPostgre() {
	dsn := "postgresql://postgres:password@localhost:5432/postgres"
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Successfull Connect to Postgres")
}