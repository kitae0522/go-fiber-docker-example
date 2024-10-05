package dto

import (
	"time"

	"go-fiber-docker-example/internal/model"
)

type NewScheduleItemReq struct {
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	AuthorID  string    `json:"authorID" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate   time.Time `json:"endDate" validate:"required"`
}

type NewScheduleItemRes struct {
	IsError 	bool   	`json:"isError"`
	StatusCode	int		`json:"statusCode"`	
	Message  	string 	`json:"message"`
}

type GetScheduleItemReq struct {
	ID	string	`json:"id" validate:"required"`
}

type GetScheduleItemRes struct {
	IsError 	bool   						`json:"isError"`
	StatusCode	int							`json:"statusCode"`	
	Message  	string 						`json:"message"`
	Schedule	*model.SchedulePostModel	`json:"schedule"`
}

type ListScheduleItemRes struct {
	IsError 		bool   						`json:"isError"`
	StatusCode		int							`json:"statusCode"`	
	Message  		string 						`json:"message"`
	ScheduleCount	int							`json:"scheduleCount"`
	ScheduleList	[]*model.SchedulePostModel	`json:"scheduleList"`
}

type UpdateScheduleItemReq struct {
	ID			string		`json:"id" validate:"required"`
	Title     	string    	`json:"title" validate:"required"`
	Content   	string    	`json:"content" validate:"required"`
	Author    	string    	`json:"author" validate:"required"`
	AuthorID  	string    	`json:"authorID" validate:"required"`
	StartDate 	time.Time 	`json:"startDate" validate:"required"`
	EndDate   	time.Time 	`json:"endDate" validate:"required"`
}

type UpdateScheduleItemRes struct {
	IsError 	bool   						`json:"isError"`
	StatusCode	int							`json:"statusCode"`	
	Message  	string 						`json:"message"`
	Schedule	*model.SchedulePostModel	`json:"schedule"`
}

type DeleteScheduleItemReq struct {
	ID			string		`json:"id" validate:"required"`
}

type DeleteScheduleItemRes struct {
	IsError 	bool   						`json:"isError"`
	StatusCode	int							`json:"statusCode"`	
	Message  	string 						`json:"message"`
}