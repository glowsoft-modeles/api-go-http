package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/sante", func(w http.ResponseWriter, _ *http.Request) { w.Write([]byte("ok")) })
	http.ListenAndServe(":"+port, nil)
}
