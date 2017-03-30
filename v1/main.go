package main

import (
	"log"
	"net/http"

	"github.com/shiva0705/goApi/v1/api/coreapi"
	"github.com/shiva0705/goApi/v1/sqs"
)

func main() {

	sqs.NewSession()

	router := coreapi.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
