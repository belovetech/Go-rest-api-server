package backend

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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
	a.initializeRoutes()
}

// Routers
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", helloWorld).Methods("GET")

	a.Router.HandleFunc("/products", a.newProduct).Methods("POST")
	a.Router.HandleFunc("/products", a.allProducts).Methods("GET")
	a.Router.HandleFunc("/products/{id}", a.fetchProduct).Methods("GET")

	a.Router.HandleFunc("/orders", a.newOrder).Methods("POST")
	a.Router.HandleFunc("/orderitems", a.newOrderItem).Methods("POST")
	a.Router.HandleFunc("/orders", a.allOrders).Methods("GET")
	a.Router.HandleFunc("/orders/{id}", a.fetchOrder).Methods("GET")

}

// Handlers
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// Products
func (a *App) newProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var p product
	json.Unmarshal(reqBody, &p)

	err := p.createProduct(a.DB)
	if err != nil {
		log.Printf("newProduct error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJson(w, http.StatusCreated, p)
}

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(a.DB)

	if err != nil {
		log.Printf("allProducts error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, products)
}

func (a *App) fetchProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p product
	p.ID, _ = strconv.Atoi(id)

	err := p.getProduct(a.DB)

	if err != nil {
		log.Printf("fetchProducts error: %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, p)
}

// orders
func (a *App) newOrder(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var o order
	json.Unmarshal(reqBody, &o)

	err := o.createOrder(a.DB)
	if err != nil {
		fmt.Printf("createOrder error %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, item := range o.Items {
		var oi orderItem
		oi = item
		oi.OrderID = o.ID

		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("createOrder, createOrderItem error %s\n", err.Error())
			responseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	responseWithJson(w, http.StatusCreated, o)
}

func (a *App) newOrderItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ois []orderItem
	json.Unmarshal(reqBody, &ois)

	for _, item := range ois {
		var oi orderItem
		oi = item
		err := oi.createOrderItem(a.DB)
		if err != nil {
			fmt.Printf("createOrderItem error %s\n", err.Error())
			responseWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

	}
	responseWithJson(w, http.StatusCreated, ois)
}

func (a *App) allOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := getOrders(a.DB)
	if err != nil {
		fmt.Printf("allOrder error %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, orders)
}

func (a *App) fetchOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var o order
	o.ID, _ = strconv.Atoi(id)

	err := o.getOrder(a.DB)
	if err != nil {
		fmt.Printf("fetchOrder error %s\n", err.Error())
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, o)
}

// Helper function
func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJson(w, code, map[string]string{"Error": message})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Run server
func (a *App) Run() {
	fmt.Println("Server started and listening on localhost:", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, a.Router))
}
