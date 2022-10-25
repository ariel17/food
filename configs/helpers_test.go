package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		mustSetValue bool
		value        string
		intValue     int
	}{
		{"key exists and is numeric", "TEST_KEY", true, "10", 10},
		{"key exists and is NOT numeric", "TEST_KEY", true, "xxxyyyzzz", 0},
		{"key NOT exists", "TEST_KEY", false, "", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv(tc.key)
			if tc.mustSetValue {
				os.Setenv(tc.key, tc.value)
			}
			v := getInt(tc.key)
			assert.Equal(t, tc.intValue, v)
		})
	}
}
