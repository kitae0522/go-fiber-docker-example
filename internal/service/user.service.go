package service

import (
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) CreateProfile(req *dto.UserCreateProfileReq) error {
	_, err := s.userRepo.CreateProfile(req)
	return err
}

func (s *UserService) GetProfileByTag(req *dto.UserGetProfileTagReq) (*model.UserModel, error) {
	user, err := s.userRepo.GetProfileByTag(req)
	return user, err
}

func (s *UserService) GetProfileByID(req *dto.UserGetProfileIdReq) (*model.UserModel, error) {
	user, err := s.userRepo.GetProfileByID(req)
	return user, err
}

func (s *UserService) UpdateProfile(req *dto.UserUpdateProfileReq) (*dto.UserUpdateProfileEntity, error) {
	updatedUser, err := s.userRepo.UpdateProfile(req)
	return updatedUser, err
}

// func (s *UserService) DeleteProfile() {}
