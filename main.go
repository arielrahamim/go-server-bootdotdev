package main

import (
	"log"
	"net/http"
)

// HealthHandler is a handler (do x when someone hits the endpoint) for the healthcheck endpoint
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
    r.Header.Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// go runtime runs main automatically
func main() {
	// create a new multiplexer (router)
	mux := http.NewServeMux()

	// serve the healthcheck endpoint
	mux.HandleFunc("/healthz", HealthzHandler)

	// serve the app directory, stripping /app/ 
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	// assign the multiplexer and address to the server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// listen and serve the server
	log.Println("listening on port", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}