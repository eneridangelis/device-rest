package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/eneridangelis/device-rest/internal/model"
	"github.com/eneridangelis/device-rest/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var repo repository.DeviceRepository

func TestMain(m *testing.M) {
	dsn := "host=localhost user=user password=1234 dbname=db_test_integration port=5433 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&model.Device{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo = repository.NewDeviceRepository(db)

	code := m.Run()

	db.Exec("DROP TABLE devices")

	os.Exit(code)
}

func TestAddDevice(t *testing.T) {
	device := &model.Device{Name: "Test Device", Brand: "Test Brand"}

	err := repo.Add(device)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Device
	if err := db.First(&result, device.ID).Error; err != nil {
		t.Fatalf("device not found in database: %v", err)
	}

	if result.Name != "Test Device" {
		t.Errorf("expected name to be 'Test Device', got %v", result.Name)
	}

	tearDown()
}

func TestGetDeviceByID(t *testing.T) {
	device := &model.Device{Name: "Test Device 2", Brand: "Test Brand 2"}
	repo.Add(device)

	result, err := repo.GetByID(device.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.Name != device.Name {
		t.Errorf("expected name %v, got %v", device.Name, result.Name)
	}

	tearDown()
}

func TestListDevices(t *testing.T) {
	devices := []*model.Device{
		{Name: "Device 1", Brand: "Brand A"},
		{Name: "Device 2", Brand: "Brand B"},
	}
	for _, device := range devices {
		repo.Add(device)
	}

	result, err := repo.List()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != len(devices) {
		t.Errorf("expected %d devices, got %d", len(devices), len(result))
	}

	tearDown()
}

func TestUpdateDevice(t *testing.T) {
	device := &model.Device{Name: "Device Original", Brand: "Brand Original"}
	repo.Add(device)

	device.Name = "Device Updated"
	device.Brand = "Brand Updated"
	if err := repo.Update(device); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Device
	if err := db.First(&result, device.ID).Error; err != nil {
		t.Fatalf("device not found in database: %v", err)
	}

	if result.Name != "Device Updated" || result.Brand != "Brand Updated" {
		t.Errorf("expected updated device name and brand, got name: %v, brand: %v", result.Name, result.Brand)
	}

	tearDown()
}

func TestDeleteDevice(t *testing.T) {
	device := &model.Device{Name: "Device to Delete", Brand: "Brand Delete"}
	repo.Add(device)

	if err := repo.Delete(device.ID); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result model.Device
	err := db.First(&result, device.ID).Error
	if err == nil {
		t.Errorf("expected device to be deleted, but found in database")
	}

	tearDown()
}

func TestSearchByBrand(t *testing.T) {
	devices := []*model.Device{
		{Name: "Device Brand A", Brand: "Brand A"},
		{Name: "Device Brand A2", Brand: "Brand A"},
		{Name: "Device Brand B", Brand: "Brand B"},
	}
	for _, device := range devices {
		repo.Add(device)
	}

	result, err := repo.SearchByBrand("Brand A")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 devices with Brand A, got %d", len(result))
	}

	for _, device := range result {
		if device.Brand != "Brand A" {
			t.Errorf("expected brand to be 'Brand A', got %v", device.Brand)
		}
	}

	tearDown()
}

func tearDown() {
	db.Exec("TRUNCATE TABLE devices RESTART IDENTITY CASCADE")
}
