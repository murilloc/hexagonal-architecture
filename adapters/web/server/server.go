package server

import (
	"github.com/gorilla/mux"
	"github.com/murilloc/go-hexagonal/adapters/web/handler"
	"github.com/murilloc/go-hexagonal/application"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (s WebServer) Serve() {
	r := mux.Router{}
	n := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(&r, n, s.Service)
	http.Handle("/", &r)

	server := &http.Server{
		ReadHeaderTimeout: 10. * time.Second,
		WriteTimeout:      10. * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.LstdFlags),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
