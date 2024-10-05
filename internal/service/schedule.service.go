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

func (s *ScheduleService) ListSchedule() ([]*model.SchedulePostModel, error) {
	scheduleList, err := s.scheduleRepo.ListSchedule()
	return scheduleList, err
}

func (s *ScheduleService) GetScheduleByID(req *dto.GetScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := s.scheduleRepo.GetScheduleByID(req)
	return schedule, err
}

func (s *ScheduleService) UpdateSchedule(req *dto.UpdateScheduleItemReq) (*model.SchedulePostModel, error) {
	schedule, err := s.scheduleRepo.UpdateSchedule(req)
	return schedule, err
}

func (s *ScheduleService) DeleteSchedule(req *dto.DeleteScheduleItemReq) (bool, error) {
	success, err := s.scheduleRepo.DeleteSchedule(req)
	return success, err
}