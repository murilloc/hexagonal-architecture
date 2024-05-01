package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/murilloc/go-hexagonal/application"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", getProduct(service)).Methods("GET", "OPTIONS")
	n.UseHandler(router)
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
