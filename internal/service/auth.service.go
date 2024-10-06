package service

import (
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/repository"
)

type AuthService struct {
	authRepo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{authRepo: repo}
}

func (s *AuthService) Register(req *dto.AuthRegisterReq) error {
	_, err := s.authRepo.Register(req)
	return err
}

func (s *AuthService) Login(req *dto.AuthLoginReq) (*model.UserModel, error) {
	user, err := s.authRepo.Login(req)
	if err != nil {
		return nil, err
	}

	// TODO: Password SHA256 Hashing
	// if !verifyPassword(user.Password, req.Password) { something }
	if user.Password != req.Password {
		return nil, err
	}
	return user, err
}
