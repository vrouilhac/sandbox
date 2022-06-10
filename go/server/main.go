package main

import (
	"net/http"
	"log"

	"vrouilhac/server/routes"
)

func main() {
	http.HandleFunc("/", routes.HandleHome)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
