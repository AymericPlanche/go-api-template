package api

import (
	"encoding/json"
	"io"
	"myapp/internal/app/response"
	"myapp/internal/things"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type ThingsHandler struct {
	ThingsRepo things.Repository
	Logger     zerolog.Logger
}

// handler for GET /things
func (h ThingsHandler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		things, err := h.ThingsRepo.LoadAll(r.Context())
		if err != nil {
			h.Logger.Error().Err(err).Msg("error while loading things")
			response.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Success(r.Context(), w, things, http.StatusOK)
	}
}

// handler for GET /thing/{id}
func (h ThingsHandler) Details() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		h.Logger.Info().Str("id", id).Msg("retrieving details")

		thing, err := h.ThingsRepo.Load(r.Context(), things.ID(id))
		if err != nil {
			if err == things.ErrNotFound {
				response.Error(w, "not found", http.StatusNotFound)
				return
			}
			response.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Success(r.Context(), w, thing, http.StatusOK)
	}
}

// handler for POST /thing
func (h ThingsHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var thing things.Thing
		if err := json.Unmarshal(body, &thing); err != nil {
			response.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.ThingsRepo.Create(r.Context(), thing); err != nil {
			if err == things.ErrAlreadyExists {
				response.Error(w, err.Error(), http.StatusConflict)
				return
			}
			response.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Success(r.Context(), w, thing, http.StatusCreated)
	}
}
