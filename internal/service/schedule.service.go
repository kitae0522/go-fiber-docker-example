package service

import (
	"go-fiber-docker-example/internal/dto"
	"go-fiber-docker-example/internal/repository"
	"go-fiber-docker-example/internal/model"
)

type ScheduleService struct {
	scheduleRepo *repository.ScheduleRepository
}

func NewScheduleService(repo *repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{scheduleRepo: repo}
}

func (s *ScheduleService) CreateSchedule(req *dto.NewScheduleItemReq) error {
	_, err := s.scheduleRepo.CreateSchedule(req)
	return err
}

func (s *ScheduleService) GetScheduleByID(req *dto.GetScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(req)
	return schedule, err
}