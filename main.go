package main

import (
	"fmt"
	"os"

	"github.com/tlmiller/janitor/cmd"
)

//func PruneBackups(res http.ResponseWriter, req *http.Request) {
//	log.Println("got event")
//	if b, err := ioutil.ReadAll(req.Body); err == nil {
//		log.Println(string(b))
//	}
//}

func main() {
	if err := cmd.New().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "8080"
	//}

	//server := http.Server{
	//	Addr:    fmt.Sprintf(":%s", port),
	//	Handler: http.HandlerFunc(PruneBackups),
	//}

	//log.Printf("starting server on port %s", port)
	//log.Fatal(server.ListenAndServe())
}
