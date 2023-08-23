package integration_tests

import (
	"bytes"
	"io"
	"net/http"

	"myapp/internal/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type AppContext struct {
	Configuration app.Config
	Services      app.Services
	LogBuffer     *bytes.Buffer
}

func InitAppContext() AppContext {
	cfg := initConfiguration()
	logger, buf := initLogger()

	// override a few services
	services := app.InitServices(cfg)
	services.Logger = logger

	return AppContext{
		Configuration: cfg,
		Services:      services,
		LogBuffer:     buf,
	}
}

func initConfiguration() app.Config {
	return app.Config{
		Port:        "8080",
		Environment: "test",
		Database: app.DatabaseConfig{
			Username: "postgres",
			Password: "postgres",
			Host:     "db",
			Port:     "5432",
			DBName:   "integration",
		},
	}
}

func (appContext AppContext) GetHTTPHandler() http.Handler {
	return app.Handler(appContext.Configuration, appContext.Services)
}

func (appContext AppContext) GetLogsAsString() string {
	logs, _ := io.ReadAll(appContext.LogBuffer)
	return string(logs)
}

func initLogger() (zerolog.Logger, *bytes.Buffer) {
	var buf bytes.Buffer

	return log.Output(zerolog.ConsoleWriter{Out: &buf, NoColor: true}), &buf
}
