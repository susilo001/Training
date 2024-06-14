package main

import (
	"github.com/susilo001/golang-advance/session-6-db-pgx/config"
)

func main() {
	config.InitPostgre()
	config.InitRoute()
}
