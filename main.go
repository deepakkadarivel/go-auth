package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"go-auth-sample/product"
	"github.com/gorilla/handlers"
	"os"
	"go-auth-sample/jwthandler"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", renderHome).Methods("GET")
	router.Handle("/status", product.StatusHandler).Methods("GET")
	router.Handle("/products", jwthandler.JwtValidator.Handler(product.ProductsHandler)).Methods("GET")
	router.Handle("/products/{slug}/feedback", jwthandler.JwtValidator.Handler(product.AddFeedbackHandler)).Methods("POST")
	router.Handle("/get-token", jwthandler.GetToken).Methods("GET")
	http.ListenAndServe(":7700", handlers.LoggingHandler(os.Stdout, router))
}

var renderHome = http.HandlerFunc(
	func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res, "Hello")
	},
)
