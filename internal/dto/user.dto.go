package dto

import "go-fiber-docker-example/internal/model"

type UserCreateProfileReq struct {
	UserTag    string `json:"userTag" validate:"required"`
	UserName   string `json:"userName" validate:"required"`
	ProfilePic string `json:"profilePic"`
	Bio        string `json:"bio"`
}

type UserCreateProfileRes struct {
	IsError    bool   `json:"isError"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type UserGetProfileTagReq struct {
	UserTag string `json:"userTag" validate:"required"`
}

type UserGetProfileIdReq struct {
	ID string `json:"id" validate:"required"`
}

type UserGetProfileRes struct {
	IsError    bool             `json:"isError"`
	StatusCode int              `json:"statusCode"`
	Message    string           `json:"message"`
	User       *model.UserModel `json:"user"`
}

type UserUpdateProfileReq struct {
	ID         string `json:"id" validate:"required"`
	UserTag    string `json:"userTag"`
	UserName   string `json:"userName"`
	ProfilePic string `json:"profilePic"`
	Bio        string `json:"bio"`
}

type UserUpdateProfileEntity struct {
	ID         string `json:"id"`
	UserTag    string `json:"userTag"`
	UserName   string `json:"userName"`
	ProfilePic string `json:"profilePic"`
	Bio        string `json:"bio"`
}

type UserUpdateProfileRes struct {
	IsError    bool                     `json:"isError"`
	StatusCode int                      `json:"statusCode"`
	Message    string                   `json:"message"`
	User       *UserUpdateProfileEntity `json:"user"`
}

type UserDeleteProfileReq struct{}

type UserDeleteProfileRes struct{}
