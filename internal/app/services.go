package app

import (
	"database/sql"
	"fmt"
	"myapp/internal/app/handlers/api"
	"myapp/internal/things"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
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
// we don't want to start an half baked application
func InitServices(cfg Config) Services {
	logger := initLogger(cfg)
	db := initDB(cfg.Database, logger)

	return Services{
		Logger: logger,
		APIHandlers: APIHandlers{
			Things: api.ThingsHandler{
				Logger:     logger,
				ThingsRepo: things.NewRepository(logger, db),
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

func initDB(cfg DatabaseConfig, logger zerolog.Logger) *goqu.Database {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	cnn, err := sql.Open("pgx", url)
	if err != nil {
		panic(err)
	}

	db := goqu.Dialect("postgres").DB(cnn)
	db.Logger(&logger)
	return db

}
