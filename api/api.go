package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AnnaGranovsky/blockdaemon-service/block"
	"github.com/AnnaGranovsky/blockdaemon-service/blockchain"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// API - type for dependency injection
type API struct {
	blockchain *blockchain.Manager
	block      *block.Manager
}

// New - init API
func New() *API {
	return &API{
		blockchain: blockchain.New(),
		block:      block.New(),
	}
}

// InitRouter init application routes
func (a *API) InitRouter(withLogger bool) *chi.Mux {
	r := chi.NewRouter()

	if withLogger {
		r.Use(middleware.Logger)
	}

	r.Use(middleware.Recoverer)

	r.Route("/blockchain", func(r chi.Router) {
		r.Route("/", a.blockchainRoutes)
	})

	return r
}

func response(status int, data interface{}, w http.ResponseWriter) {
	(*&w).Header().Set("Access-Control-Allow-Origin", "*")
	if status == 0 {
		status = http.StatusOK
	}
	w.WriteHeader(status)

	e := json.NewEncoder(w)
	if err := e.Encode(data); err != nil {
		log.Println("Encoding response failed:", err)
		http.Error(w, "Encoding error occured", http.StatusInternalServerError)
		return
	}
}
