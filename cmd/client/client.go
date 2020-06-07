package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/beruangcoklat/share-variables/internal/usecase"

	"github.com/beruangcoklat/share-variables/internal/handler"
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	h := handler.GetHandler()
	r.HandleFunc("/store-variable", h.StoreVariable)
	return r
}

func main() {
	port := flag.String("port", "", "port")
	flag.Parse()

	go usecase.GetUseCase().PrintVariable(context.Background())
	go usecase.GetUseCase().ReceiveVariable(context.Background())

	err := http.ListenAndServe(":"+*port, router())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
