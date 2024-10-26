package main

import (
	"log"

	"github.com/eneridangelis/device-rest/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// [ene] passar isso para variáveis locais depois
	dsn := "host=localhost user=user password=1234 dbname=db_test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("successfully connected to the database!")

	deviceRepo := repository.NewDeviceRepository(db)
}
