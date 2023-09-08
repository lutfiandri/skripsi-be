package service

import (
	"context"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceTypeService interface {
	GetDeviceTypes(ctx context.Context) ([]rest.DeviceType, error)
	GetDeviceTypeById(ctx context.Context, id primitive.ObjectID) (rest.DeviceType, error)
	CreateDeviceType(ctx context.Context, deviceType rest.DeviceType)
	UpdateDeviceType(ctx context.Context, id primitive.ObjectID, deviceType rest.DeviceType) error
	DeleteDeviceType(ctx context.Context, id primitive.ObjectID) error
}

type deviceTypeService struct {
	repository repository.DeviceTypeRepository
}

func NewDeviceTypeService(repository repository.DeviceTypeRepository) DeviceTypeService {
	return &deviceTypeService{
		repository: repository,
	}
}

func (service *deviceTypeService) GetDeviceTypes(ctx context.Context) ([]rest.DeviceType, error) {
	result, err := service.repository.GetDeviceTypes(ctx)
	if err != nil {
		return []rest.DeviceType{}, err
	}

	deviceTypes := factory.DeviceTypeDbToRestMany(result)

	return deviceTypes, nil
}

func (service *deviceTypeService) GetDeviceTypeById(ctx context.Context, id primitive.ObjectID) (rest.DeviceType, error) {
	panic("unimplemented")
}

func (service *deviceTypeService) CreateDeviceType(ctx context.Context, deviceType rest.DeviceType) {
	panic("unimplemented")
}

func (service *deviceTypeService) UpdateDeviceType(ctx context.Context, id primitive.ObjectID, deviceType rest.DeviceType) error {
	panic("unimplemented")
}

func (service *deviceTypeService) DeleteDeviceType(ctx context.Context, id primitive.ObjectID) error {
	panic("unimplemented")
}
