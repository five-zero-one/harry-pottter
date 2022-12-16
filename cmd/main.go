package main

import (
	"fmt"
	"harry-potter/service"
	"log"
	"net/http"

	"github.com/rs/cors"
)

const port = 8080
const web = "http://localhost:8000"

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	o := cors.Options{
		AllowedOrigins: []string{web},
		AllowedMethods: []string{http.MethodGet},
	}

	h := service.New()

	srv := http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		Handler:  cors.New(o).Handler(h),
		ErrorLog: log.Default(),
	}

	srv.ErrorLog.Printf("listening to http://localhost:%d", port)
	return srv.ListenAndServe()
}
