package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	mux := http.NewServeMux()
	
	//Serve index.html
	filepathRoot := "."
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	//Serve healthz, return header`200 OK` if server is running
	mux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
