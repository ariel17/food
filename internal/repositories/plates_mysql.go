package repositories

import (
	"database/sql"
	"github.com/ariel17/food/internal/entities"
)

type PlatesRepositoryMySQL struct {
	db *sql.DB
}

func NewPlatesRepositoryMySQL(db *sql.DB) *PlatesRepositoryMySQL {
	return &PlatesRepositoryMySQL{
		db: db,
	}
}

func (s *PlatesRepositoryMySQL) GetAllPlates() ([]entities.Plate, error) {
	rows, err := s.db.Query("SELECT id, name, only_on FROM plates")
	if err != nil {
		return nil, err
	}
	plates := []entities.Plate{}
	for rows.Next() {
		var (
			id int
			name string
			rawOnlyOn *string
		)
		if err := rows.Scan(&id, &name, &rawOnlyOn); err != nil {
			return nil, err
		}
		var onlyOn string
		if rawOnlyOn != nil {
			onlyOn = *rawOnlyOn
		}
		plates = append(plates, entities.Plate{
			ID: id,
			Name: name,
			OnlyOn: onlyOn,
		})
	}
	return plates, nil
}
