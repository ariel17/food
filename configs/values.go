package configs

import (
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	host     string
	port     int
	username string
	password string
	name     string
}

func (d DatabaseConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.username, d.password, d.host, d.port, d.name)
}

var databaseConfig DatabaseConfig

func GetDatabaseConfig() DatabaseConfig {
	return databaseConfig
}

func getInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}
	return value
}

func init() {
	databaseConfig = DatabaseConfig{
		host:     os.Getenv("DATABASE_HOST"),
		port:     getInt("DATABASE_PORT"),
		username: os.Getenv("DATABASE_USER"),
		password: os.Getenv("DATABASE_PASS"),
		name:     os.Getenv("DATABASE_NAME"),
	}
}
