package main

import (
	"myapp/internal/app"
	"os"
	"strings"
)

func main() {
	app.Start(app.Config{
		Port:                  env("PORT"),
		Environment:           env("ENVIRONMENT"),
		Version:               env("VERSION"),
		ProjectID:             env("PROJECT_ID"),
		FirestoreEmulatorHost: env("FIRESTORE_EMULATOR_HOST"),
		CSRFSecret:            env("CSRF_SECRET"),
		SessionPassphrase:     env("SESSION_PASSPHRASE"),
		AuthCallbackHost:      env("AUTH_CALLBACK_HOST"),
		GoogleClientKey:       env("GOOGLE_KEY"),
		GoogleSecret:          env("GOOGLE_SECRET"),
		Buckets:               strings.Split(env("GCS_BUCKETS"), "|"),
		UserWhitelist:         strings.Split(env("USER_WHITE_LIST"), "|"),
	})
}

// Some env vars retrieved from secret manager have unexpected new lines
func env(env string) string {
	return strings.Trim(os.Getenv(env), " \n")
}
