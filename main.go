package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AnnaGranovsky/blockdaemon-service/api"
)

const port = ":3000"

func main() {
	a := api.New()
	server := &http.Server{
		Addr:           port,
		Handler:        a.InitRouter(true),
		ReadTimeout:    time.Duration(300) * time.Second,
		WriteTimeout:   time.Duration(30) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 Mb
	}

	log.Println("launching the service at", port)

	err := server.ListenAndServe()
	if err != nil && err.Error() == "http: Server closed" {
		log.Println("api port is closed")
		return
	}

	log.Panic(err)
}
