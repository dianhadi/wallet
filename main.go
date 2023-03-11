package main

import (
	"log"
	"net/http"

	"github.com/dianhadi/wallet/internal/config"
	hHelper "github.com/dianhadi/wallet/internal/handler/helper"
	hTransaction "github.com/dianhadi/wallet/internal/handler/transaction"
	hWallet "github.com/dianhadi/wallet/internal/handler/wallet"

	rTransaction "github.com/dianhadi/wallet/internal/repo/transaction"
	rWallet "github.com/dianhadi/wallet/internal/repo/wallet"
	"github.com/dianhadi/wallet/pkg/database"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("Get Configuration")
	config, err := config.GetConfig("config/config.yaml")
	if err != nil {
		panic(err)
	}

	log.Println("Connect to Database")
	db := database.New(config.Database.Username, config.Database.Password, config.Database.Host, config.Database.DBName, config.Database.Port)

	log.Println("Init Repo")
	repoWallet, err := rWallet.New(db)
	if err != nil {
		panic(err)
	}
	repoTransaction, err := rTransaction.New(db)
	if err != nil {
		panic(err)
	}

	log.Println("Init Handler")
	handlerHelper, err := hHelper.New(repoWallet)
	if err != nil {
		panic(err)
	}
	handlerWallet, err := hWallet.New(repoWallet)
	if err != nil {
		panic(err)
	}
	handlerTransaction, err := hTransaction.New(repoTransaction, repoWallet)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	authToken := chi.Chain(handlerHelper.Authorize)

	r.Post("/api/v1/init", handlerWallet.Init)

	r.With(authToken...).Post("/api/v1/wallet", handlerWallet.Enable)
	r.With(authToken...).Get("/api/v1/wallet", handlerWallet.View)
	r.With(authToken...).Patch("/api/v1/wallet", handlerWallet.Disable)

	r.With(authToken...).Get("/api/v1/wallet/transactions", handlerTransaction.View)
	r.With(authToken...).Post("/api/v1/wallet/deposits", handlerTransaction.Deposit)
	r.With(authToken...).Post("/api/v1/wallet/withdrawals", handlerTransaction.Withdraw)

	log.Println("Service is up at port", config.Server.Port)
	http.ListenAndServe(config.Server.Port, r)
}
