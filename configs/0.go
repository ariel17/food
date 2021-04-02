package configs

import "os"

var environment string

func IsProduction() bool {
	return environment == "production"
}

func init() {
	environment = os.Getenv("ENVIRONMENT")
}