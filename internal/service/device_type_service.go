package service

import (
	"context"
	"sync"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory/modelfactory"

	"github.com/google/uuid"
)

type DeviceTypeService interface {
	GetDeviceTypes(ctx context.Context) ([]rest.DeviceTypeResponse, error)
	GetDeviceType(ctx context.Context, request rest.GetDeviceTypeRequest) (rest.DeviceTypeResponse, error)
	CreateDeviceType(ctx context.Context, request rest.CreateDeviceTypeRequest) (rest.DeviceTypeResponse, error)
	UpdateDeviceType(ctx context.Context, request rest.UpdateDeviceTypeRequest) (rest.DeviceTypeResponse, error)
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

func (service *deviceTypeService) GetDeviceTypes(ctx context.Context) ([]rest.DeviceTypeResponse, error) {
	result, err := service.repository.GetDeviceTypes(ctx)
	if err != nil {
		return []rest.DeviceTypeResponse{}, err
	}

	deviceTypes := modelfactory.DeviceTypeDbToRestMany(result)

	return deviceTypes, nil
}

func (service *deviceTypeService) GetDeviceType(ctx context.Context, request rest.GetDeviceTypeRequest) (rest.DeviceTypeResponse, error) {
	result, err := service.repository.GetDeviceTypeById(ctx, request.Id)
	if err != nil {
		return rest.DeviceTypeResponse{}, err
	}

	deviceType := modelfactory.DeviceTypeDbToRest(result)

	return deviceType, nil
}

func (service *deviceTypeService) CreateDeviceType(ctx context.Context, request rest.CreateDeviceTypeRequest) (rest.DeviceTypeResponse, error) {
	now := time.Now()
	id := uuid.NewString()

	deviceType := modelfactory.CreateDeviceTypeRestToDb(request)
	deviceType.Id = id
	deviceType.CreatedAt = now
	deviceType.UpdatedAt = now

	if err := service.repository.UpsertDeviceType(ctx, id, deviceType); err != nil {
		return rest.DeviceTypeResponse{}, err
	}

	return modelfactory.DeviceTypeDbToRest(deviceType), nil
}

func (service *deviceTypeService) UpdateDeviceType(ctx context.Context, request rest.UpdateDeviceTypeRequest) (rest.DeviceTypeResponse, error) {
	service.Lock()
	defer service.Unlock()

	if _, err := service.repository.GetDeviceTypeById(ctx, request.Id); err != nil {
		return rest.DeviceTypeResponse{}, err
	}

	deviceType := modelfactory.UpdateDeviceTypeRestToDb(request)
	deviceType.UpdatedAt = time.Now()

	if err := service.repository.UpsertDeviceType(ctx, request.Id, deviceType); err != nil {
		return rest.DeviceTypeResponse{}, err
	}

	return modelfactory.DeviceTypeDbToRest(deviceType), nil
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
