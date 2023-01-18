package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mayank979/go/storage"
	"github.com/gorilla/mux"
)

type APIServer struct {
	port string
	db   storage.PG
}

func RunAPIServer(port string, db storage.PG) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.Handle("/account", makeHTTPHandlerFunc(s.createAccount)).Methods("POST")
	router.Handle("/account", makeHTTPHandlerFunc(s.getAllAccounts)).Methods("GET")
	router.Handle("/account/{id}", makeHTTPHandlerFunc(s.getAccountById)).Methods("GET")

	fmt.Println("Server running on port", s.port)

	http.ListenAndServe(s.port, router)

}

func (s *APIServer) createAccount(w http.ResponseWriter, r *http.Request) error {

	account := storage.NewAccount("Mayank", "Yadav")

	return WriteJSON(w, 200, account)
}

func (s *APIServer) getAllAccounts(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, 200, map[string]string{"result": "success"})
}

func (s *APIServer) getAccountById(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, 200, map[string]string{"result": "success"})
}

type APIError struct {
	Error string `json:"error"`
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHTTPHandlerFunc(f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, &APIError{Error: err.Error()})
		}
	}
}
