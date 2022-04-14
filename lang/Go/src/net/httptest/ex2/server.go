package main

import (
	"net/http"
)

func handleSlothfulMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Stay slothful!"}`))
}

func appRouter() http.Handler {
	rt := http.NewServeMux()
	rt.HandleFunc("/sloth", handleSlothfulMessage)
	return rt
}

func main() { http.ListenAndServe(":1123", appRouter()) }
