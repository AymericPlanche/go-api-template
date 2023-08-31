package app

import (
	"context"
	"log"
	"myapp/internal/app/response"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
)

func Start(cfg Config) {
	services := InitServices(cfg)

	log := services.Logger

	log.Info().
		Str("port", cfg.Port).
		Str("env", cfg.Environment).
		Msg("Application is starting")

	server := &http.Server{Addr: ":" + cfg.Port, Handler: Handler(cfg, services)}
	gracefulShutdown(server, services.Logger)
}

func Handler(cfg Config, services Services) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(services.Logger))
	r.Use(middleware.Heartbeat("/healthcheck"))
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(middleware.Recoverer)
	// TODO: auhtentication middleware depending on auth type

	r.Route("/hello", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			response.Success(r.Context(), w, "Hello PxP enthousiasts !", http.StatusOK)
		})
	})

	r.Route("/things", func(r chi.Router) {
		h := services.APIHandlers.Things

		r.Get("/", h.Index())
		r.Get("/{id}", h.Details())
		r.Post("/", h.Create())
	})

	return r
}

func gracefulShutdown(server *http.Server, logger zerolog.Logger) {
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		// there is nothing useful to do with the cancel function at this level, so we can discard it
		_ = cancel

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			logger.Fatal().Err(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
