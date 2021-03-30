package repositories

import "github.com/ariel17/food/internal/entities"

type MockRepository struct {
	Plates []entities.Plate
	Steps []entities.Step
	Err error
}

func (m *MockRepository) JoinPlatesSteps(_ []entities.Plate) ([]entities.Step, error) {
	return m.Steps, m.Err
}

func (m *MockRepository) GetStepsForPlate(_ entities.Plate) ([]entities.Step, error) {
	return m.Steps, m.Err
}

func (m *MockRepository) GetAllPlates() ([]entities.Plate, error) {
	return m.Plates, m.Err
}