package repositories

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ariel17/food/internal/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJoinPlatesIngredients(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	plates := []entities.Plate{
		{
			ID: 1,
			Name: "milanesa con pure",
		},
		{
			ID: 2,
			Name: "pollo con papas al horno",
		},
	}

	r := NewIngredientsRepositoryMySQL(db)

	testCases := []struct{
		name string
		isSuccessful bool
	}{
		{"ok", true},
		{"failed by database error", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sql := `SELECT i.name, SUM(pi.amount), pi.unit FROM plates_ingredients pi INNER JOIN ingredients i ON (i.id=pi.ingredient_id) WHERE pi.plate_id IN (?,?) GROUP BY i.name, pi.unit`
			eq := mock.ExpectQuery(sql).WithArgs(1, 2)
			if tc.isSuccessful {
				columns := []string{"name", "amount", "unit"}
				eq.WillReturnRows(sqlmock.NewRows(columns).
					AddRow("carne", 1000, "g").
					AddRow("papa", 2000, "g").
					AddRow("pollo", 500, "g"))
			} else {
				eq.WillReturnError(errors.New("mocked error"))
			}

			steps, err := r.JoinPlatesSteps(plates)

			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotEmpty(t, steps)
				assert.Equal(t, "carne", steps[0].Ingredient.Name)
				assert.Equal(t, float64(1000), steps[0].Amount)
				assert.Equal(t, "papa", steps[1].Ingredient.Name)
				assert.Equal(t, float64(2000), steps[1].Amount)
				assert.Equal(t, "pollo", steps[2].Ingredient.Name)
				assert.Equal(t, float64(500), steps[2].Amount)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, steps)
			}
		})
	}
}
