package config

import "os"

func GetDatabaseConnectionString() string {
	return os.Getenv("DATABASE_URL")
}
