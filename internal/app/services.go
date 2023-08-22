package app

import (
	"myapp/internal/app/handlers/api"
	"myapp/internal/things"
	"os"

	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// List of public services
// Services should only be referenced in this struct if they are called directly from the app startup phase
// If not, they stay just a private var in initServices
type Services struct {
	Logger zerolog.Logger

	APIHandlers APIHandlers
}

// Contains the API http handlers defined as services
type APIHandlers struct {
	Things api.ThingsHandler
}

// builts the entire dependency tree
// it is called on app startup and should panic if something is wrong
// we don't want to start and half baked application
func initServices(cfg Config) Services {
	logger := initLogger(cfg)

	return Services{
		Logger: logger,
		APIHandlers: APIHandlers{
			Things: api.ThingsHandler{
				Logger:     logger,
				ThingsRepo: things.NewRepository(logger),
			},
		},
	}
}

func initLogger(cfg Config) zerolog.Logger {
	if cfg.Environment == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	return httplog.NewLogger("myapp", httplog.Options{
		JSON: true,
	})
}
