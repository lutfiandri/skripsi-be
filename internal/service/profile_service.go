package service

import (
	"context"
	"sync"
	"time"

	"skripsi-be/internal/model/rest"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory/modelfactory"
)

type ProfileService interface {
	GetProfile(ctx context.Context, claims rest.JWTClaims) (rest.ProfileResponse, error)
	UpdateProfile(ctx context.Context, claims rest.JWTClaims, request rest.UpdateProfileRequest) (rest.ProfileResponse, error)
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

func (service *profileService) GetProfile(ctx context.Context, claims rest.JWTClaims) (rest.ProfileResponse, error) {
	user, err := service.userRepository.GetUserByEmail(ctx, claims.User.Email)
	if err != nil {
		return rest.ProfileResponse{}, err
	}

	response := modelfactory.ProfileDbToRest(user)

	return response, nil
}

func (service *profileService) UpdateProfile(ctx context.Context, claims rest.JWTClaims, request rest.UpdateProfileRequest) (rest.ProfileResponse, error) {
	service.Lock()
	defer service.Unlock()

	user, err := service.userRepository.GetUserByEmail(ctx, claims.User.Email)
	if err != nil {
		return rest.ProfileResponse{}, err
	}

	user.Name = request.Name
	user.UpdatedAt = time.Now()

	if err := service.userRepository.UpsertUser(ctx, claims.User.Id, user); err != nil {
		return rest.ProfileResponse{}, err
	}

	response := modelfactory.ProfileDbToRest(user)

	return response, nil
}
