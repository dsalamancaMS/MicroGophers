package main

import (
	"net/http"
	"os"
)

func minionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Minion Gopher working on " + os.Getenv("HOSTNAME")))
}

func main() {
	http.HandleFunc("/", minionHandler)
	http.ListenAndServe(":8002", nil)
}
