package handler

import (
	"fmt"
	"net/http"
)

func Route() {
	http.HandleFunc("/ping", pingHandler)
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
