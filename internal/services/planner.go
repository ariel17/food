package services

import (
	"errors"
	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/repositories"
	"math/rand"
)

const (
	DaysOfWeek = 7
)

type PlannerService interface {
	PlanForTheWeek() ([]entities.Plate, error)
}

func NewPlannerService(repository repositories.Repository) PlannerService {
	return &service{
		repository: repository,
	}
}

type service struct {
	repository repositories.Repository
}

func (s *service) PlanForTheWeek() ([]entities.Plate, error) {
	plates, err := s.repository.GetAllPlates()
	if err != nil {
		return nil, err
	}
	if len(plates) < DaysOfWeek {
		return nil, errors.New("not enough plates to plan for the whole week")
	}
	plan := []entities.Plate{}
	for day := 0; day < DaysOfWeek; day++ {
		v := rand.Intn(len(plates))
		plan = append(plan, plates[v])
		plates = append(plates[:v], plates[v+1:]...)
	}
	return plan, nil
}
