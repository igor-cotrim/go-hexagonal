package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/igor-cotrim/go-hexagonal/adapters/web/handlers"
	"github.com/igor-cotrim/go-hexagonal/application"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer(service application.ProductServiceInterface) *WebServer {
	return &WebServer{Service: service}
}

func (w WebServer) Serve() {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handlers.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
