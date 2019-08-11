package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hisshoes/crypto-rebalancer/pkg/http/rest"
	"github.com/hisshoes/crypto-rebalancer/pkg/portfolio"
	"github.com/hisshoes/crypto-rebalancer/pkg/storage/memory"
)

func main() {
	dbType := "memory"

	var p portfolio.Service

	switch dbType {
	case "memory":
		s := new(memory.Storage)
		p = portfolio.NewService(s)

	default:
		panic("Unknown database")
	}

	router := rest.Handler(p)

	fmt.Println("Portfolio server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}
