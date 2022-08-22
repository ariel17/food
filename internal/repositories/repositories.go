package repositories

import "github.com/ariel17/food/internal/entities"

type Repository interface {
	JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error)
	GetStepsForPlate(plate entities.Plate) ([]entities.Step, error)
	GetAllPlates() ([]entities.Plate, error)
	Close()
}

// New creates a new repository instance based on indicated recipes source.
func New(source string) Repository {
	if source == "database" {
		return NewRepositoryMySQL()
	}
	return NewRepositoryYAML()
}