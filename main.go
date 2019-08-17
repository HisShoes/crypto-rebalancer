package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/hisshoes/crypto-rebalancer/pkg/http/rest"
	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
	"github.com/hisshoes/crypto-rebalancer/pkg/storage/memory"
	"github.com/hisshoes/crypto-rebalancer/pkg/storage/mongodb"
)

func main() {
	dbType := "mongodb"

	var p portfolio.Service

	//setup storage and initialize service with it
	switch dbType {
	case "memory":
		s := new(memory.Storage)
		p = portfolio.NewService(s)

	case "mongodb":
		s := mongodb.NewRepository("mongodb://192.168.99.100:27017")
		p = portfolio.NewService(s)

	default:
		panic("Unknown database")
	}

	//create router using service
	router := rest.Handler(p)

	//initialize logger
	log.SetLevel(log.InfoLevel)

	//start server
	log.Info("Portfolio server starting on ", ":8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
