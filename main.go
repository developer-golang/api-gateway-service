package main

import (
	"log"
	"net/http"

	r "api-gateway-service/router"
)

func main() {
	gateway := r.NewRouter()

	log.Println("API Gateway running on :8087")

	err := http.ListenAndServe(":8087", gateway)
	if err != nil {
		log.Fatal(err)
	}
}
