package public

import (
	"cryptoProject/internal/entities"
	"cryptoProject/pkg/dto"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strings"
)

type Server struct {
	router  *chi.Mux
	service Service
}

func NewServer(service Service) (*Server, error) {
	if service == nil || service == Service(nil) {
		return nil, errors.Wrap(entities.ErrInvalidParameter, "Service can't be nil")
	}
	r := chi.NewRouter()
	s := &Server{
		router:  r,
		service: service,
	}
	return s, nil
}

func (s *Server) Run() {
	s.router.Get("/v1/getmax", s.GetMaxRate)
	s.router.Get("/v1/getmin", s.GetMinRate)
	s.router.Get("/v1/getavg", s.GetAvgRate)
	s.router.Get("/v1/getlast", s.GetLastRate)
	if err := http.ListenAndServe(":8080", s.router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// @Summary		Get max coin rates
// @Description	Returns a list of coins with their max rates in USD
// @Tags			coins
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CoinsDTO	"Coin list"
// @Failure		400
// @Failure		404
// @Failure		500
// @Router			/getmax [get]
// @Param			titles	query	string	true	"List of coin titles (separated by commas, for example: BTC,ETH)"
func (s *Server) GetMaxRate(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	titlesStr := req.URL.Query().Get("titles")
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetMaxRate(ctx, titles)
	if err != nil {
		http.Error(rw, "Failed to get max rate", http.StatusBadRequest)
		return
	}
	if coins == nil {
		http.Error(rw, "No max rate", http.StatusNotFound)
	}

	var data dto.CoinsDTO
	for _, coin := range coins {
		data = append(data, dto.CoinDTO{
			Title:       coin.Title,
			CurrentRate: coin.CurrentRate,
			Timestamp:   coin.Timestamp,
		})
	}

	rw.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// @Summary		Get min coin rates
// @Description	Returns a list of coins with their min rates in USD
// @Tags			coins
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CoinsDTO	"Coin list"
// @Failure		400
// @Failure		404
// @Failure		500
// @Router			/getmin [get]
// @Param			titles	query	string	true	"List of coin titles (separated by commas, for example: BTC,ETH)"
func (s *Server) GetMinRate(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	titlesStr := req.URL.Query().Get("titles")
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetMinRate(ctx, titles)
	if err != nil {
		http.Error(rw, "Failed to get min rate", http.StatusBadRequest)
		return
	}
	if coins == nil {
		http.Error(rw, "No min rate", http.StatusNotFound)
	}

	var data dto.CoinsDTO
	for _, coin := range coins {
		data = append(data, dto.CoinDTO{
			Title:       coin.Title,
			CurrentRate: coin.CurrentRate,
			Timestamp:   coin.Timestamp,
		})
	}

	rw.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// @Summary		Get average coin rates
// @Description	Returns a list of coins with their average rates in USD
// @Tags			coins
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CoinsDTO	"Coin list"
// @Failure		400
// @Failure		404
// @Failure		500
// @Router			/getavg [get]
// @Param			titles	query	string	true	"List of coin titles (separated by commas, for example: BTC,ETH)"
func (s *Server) GetAvgRate(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	titlesStr := req.URL.Query().Get("titles")
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetAvgRate(ctx, titles)
	if err != nil {
		http.Error(rw, "Failed to get average rate", http.StatusBadRequest)
		return
	}
	if coins == nil {
		http.Error(rw, "No average rate", http.StatusNotFound)
	}

	var data dto.CoinsDTO
	for _, coin := range coins {
		data = append(data, dto.CoinDTO{
			Title:       coin.Title,
			CurrentRate: coin.CurrentRate,
			Timestamp:   coin.Timestamp,
		})
	}

	rw.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

// @Summary		Get last coin rates
// @Description	Returns a list of coins with their latest rates in USD
// @Tags			coins
// @Accept			json
// @Produce		json
// @Success		200	{object}	dto.CoinsDTO	"Coin list"
// @Failure		400
// @Failure		404
// @Failure		500
// @Router			/getlast [get]
// @Param			titles	query	string	true	"List of coin titles (separated by commas, for example: BTC,ETH)"
func (s *Server) GetLastRate(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	titlesStr := req.URL.Query().Get("titles")
	titles := strings.Split(titlesStr, ",")

	coins, err := s.service.GetLastRate(ctx, titles)
	if err != nil {
		http.Error(rw, "Failed to get last rate", http.StatusBadRequest)
		return
	}
	if coins == nil {
		http.Error(rw, "No last rate", http.StatusNotFound)
	}

	var data dto.CoinsDTO
	for _, coin := range coins {
		data = append(data, dto.CoinDTO{
			Title:       coin.Title,
			CurrentRate: coin.CurrentRate,
			Timestamp:   coin.Timestamp,
		})
	}

	rw.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, "Failed to encode response", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
