package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

const (
	host        = "localhost"
	port        = "5432"
	user    = "admin"
	password    = "C0mplexPwd!1234"
	dbname      = "goauth"
	search_path = "auth"
)

func main() {
	//router := mux.NewRouter()
	//router.Handle("/", renderHome).Methods("GET")
	//router.Handle("/status", product.StatusHandler).Methods("GET")
	//router.Handle("/products", jwthandler.JwtValidator.Handler(product.ProductsHandler)).Methods("GET")
	//router.Handle("/products/{slug}/feedback", jwthandler.JwtValidator.Handler(product.AddFeedbackHandler)).Methods("POST")
	//router.Handle("/get-token", jwthandler.GetToken).Methods("GET")
	//http.ListenAndServe(":7700", handlers.LoggingHandler(os.Stdout, router))

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s search_path=%s sslmode=disable", host, port, user, password, dbname, search_path)

	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully Connected.")

	a := App{}
	a.Initialize(connectionString)
	a.Run(":7700")
}

var renderHome = http.HandlerFunc(
	func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res, "Hello")
	},
)
