package repositories

import (
	"errors"
	"github.com/ariel17/food/internal/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetStepsForPlate(t *testing.T) {
	p := entities.Plate{
		ID:   1,
	}
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	r := NewStepsRepositoryMySQL(db)
	testCases := []struct {
		name         string
		isSuccessful bool
	}{
		{"ok", true},
		{"failed by database error on query", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sql := "SELECT ingredient_id, amount, unit FROM plates_ingredients WHERE plate_id = ?"
			eq1 := mock.ExpectQuery(sql).WithArgs(p.ID)

			sql = "SELECT name, type FROM ingredients WHERE id = ?"
			eq2 := mock.ExpectQuery(sql).WithArgs(2)

			if tc.isSuccessful {
				columns := []string{"ingredient_id", "amount", "unit"}
				eq1.WillReturnRows(sqlmock.NewRows(columns).AddRow(2, 100, "g"))

				columns = []string{"name", "type"}
				eq2.WillReturnRows(sqlmock.NewRows(columns).AddRow("comino", "condimento"))
			} else {
				eq1.WillReturnError(errors.New("mocked error"))
				eq2.WillReturnError(errors.New("mocked error"))
			}

			steps, err := r.GetStepsForPlate(p)

			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotEmpty(t, steps)
				assert.Equal(t, 2, steps[0].Ingredient.ID)
				assert.Equal(t,"comino", steps[0].Ingredient.Name)
				assert.Equal(t,"condimento", steps[0].Ingredient.Type)
				assert.Equal(t, float64(100), steps[0].Amount)
				assert.Equal(t, "g", steps[0].Unit)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, steps)
			}
		})
	}
}
