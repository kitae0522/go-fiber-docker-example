package dto

import "time"

type ScheduleItem struct {
	ID			string 		`json:"id"`
	Title		string		`json:"title"`
	content		string		`json:"content"`
	author		string 		`json:"author"`
	authorID	string 		`json:"authorID"`
	startDate	time.Time	`json:"startDate"`
	endDate		time.Time	`json:"endDate"`
	createdAt	time.Time	`json:"createdAt"`
	updatedAt	time.Time	`json:"updatedAt"`
}

type NewScheduleItemReq struct {
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    string    `json:"author" validate:"required"`
	AuthorID  string    `json:"authorID" validate:"required"`
	StartDate time.Time `json:"startDate" validate:"required"`
	EndDate   time.Time `json:"endDate" validate:"required"`
}

type NewScheduleItemRes struct {
	IsError 	bool   `json:"isError"`
	Message  	string `json:"message"`
}

type GetScheduleItemReq struct {
	ID	string	`json:"id" validate:"required"`
}

type GetScheduleItemRes struct {
	IsError 	bool   			`json:"isError"`
	Message  	string 			`json:"message"`
	Schedule	interface{}	`json:"schedule"`
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
	IsError 	bool   			`json:"isError"`
	Message  	string 			`json:"message"`
	Schedule	ScheduleItem	`json:"schedule"`
}

type DeleteScheduleItemReq struct {
	ID			string		`json:"id" validate:"required"`
}

type DeleteScheduleItemRes struct {
	IsError 	bool   			`json:"isError"`
	Message  	string 			`json:"message"`
	Schedule	ScheduleItem	`json:"schedule"`
}