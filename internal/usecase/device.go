package usecase

import (
	"github.com/eneridangelis/device-rest/internal/model"
	db "github.com/eneridangelis/device-rest/internal/repository"
)

type DeviceUsecase struct {
	repository db.DeviceRepository
}

func NewDeviceUsecase(repository db.DeviceRepository) *DeviceUsecase {
	return &DeviceUsecase{repository: repository}
}

func (d *DeviceUsecase) AddDevice(device *model.Device) error {
	return d.repository.Add(device)
}

func (d *DeviceUsecase) GetDeviceByID(id uint) (*model.Device, error) {
	return d.repository.GetByID(id)
}

func (d *DeviceUsecase) ListAllDevices() ([]*model.Device, error) {
	return d.repository.List()
}

func (d *DeviceUsecase) UpdateDevice(device *model.Device) error {
	return d.repository.Update(device)
}

func (d *DeviceUsecase) DeleteDevice(id uint) error {
	return d.repository.Delete(id)
}

func (d *DeviceUsecase) SearchDeviceByBrand(brand string) ([]*model.Device, error) {
	return d.repository.SearchByBrand(brand)
}
