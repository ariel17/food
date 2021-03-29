package repositories

import (
	"database/sql"
	"github.com/ariel17/food/internal/entities"
)

type StepsRepositoryMySQL struct {
	db *sql.DB
}

func NewStepsRepositoryMySQL(db *sql.DB) *StepsRepositoryMySQL {
	return &StepsRepositoryMySQL{
		db: db,
	}
}

func (s *StepsRepositoryMySQL) GetStepsForPlate(plate entities.Plate) ([]entities.Step, error) {
	rows, err := s.db.Query("SELECT ingredient_id, amount, unit FROM plates_ingredients WHERE plate_id = ?", plate.ID)
	if err != nil {
		return nil, err
	}
	steps := []entities.Step{}
	for rows.Next() {
		var (
			ingredientID int
			amount       float64
			unit         string
		)
		if err := rows.Scan(&ingredientID, &amount, &unit); err != nil {
			return nil, err
		}
		steps = append(steps, entities.Step{
			Ingredient: entities.Ingredient{
				ID: ingredientID,
			},
			Amount:     amount,
			Unit:       unit,
		})
	}
	for i, step := range steps {
		rows, err := s.db.Query("SELECT name, type FROM ingredients WHERE id = ?", step.Ingredient.ID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var name, itype  string
			if err := rows.Scan(&name, &itype); err != nil {
				return nil, err
			}
			steps[i].Ingredient.Name = name
			steps[i].Ingredient.Type = itype
		}
	}
	return steps, nil
}
