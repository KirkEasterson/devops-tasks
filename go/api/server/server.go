package server

import (
	"fmt"
	"gotask/api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ServerConfig struct {
	InMemory bool
	Url      string
	Port     string
}

func Serve(s ServerConfig) {
	counter := counter.Counter(0)

	router := mux.NewRouter()
	router.HandleFunc("/metrics", counter.Count()).Methods("GET")

	port := fmt.Sprintf(":%s", s.Port)
	log.Printf("Listening at %s%s\n", s.Url, port)
	log.Print(http.ListenAndServe(port, logRequest(router)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			handler.ServeHTTP(w, r)
		})
}
