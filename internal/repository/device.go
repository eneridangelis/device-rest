package repository

import (
	"github.com/eneridangelis/device-rest/internal/model"
)

func (r *deviceRepository) Add(device *model.Device) error {
	return r.db.Create(device).Error
}

func (r *deviceRepository) GetByID(id uint) (*model.Device, error) {
	var device model.Device
	err := r.db.First(&device, id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) List() ([]*model.Device, error) {
	var devices []*model.Device
	err := r.db.Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *deviceRepository) Update(device *model.Device) error {
	return r.db.Save(device).Error
}

func (r *deviceRepository) Delete(id uint) error {
	return r.db.Delete(&model.Device{}, id).Error
}

func (r *deviceRepository) SearchByBrand(brand string) ([]*model.Device, error) {
	var devices []*model.Device
	err := r.db.Where("brand = ?", brand).Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}
