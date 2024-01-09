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

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}
	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
