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
	})
}

// Some env vars retrieved from secret manager have unexpected new lines
func env(env string) string {
	return strings.Trim(os.Getenv(env), " \n")
}
