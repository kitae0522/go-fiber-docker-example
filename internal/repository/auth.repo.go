package repository

import (
	"context"
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/model"
)

type AuthRepository struct {
	client *model.PrismaClient
}

func NewAuthRepository(prismaClient *model.PrismaClient) *AuthRepository {
	return &AuthRepository{client: prismaClient}
}

func (r *AuthRepository) Register(req *dto.AuthRegisterReq) (*model.UserModel, error) {
	user, err := r.client.User.CreateOne(
		model.User.UserTag.Set(req.UserTag),
		model.User.UserName.Set(req.UserName),
		model.User.Password.Set(req.Password),
		model.User.Email.Set(req.Email),
	).Exec(context.Background())
	return user, err
}

func (r *AuthRepository) Login(req *dto.AuthLoginReq) (*model.UserModel, error) {
	user, err := r.client.User.FindUnique(
		model.User.Email.Equals(req.Email),
	).Exec(context.Background())
	return user, err
}
