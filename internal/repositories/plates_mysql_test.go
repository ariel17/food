package repositories

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllPlates(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	if err != nil {
		t.Fatal(err)
	}

	r := NewPlatesRepositoryMySQL(db)

	testCases := []struct{
		name string
		isSuccessful bool
	}{
		{"ok", true},
		{"failed by database error", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sql := "SELECT id, name, only_on FROM plates"
			eq := mock.ExpectQuery(sql)
			if tc.isSuccessful {
				columns := []string{"id", "name", "only_on"}
				eq.WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "milanesa", nil))
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
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, plates)
			}
		})
	}
}
