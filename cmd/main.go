package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eneridangelis/device-rest/internal/api"
	"github.com/eneridangelis/device-rest/internal/repository"
	"github.com/eneridangelis/device-rest/internal/usecase"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := "host=" + os.Getenv("POSTGRES_HOST") +
		" user=" + os.Getenv("POSTGRES_USER") +
		" password=" + os.Getenv("POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("POSTGRES_DB") +
		" port=" + os.Getenv("POSTGRES_PORT") +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("successfully connected to the database!")

	deviceRepository := repository.NewDeviceRepository(db)

	deviceUsecase := usecase.NewDeviceUsecase(deviceRepository)

	deviceHandler := api.NewDeviceHandler(deviceUsecase)

	router := api.NewRouter(deviceHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
