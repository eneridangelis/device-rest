package repository

import (
	"github.com/eneridangelis/device-rest/internal/model"
	"gorm.io/gorm"
)

type DeviceRepository interface {
	Add(device *model.Device) error
	GetByID(id uint) (*model.Device, error)
	List() ([]*model.Device, error)
	Update(device *model.Device) error
	Delete(id uint) error
	SearchByBrand(brand string) ([]*model.Device, error)
}

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{db: db}
}
