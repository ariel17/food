package repositories

import (
	"database/sql"
	"github.com/ariel17/food/internal/entities"
	"strings"
)

type Repository interface {
	JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error)
	GetStepsForPlate(plate entities.Plate) ([]entities.Step, error)
	GetAllPlates() ([]entities.Plate, error)
}

func NewRepositoryMySQL(db *sql.DB) Repository {
	return &repositoryMySQL{
		db: db,
	}
}

type repositoryMySQL struct {
	db *sql.DB
}

func (r *repositoryMySQL) JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error) {
	ids := make([]interface{}, 0)
	for _, plate := range plates {
		ids = append(ids, plate.ID)
	}
	query := "SELECT i.name, SUM(pi.amount), pi.unit FROM plates_ingredients pi INNER JOIN ingredients i ON (i.id=pi.ingredient_id) WHERE pi.plate_id IN (?" + strings.Repeat(",?", len(plates)-1) + ") GROUP BY i.name, pi.unit"
	var rows, err = r.db.Query(query, ids...)
	if err != nil {
		return nil, err
	}
	steps := make([]entities.Step, 0)
	for rows.Next() {
		var (
			name    string
			amount  float64
			rawUnit *string
		)
		if err := rows.Scan(&name, &amount, &rawUnit); err != nil {
			return nil, err
		}
		var unit string
		if rawUnit != nil {
			unit = *rawUnit
		}
		steps = append(steps, entities.Step{
			Ingredient: entities.Ingredient{
				Name: name,
			},
			Amount: amount,
			Unit:   unit,
		})
	}
	return steps, nil
}

func (r *repositoryMySQL) GetStepsForPlate(plate entities.Plate) ([]entities.Step, error) {
	rows, err := r.db.Query("SELECT ingredient_id, amount, unit FROM plates_ingredients WHERE plate_id = ?", plate.ID)
	if err != nil {
		return nil, err
	}
	steps := make([]entities.Step, 0)
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
			Amount: amount,
			Unit:   unit,
		})
	}
	for i, step := range steps {
		rows, err := r.db.Query("SELECT name, type FROM ingredients WHERE id = ?", step.Ingredient.ID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var name, iType string
			if err := rows.Scan(&name, &iType); err != nil {
				return nil, err
			}
			steps[i].Ingredient.Name = name
			steps[i].Ingredient.Type = iType
		}
	}
	return steps, nil
}

func (r *repositoryMySQL) GetAllPlates() ([]entities.Plate, error) {
	rows, err := r.db.Query("SELECT id, name, only_on, needs_mixing FROM plates")
	if err != nil {
		return nil, err
	}
	plates := make([]entities.Plate, 0)
	for rows.Next() {
		var (
			id        int
			name      string
			rawOnlyOn *string
			needsMixing bool
		)
		if err := rows.Scan(&id, &name, &rawOnlyOn, &needsMixing); err != nil {
			return nil, err
		}
		var onlyOn string
		if rawOnlyOn != nil {
			onlyOn = *rawOnlyOn
		}
		plates = append(plates, entities.Plate{
			ID:     id,
			Name:   name,
			OnlyOn: onlyOn,
			NeedsMixing: needsMixing,
		})
	}
	return plates, nil
}
