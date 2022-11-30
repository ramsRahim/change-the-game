package main

import (
	"net/http"

	"github.com/change-the-game/Http"
)

func main() {

	http.HandleFunc("/hello", Http.Hello)
	http.HandleFunc("/headers", Http.Headers)

	http.ListenAndServe(":8090", nil)
}
