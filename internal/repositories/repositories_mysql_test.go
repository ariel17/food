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
	if err != nil {
		t.Fatal(err)
	}

	plates := []entities.Plate{
		{
			ID:   1,
			Name: "milanesa con pure",
		},
		{
			ID:   2,
			Name: "pollo con papas al horno",
		},
	}

	r := repositoryMySQL{
		db: db,
	}

	testCases := []struct {
		name         string
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
					AddRow("pollo", 500, "g").
					AddRow("palta", 1, nil))
			} else {
				eq.WillReturnError(errors.New("mocked error"))
			}

			steps, err := r.JoinPlatesSteps(plates)

			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotEmpty(t, steps)
				assert.Equal(t, "carne", steps[0].Ingredient.Name)
				assert.Equal(t, float64(1000), steps[0].Amount)
				assert.Equal(t, "g", steps[0].Unit)
				assert.Equal(t, "papa", steps[1].Ingredient.Name)
				assert.Equal(t, float64(2000), steps[1].Amount)
				assert.Equal(t, "g", steps[1].Unit)
				assert.Equal(t, "pollo", steps[2].Ingredient.Name)
				assert.Equal(t, float64(500), steps[2].Amount)
				assert.Equal(t, "g", steps[2].Unit)
				assert.Equal(t, "palta", steps[3].Ingredient.Name)
				assert.Equal(t, float64(1), steps[3].Amount)
				assert.Empty(t, steps[3].Unit)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, steps)
			}
		})
	}
}

func TestGetStepsForPlate(t *testing.T) {
	p := entities.Plate{
		ID: 1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	r := repositoryMySQL{
		db: db,
	}
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
				assert.Equal(t, "comino", steps[0].Ingredient.Name)
				assert.Equal(t, "condimento", steps[0].Ingredient.Type)
				assert.Equal(t, float64(100), steps[0].Amount)
				assert.Equal(t, "g", steps[0].Unit)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, steps)
			}
		})
	}
}

func TestGetAllPlates(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	r := repositoryMySQL{
		db: db,
	}

	testCases := []struct {
		name         string
		isSuccessful bool
	}{
		{"ok", true},
		{"failed by database error", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sql := "SELECT id, name, only_on, needs_mixing FROM plates"
			eq := mock.ExpectQuery(sql)
			if tc.isSuccessful {
				columns := []string{"id", "name", "only_on", "needs_mixing"}
				eq.WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "milanesa", nil, true))
			} else {
				eq.WillReturnError(errors.New("mocked error"))
			}

			plates, err := r.GetAllPlates()

			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotEmpty(t, plates)
				assert.Equal(t, 1, plates[0].ID)
				assert.Equal(t, "milanesa", plates[0].Name)
				assert.Equal(t, "", plates[0].OnlyOn)
				assert.True(t, plates[0].NeedsMixing)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, plates)
			}
		})
	}
}
