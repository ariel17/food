package repositories

import (
	"database/sql"
	"github.com/ariel17/food/internal/entities"
	"strings"
)

type IngredientsRepositoryMySQL struct {
	db *sql.DB
}

func NewIngredientsRepositoryMySQL(db *sql.DB) *IngredientsRepositoryMySQL {
	return &IngredientsRepositoryMySQL{
		db: db,
	}
}

func (i *IngredientsRepositoryMySQL) JoinPlatesSteps(plates []entities.Plate) ([]entities.Step, error) {
	ids := []interface{}{}
	for _, plate := range plates {
		ids = append(ids, plate.ID)
	}
	sql := "SELECT i.name, SUM(pi.amount), pi.unit FROM plates_ingredients pi INNER JOIN ingredients i ON (i.id=pi.ingredient_id) WHERE pi.plate_id IN (?" + strings.Repeat(",?", len(plates)-1) + ") GROUP BY i.name, pi.unit"
	var rows, err = i.db.Query(sql, ids...)
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
