package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
)

//Handler return a handler to handle REST API requests
func Handler(p portfolio.Service) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/portfolio/{id}", getPortfolio(p)).Methods("GET")
	router.HandleFunc("/portfolio", listPortfolios(p)).Methods("GET")
	router.HandleFunc("/portfolio", createPortfolio(p)).Methods("POST")

	return router
}

func createPortfolio(s portfolio.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var p portfolio.Portfolio
		_ = json.NewDecoder(r.Body).Decode(&p)

		id, err := s.CreatePortfolio(p)
		if err != nil {
			http.Error(w, "Couldnt create portfolio", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(id)
	}
}

func listPortfolios(s portfolio.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := s.ListPortfolios()
		if err == portfolio.ErrMissing {
			http.Error(w, "Portfolios could not be retrieved", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}

func getPortfolio(s portfolio.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		p, err := s.Portfolio(id)
		if err == portfolio.ErrMissing {
			http.Error(w, "Portfolio could not be retrieved", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
	}
}
