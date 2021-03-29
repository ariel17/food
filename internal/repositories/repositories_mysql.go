package repositories

import (
	"database/sql"
	"github.com/ariel17/food/internal/entities"
	"strings"
)

type RepositoryMySQL struct {
	db *sql.DB
}

func (r *RepositoryMySQL) JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error) {
	ids := []interface{}{}
	for _, plate := range plates {
		ids = append(ids, plate.ID)
	}
	sql := "SELECT i.name, SUM(pi.amount), pi.unit FROM plates_ingredients pi INNER JOIN ingredients i ON (i.id=pi.ingredient_id) WHERE pi.plate_id IN (?" + strings.Repeat(",?", len(plates)-1) + ") GROUP BY i.name, pi.unit"
	var rows, err = r.db.Query(sql, ids...)
	if err != nil {
		return nil, err
	}
	steps := []entities.Step{}
	for rows.Next() {
		var (
			name   string
			amount float64
			unit   string
		)
		if err := rows.Scan(&name, &amount, &unit); err != nil {
			return nil, err
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

func (r *RepositoryMySQL) GetStepsForPlate(plate entities.Plate) ([]entities.Step, error) {
	rows, err := r.db.Query("SELECT ingredient_id, amount, unit FROM plates_ingredients WHERE plate_id = ?", plate.ID)
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
		rows, err := r.db.Query("SELECT name, type FROM ingredients WHERE id = ?", step.Ingredient.ID)
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
