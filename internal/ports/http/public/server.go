package v1

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
	s.router.Get("/v1/getmaxrate", s.GetMaxRate)
	s.router.Get("/v1/getminrate", s.GetMinRate)
	s.router.Get("/v1/getavgrate", s.GetAvgRate)
	s.router.Get("/v1/getlastrate", s.GetLastRate)
	if err := http.ListenAndServe(":8080", s.router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} // как лучше обработать ошибку?
}

// @Summary Get max coin rates
// @Description Returns a list of coins with their max rates in USD
// @Tags coins
// @Accept json
// @Produce json
// @Param titles query string true "List of coin titles (separated by commas, for example: BTC,ETH)"
// @Success 200 {object} dto.CoinsDTO "Coin list"
// @Failure 400 {object} map[string]string "Request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1 [get]
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

	var data dto.CoinsDTO // посмотреть про дто, почему не используем entities
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

// @Summary Get min coin rates
// @Description Returns a list of coins with their min rates in USD
// @Tags coins
// @Accept json
// @Produce json
// @Param titles query string true "List of coin titles (separated by commas, for example: BTC,ETH)"
// @Success 200 {array} entities.Coin "Coin list"
// @Failure 400 {object} map[string]string "Request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1 [get]
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

// @Summary Get average coin rates
// @Description Returns a list of coins with their average rates in USD
// @Tags coins
// @Accept json
// @Produce json
// @Param titles query string true "List of coin titles (separated by commas, for example: BTC,ETH)"
// @Success 200 {array} entities.Coin "Coin list"
// @Failure 400 {object} map[string]string "Request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1 [get]
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

// @Summary Get last coin rates
// @Description Returns a list of coins with their latest rates in USD
// @Tags coins
// @Accept json
// @Produce json
// @Param titles query string true "List of coin titles (separated by commas, for example: BTC,ETH)"
// @Success 200 {array} entities.Coin "Coin list"
// @Failure 400 {object} map[string]string "Request error"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1 [get]
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
