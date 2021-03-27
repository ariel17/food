package fixtures

import (
	"encoding/json"
	"testing"
)

type dummy struct {
	Dummy1 string `json:"dummy_1"`
	Dummy2 int    `json:"dummy_2"`
}

func (d1 dummy) Equals(d2 dummy) bool {
	return d1.Dummy1 == d2.Dummy1 && d1.Dummy2 == d2.Dummy2
}

func TestLoad(t *testing.T) {
	testCases := []struct {
		name         string
		goldenPath   string
		isSuccessful bool
		expected     dummy
	}{
		{"ok", "ok.json", true, dummy{"1", 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			b := goldies.Load(t, tc.goldenPath)
			var d dummy
			err := json.Unmarshal(b, &d)
			assert.Nil(t, err)
			assert.True(t, d.Equals(tc.expected))
		})
	}
}
