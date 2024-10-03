package repository

import (
	"context"

	"go-fiber-docker-example/internal/model"
	"go-fiber-docker-example/internal/dto"
)

type ScheduleRepository struct {
	client *model.PrismaClient
}

func NewScheduleRepository(prismaClient *model.PrismaClient) *ScheduleRepository {
	return &ScheduleRepository{client: prismaClient}
}

func (r *ScheduleRepository) CreateSchedule(req *dto.NewScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := r.client.SchedulePost.CreateOne(
		model.SchedulePost.Title.Set(req.Title),
		model.SchedulePost.Content.Set(req.Content),
		model.SchedulePost.Author.Set(req.Author),
		model.SchedulePost.AuthorID.Set(req.AuthorID),
		model.SchedulePost.StartDate.Set(req.StartDate),
		model.SchedulePost.EndDate.Set(req.EndDate),
	).Exec(context.Background())
	return schedule, err
}

func (r *ScheduleRepository) GetScheduleByID(req *dto.GetScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := r.client.SchedulePost.FindUnique(
		model.SchedulePost.ID.Equals(req.ID),
	).Exec(context.Background())
	return schedule, err
}

func (r *ScheduleRepository) UpdateSchedule(req *dto.UpdateScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := r.client.SchedulePost.FindUnique(
		model.SchedulePost.ID.Equals(req.ID),
	).Update(
		model.SchedulePost.Title.Set(req.Title),
		model.SchedulePost.Content.Set(req.Content),
		model.SchedulePost.Author.Set(req.Author),
		model.SchedulePost.StartDate.Set(req.StartDate),
		model.SchedulePost.EndDate.Set(req.EndDate),
	).Exec(context.Background())
	return schedule, err
}

func (r *ScheduleRepository) DeleteSchedule(req *dto.DeleteScheduleItemReq) (bool, error) {
	_, err := r.client.SchedulePost.FindUnique(
		model.SchedulePost.ID.Equals(req.ID),
	).Delete().Exec(context.Background())
	
	if err != nil {
		return true, err
	}
	return false, nil
}