package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"encoding/json"
	"go-auth-sample/model"
)

const (
	host        = "localhost"
	port        = "5432"
	user        = "admin"
	password    = "C0mplexPwd!1234"
	dbname      = "goauth"
	search_path = "auth"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func main() {
	//router.Handle("/", renderHome).Methods("GET")
	//router.Handle("/status", product.StatusHandler).Methods("GET")
	//router.Handle("/products", jwthandler.JwtValidator.Handler(product.ProductsHandler)).Methods("GET")
	//router.Handle("/products/{slug}/feedback", jwthandler.JwtValidator.Handler(product.AddFeedbackHandler)).Methods("POST")
	//router.Handle("/get-token", jwthandler.GetToken).Methods("GET")
	//http.ListenAndServe(":7700", handlers.LoggingHandler(os.Stdout, router))

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s search_path=%s sslmode=disable", host, port, user, password, dbname, search_path)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//app.Router = mux.NewRouter()
	a := App{mux.NewRouter(), db}

	//a.Initialize(connectionString)
	a.initializeRoutes()
	a.Run(":7700")
}

func (app *App) Initialize(connectionString string) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/register", app.register).Methods("POST")
	app.Router.Handle("/", renderHome).Methods("GET")
}

func (app *App) Run(addr string) {
	fmt.Println("Enableling Routes")
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) register(res http.ResponseWriter, req *http.Request) {
	var account model.Account
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&account); err != nil {
		respondWithError(res, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer req.Body.Close()

	fmt.Println("account : ", account)

	user, err := account.Register(app.DB, account)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
	}

	respondWithJSON(res, http.StatusOK, user)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	fmt.Println("response : ", string(response))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var renderHome = http.HandlerFunc(
	func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res, "Hello")
	},
)
