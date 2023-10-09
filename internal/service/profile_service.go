package service

import (
	"context"
	"sync"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
)

type ProfileService interface {
	GetProfile(ctx context.Context, claims rest.JWTClaims) (rest.UserResponse, error)
	EditProfile(ctx context.Context, claims rest.JWTClaims, request rest.EditProfileRequest) (rest.UserResponse, error)
}

type profileService struct {
	sync.Mutex
	userRepository repository.UserRepository
}

func NewProfileService(userRepository repository.UserRepository) ProfileService {
	return &profileService{
		userRepository: userRepository,
	}
}

func (service *profileService) GetProfile(ctx context.Context, claims rest.JWTClaims) (rest.UserResponse, error) {
	user, err := service.userRepository.GetUserByEmail(ctx, claims.User.Email)
	if err != nil {
		return rest.UserResponse{}, err
	}

	response := rest.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	return response, nil
}

func (service *profileService) EditProfile(ctx context.Context, claims rest.JWTClaims, request rest.EditProfileRequest) (rest.UserResponse, error) {
	service.Lock()
	defer service.Unlock()

	user, err := service.userRepository.GetUserByEmail(ctx, claims.User.Email)
	if err == nil {
		return rest.UserResponse{}, err
	}

	user.Name = request.Name
	user.UpdatedAt = time.Now()

	if err := service.userRepository.UpsertUser(ctx, claims.User.Id, user); err != nil {
		return rest.UserResponse{}, err
	}

	response := rest.UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}

	return response, nil
}
