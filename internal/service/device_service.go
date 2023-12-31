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

type DeviceService interface {
	GetDevices(ctx context.Context) ([]rest.DeviceResponse, error)
	GetDevice(ctx context.Context, request rest.GetDeviceRequest) (rest.DeviceResponse, error)
	CreateDevice(ctx context.Context, request rest.CreateDeviceRequest) (rest.DeviceResponse, error)
	UpdateDevice(ctx context.Context, request rest.UpdateDeviceRequest) (rest.DeviceResponse, error)
	UpdateDeviceVersion(ctx context.Context, request rest.UpdateDeviceVersionRequest) (rest.DeviceResponse, error)
	DeleteDevice(ctx context.Context, request rest.DeleteDeviceRequest) error

	AcquireDevice(ctx context.Context, claims rest.JWTClaims, request rest.AcquireDeviceRequest) (rest.DeviceResponse, error)
}

type deviceService struct {
	sync.Mutex
	repository repository.DeviceRepository
}

func NewDeviceService(repository repository.DeviceRepository) DeviceService {
	return &deviceService{
		repository: repository,
	}
}

func (service *deviceService) GetDevices(ctx context.Context) ([]rest.DeviceResponse, error) {
	result, err := service.repository.GetDevices(ctx)
	if err != nil {
		return []rest.DeviceResponse{}, err
	}

	devices := modelfactory.DeviceDomainToRestMany(result)

	return devices, nil
}

func (service *deviceService) GetDevice(ctx context.Context, request rest.GetDeviceRequest) (rest.DeviceResponse, error) {
	result, err := service.repository.GetDeviceById(ctx, request.Id)
	if err != nil {
		return rest.DeviceResponse{}, err
	}

	device := modelfactory.DeviceDomainToRest(result)

	return device, nil
}

func (service *deviceService) CreateDevice(ctx context.Context, request rest.CreateDeviceRequest) (rest.DeviceResponse, error) {
	now := time.Now()
	id := uuid.NewString()

	device := modelfactory.CreateDeviceRestToDomain(request)
	device.Id = id
	device.CreatedAt = now
	device.UpdatedAt = now

	if err := service.repository.UpsertDevice(ctx, id, device); err != nil {
		return rest.DeviceResponse{}, err
	}

	return modelfactory.DeviceDomainToRest(device), nil
}

func (service *deviceService) UpdateDevice(ctx context.Context, request rest.UpdateDeviceRequest) (rest.DeviceResponse, error) {
	service.Lock()
	defer service.Unlock()

	device, err := service.repository.GetDeviceById(ctx, request.Id)
	if err != nil {
		return rest.DeviceResponse{}, err
	}

	device.Name = request.Name
	device.Room = request.Room
	device.UpdatedAt = time.Now()

	if err := service.repository.UpsertDevice(ctx, request.Id, device); err != nil {
		return rest.DeviceResponse{}, err
	}

	return modelfactory.DeviceDomainToRest(device), nil
}

func (service *deviceService) UpdateDeviceVersion(ctx context.Context, request rest.UpdateDeviceVersionRequest) (rest.DeviceResponse, error) {
	service.Lock()
	defer service.Unlock()

	device, err := service.repository.GetDeviceById(ctx, request.Id)
	if err != nil {
		return rest.DeviceResponse{}, err
	}

	device.HwVersion = request.HwVersion
	device.SwVersion = request.SwVersion
	device.UpdatedAt = time.Now()

	if err := service.repository.UpsertDevice(ctx, request.Id, device); err != nil {
		return rest.DeviceResponse{}, err
	}

	return modelfactory.DeviceDomainToRest(device), nil
}

func (service *deviceService) AcquireDevice(ctx context.Context, claims rest.JWTClaims, request rest.AcquireDeviceRequest) (rest.DeviceResponse, error) {
	service.Lock()
	defer service.Unlock()

	device, err := service.repository.GetDeviceById(ctx, request.Id)
	if err != nil {
		return rest.DeviceResponse{}, err
	}

	device.UserId = &claims.User.Id
	device.UpdatedAt = time.Now()

	if err := service.repository.UpsertDevice(ctx, request.Id, device); err != nil {
		return rest.DeviceResponse{}, err
	}

	return modelfactory.DeviceDomainToRest(device), nil
}

func (service *deviceService) DeleteDevice(ctx context.Context, request rest.DeleteDeviceRequest) error {
	service.Lock()
	defer service.Unlock()

	if _, err := service.repository.GetDeviceById(ctx, request.Id); err != nil {
		return err
	}

	err := service.repository.DeleteDevice(ctx, request.Id)
	if err != nil {
		return err
	}

	return nil
}
