package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB     *sql.DB
	Port   string
	Router *mux.Router
}

func (a *App) Initialize() {
	db, err := sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	a.DB = db
	a.Router = mux.NewRouter()
	a.initializeRouter()
}

func (a *App) initializeRouter() {
	a.Router.HandleFunc("/", helloWorld).Methods("GET")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func (a *App) Run() {
	fmt.Println("Server started and listening on localhost:", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
