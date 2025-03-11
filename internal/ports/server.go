package ports

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
)

type Server interface {
	Run(ctx context.Context) error
}

type server struct {
	router  *chi.Mux
	service Service
}

func NewServer(service Service) Server {
	r := chi.NewRouter()
	s := &server{
		router:  r,
		service: service,
	}
	s.setupRoutes()
	return s
}

func (s *server) setupRoutes() {
	s.router.Get("/rates/max", s.handleGetMaxRate)
	s.router.Get("/rates/min", s.handleGetMinRate)
	s.router.Get("/rates/avg", s.handleGetAvgRate)
	s.router.Get("/rates/last", s.handleGetLastRate)
}

func (s *server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}
	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("error shutting down http server: %v", err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "failed to start server")
	}
	return nil
}

func (s *server) handleGetMaxRate(w http.ResponseWriter, r *http.Request) {
	titlesStr := r.URL.Query().Get("titles")
	if len(titlesStr) == 0 {
		http.Error(w, "titles can't be empty", http.StatusBadRequest)
		return
	}
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetMaxRate(r.Context(), titles)
	if err != nil {
		http.Error(w, "Failed to get max rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coins); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *server) handleGetMinRate(w http.ResponseWriter, r *http.Request) {
	titlesStr := r.URL.Query().Get("titles")
	if len(titlesStr) == 0 {
		http.Error(w, "titles can't be empty", http.StatusBadRequest)
		return
	}
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetMinRate(r.Context(), titles)
	if err != nil {
		http.Error(w, "Failed to get min rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coins); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *server) handleGetAvgRate(w http.ResponseWriter, r *http.Request) {
	titlesStr := r.URL.Query().Get("titles")
	if len(titlesStr) == 0 {
		http.Error(w, "titles can't be empty", http.StatusBadRequest)
		return
	}
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetAvgRate(r.Context(), titles)
	if err != nil {
		http.Error(w, "Failed to get average rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coins); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (s *server) handleGetLastRate(w http.ResponseWriter, r *http.Request) {
	titlesStr := r.URL.Query().Get("titles")
	if len(titlesStr) == 0 {
		http.Error(w, "titles can't be empty", http.StatusBadRequest)
		return
	}
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetLastRate(r.Context(), titles)
	if err != nil {
		http.Error(w, "Failed to get last rate", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(coins); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
