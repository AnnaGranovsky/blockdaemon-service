package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AnnaGranovsky/blockdaemon-service/blockchain"
	"github.com/go-chi/chi"
)

func (a *API) blockchainRoutes(r chi.Router) {
	r.Get("/", a.listBlockchains)
	r.Post("/", a.createBlockchain)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", a.oneBlockchain)
		r.Post("/block", a.createBlock)
	})
}

func (a *API) listBlockchains(w http.ResponseWriter, r *http.Request) {
	response(200, a.blockchain.List(), w)
}

func (a *API) oneBlockchain(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("no blockchain id param")
		response(400, "Blockchain id should be provided", w)
		return
	}

	resp := map[string]interface{}{
		"blocks":     a.block.List(id),
		"blockchain": a.blockchain.One(id),
	}

	response(200, resp, w)
}

func (a *API) createBlockchain(w http.ResponseWriter, r *http.Request) {
	newbc := blockchain.Blockchain{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newbc); err != nil {
		log.Println("failed to decode r.Body, err: ", err)
		response(400, "Unable to parse JSON", w)
		return
	}

	created, err := a.blockchain.Insert(newbc)
	if err != nil {
		log.Println("insert bc failed, err: ", err)
		response(400, "Failed to create blockchain", w)
		return
	}

	response(200, created, w)

}

func (a *API) createBlock(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		log.Println("no blockchain id param")
		response(400, "blockchain id should be provided", w)
		return
	}

	if err := a.blockchain.IncrementBlocks(id); err != nil {
		log.Println("failed to update blocks count, err: ", err)
		response(400, "Unable to create block", w)
		return
	}

	response(200, a.block.Insert(id), w)
}
