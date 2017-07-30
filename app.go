package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"go-auth-sample/model"
	"encoding/json"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
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
}

func (app *App) Run(addr string) {
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

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
