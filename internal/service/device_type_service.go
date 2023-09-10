package service

import (
	"context"
	"sync"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory"

	"github.com/google/uuid"
)

type DeviceTypeService interface {
	GetDeviceTypes(ctx context.Context) ([]rest.DeviceType, error)
	GetDeviceType(ctx context.Context, request rest.GetDeviceTypeRequest) (rest.DeviceType, error)
	CreateDeviceType(ctx context.Context, request rest.CreateDeviceTypeRequest) (rest.DeviceType, error)
	UpdateDeviceType(ctx context.Context, request rest.UpdateDeviceTypeRequest) (rest.DeviceType, error)
	DeleteDeviceType(ctx context.Context, request rest.DeleteDeviceTypeRequest) error
}

type deviceTypeService struct {
	sync.Mutex
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

func (service *deviceTypeService) GetDeviceType(ctx context.Context, request rest.GetDeviceTypeRequest) (rest.DeviceType, error) {
	result, err := service.repository.GetDeviceTypeById(ctx, request.Id)
	if err != nil {
		return rest.DeviceType{}, err
	}

	deviceType := factory.DeviceTypeDbToRest(result)

	return deviceType, nil
}

func (service *deviceTypeService) CreateDeviceType(ctx context.Context, request rest.CreateDeviceTypeRequest) (rest.DeviceType, error) {
	now := time.Now()
	id := uuid.NewString()

	deviceType := factory.CreateDeviceTypeRestToDb(request)
	deviceType.Id = id
	deviceType.CreatedAt = now
	deviceType.UpdatedAt = now

	if err := service.repository.UpsertDeviceType(ctx, id, deviceType); err != nil {
		return rest.DeviceType{}, err
	}

	return factory.DeviceTypeDbToRest(deviceType), nil
}

func (service *deviceTypeService) UpdateDeviceType(ctx context.Context, request rest.UpdateDeviceTypeRequest) (rest.DeviceType, error) {
	service.Lock()
	defer service.Unlock()

	if _, err := service.repository.GetDeviceTypeById(ctx, request.Id); err != nil {
		return rest.DeviceType{}, err
	}

	deviceType := factory.UpdateDeviceTypeRestToDb(request)
	deviceType.UpdatedAt = time.Now()

	if err := service.repository.UpsertDeviceType(ctx, request.Id, deviceType); err != nil {
		return rest.DeviceType{}, err
	}

	return factory.DeviceTypeDbToRest(deviceType), nil
}

func (service *deviceTypeService) DeleteDeviceType(ctx context.Context, request rest.DeleteDeviceTypeRequest) error {
	service.Lock()
	defer service.Unlock()

	if _, err := service.repository.GetDeviceTypeById(ctx, request.Id); err != nil {
		return err
	}

	err := service.repository.DeleteDeviceType(ctx, request.Id)
	if err != nil {
		return err
	}

	return nil
}
