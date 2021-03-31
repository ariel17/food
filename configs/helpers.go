package configs

import (
	"os"
	"strconv"
)

func getInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return value
}
