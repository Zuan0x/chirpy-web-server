package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	r := chi.NewRouter()


	const filepathRoot = "./app"
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: 0,
	}

	//mux := http.NewServeMux()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))	
	r.Handle("/app/*", fsHandler)
	r.Handle("/app", fsHandler)
	r.Get("/metrics", apiCfg.handlerMetrics)
	r.Get("/healthz", apiCfg.handlerMetrics)
	r.Post("/reset", apiCfg.handlerReset)

	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())	
}



	
