package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/repositories"
)

const (
	DaysOfWeek = 7
)

type PlannerService interface {
	CreatePlan() ([]entities.Plate, error)
	CreateShopList(plan []entities.Plate) ([]entities.Step, error)
}

func NewPlannerService(repository repositories.Repository) PlannerService {
	return &service{
		repository: repository,
	}
}

type service struct {
	repository repositories.Repository
}

func (s *service) CreatePlan() ([]entities.Plate, error) {
	plates, err := s.repository.GetAllPlates()
	if err != nil {
		return nil, err
	}
	if len(plates) < DaysOfWeek {
		return nil, errors.New("not enough plates to plan for the whole week")
	}
	plan := make([]entities.Plate, 0)
	for day := 0; day < DaysOfWeek; day++ {
		v := rand.Intn(len(plates))
		plan = append(plan, plates[v])
		plates = append(plates[:v], plates[v+1:]...)
	}
	return plan, nil
}

func (s *service) CreateShopList(plan []entities.Plate) ([]entities.Step, error) {
	return s.repository.JoinPlatesSteps(plan)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
