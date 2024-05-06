package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/murilloc/go-hexagonal/adapters/dto"
	"github.com/murilloc/go-hexagonal/application"
	"github.com/urfave/negroni"
	"net/http"
)

func MakeProductHandlers(router *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	router.Handle("/product/{id}", getProduct(service)).Methods("GET", "OPTIONS")
	router.Handle("/product", createProduct(service)).Methods("POST", "OPTIONS")
	router.Handle("/product/{id}/enable", enableProduct(service)).Methods("GET", "OPTIONS")
	router.Handle("/product/{id}/disable", disableProduct(service)).Methods("GET", "OPTIONS")

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

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			jsonErr, marshalErr := jsonError(err.Error())
			if marshalErr != nil {
				// handle error
				fmt.Println(marshalErr)
			} else {
				_, err := w.Write(jsonErr)
				if err != nil {
					return
				}
				return
			}
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		result, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
