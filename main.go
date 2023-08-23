package main

import (
	"myapp/internal/app"
	"os"
	"strings"
)

func main() {
	app.Start(app.Config{
		Port:        env("PORT"),
		Environment: env("ENVIRONMENT"),
		Database: app.DatabaseConfig{
			Username: env("DATABASE_USER"),
			Password: env("DATABASE_PASSWORD"),
			Host:     env("DATABASE_HOST"),
			Port:     env("DATABASE_PORT"),
			DBName:   env("DATABASE_NAME"),
		},
	})
}

// Some env vars retrieved from secret manager have unexpected new lines
func env(env string) string {
	return strings.Trim(os.Getenv(env), " \n")
}
