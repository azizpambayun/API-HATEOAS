package main

import (
	"api-hateoas/routers"
	"log"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}