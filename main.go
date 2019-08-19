package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func PruneBackups(res http.ResponseWriter, req *http.Request) {
	log.Println("got event")
	if b, err := ioutil.ReadAll(req.Body); err == nil {
		log.Println(string(b))
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: http.HandlerFunc(PruneBackups),
	}

	log.Printf("starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}
