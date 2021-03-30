package services

import (
	"github.com/ariel17/food/internal/entities"
	"github.com/ariel17/food/internal/repositories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePlan(t *testing.T) {
	testCases := []struct {
		name         string
		isSuccessful bool
	}{
		{"ok", true},
	}

	r := &repositories.MockRepository{}
	r.Plates = []entities.Plate{
		{ID: 1, Name: "milanesa"},
		{ID: 2, Name: "pollo frito"},
		{ID: 3, Name: "pizza"},
		{ID: 4, Name: "sopa"},
		{ID: 5, Name: "ensalada"},
		{ID: 6, Name: "asado"},
		{ID: 7, Name: "sushi"},
		{ID: 8, Name: "puchero"},
		{ID: 9, Name: "empanadas"},
		{ID: 10, Name: "guiso"},
	}
	s := NewPlannerService(r)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			plates, err := s.CreatePlan()
			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotEmpty(t, plates)
				assert.Equal(t, 7, len(plates))

				plates2, err2 := s.CreatePlan()
				assert.Nil(t, err2)
				assert.NotEmpty(t, plates2)
				assert.Equal(t, 7, len(plates2))
				assert.NotEqual(t, plates, plates2)
			} else {
				assert.NotNil(t, err)
				assert.Empty(t, plates)
			}
		})
	}
}
