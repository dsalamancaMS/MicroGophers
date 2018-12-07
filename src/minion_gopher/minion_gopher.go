package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Minion Gopher working on " + os.Getenv("HOSTNAME")))
	})
	http.ListenAndServe(":8002", nil)
}
