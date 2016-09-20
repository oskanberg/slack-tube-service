package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

var lastStatusCheck time.Time

const listenPort string = "1123"

func init() {
	var err error
	lastStatusCheck, err = time.Parse(time.RFC3339, "1970-01-01T00:00:00+00:00")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	loadAuthorisedTokensFromFile(authorisedTokenFileLocation)
	router := newRouter()
	fmt.Println("Ready, listening on port", listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, cors.Default().Handler(router)))
}
