package repository

import (
	"context"
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/model"
)

type UserRepository struct {
	client *model.PrismaClient
}

func NewUserRepository(prismaClient *model.PrismaClient) *UserRepository {
	return &UserRepository{client: prismaClient}
}

func (r *UserRepository) CreateProfile(req *dto.UserCreateProfileReq) (*model.UserProfileModel, error) {
	user, err := r.GetProfileByTag(&dto.UserGetProfileTagReq{UserTag: req.UserTag})
	if err != nil {
		return nil, err
	}
	profile, err := r.client.UserProfile.CreateOne(
		model.UserProfile.User.Link(
			model.User.Cuuid.Equals(user.Cuuid),
		),
		model.UserProfile.ProfilePic.Set(req.ProfilePic),
		model.UserProfile.Bio.Set(req.Bio),
	).Exec(context.Background())
	return profile, err
}

func (r *UserRepository) GetProfileByTag(req *dto.UserGetProfileTagReq) (*model.UserModel, error) {
	user, err := r.client.User.FindUnique(
		model.User.UserTag.Equals(req.UserTag),
	).Exec(context.Background())
	return user, err
}

func (r *UserRepository) GetProfileByID(req *dto.UserGetProfileIdReq) (*model.UserModel, error) {
	user, err := r.client.User.FindUnique(
		model.User.Cuuid.Equals(req.ID),
	).Exec(context.Background())
	return user, err
}

func (r *UserRepository) UpdateProfile(req *dto.UserUpdateProfileReq) (*dto.UserUpdateProfileEntity, error) {
	updatedUser := new(dto.UserUpdateProfileEntity)

	user, err := r.client.User.FindUnique(
		model.User.Cuuid.Equals(req.ID),
	).Update(
		model.User.UserTag.Set(req.UserTag),
		model.User.UserName.Set(req.UserName),
	).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	profile, err := r.client.UserProfile.FindUnique(
		model.UserProfile.UserID.Equals(req.ID),
	).Update(
		model.UserProfile.ProfilePic.Set(req.ProfilePic),
		model.UserProfile.Bio.Set(req.Bio),
	).Exec(context.Background())

	updatedUser.ID = profile.UserID
	updatedUser.UserTag = user.UserTag
	updatedUser.UserName = user.UserName
	updatedUser.ProfilePic, _ = profile.ProfilePic()
	updatedUser.Bio, _ = profile.Bio()

	return updatedUser, err
}
