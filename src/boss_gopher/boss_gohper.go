package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get("http://localhost:8001/")
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		worker_response := string(body)
		w.Write([]byte("Boss Gopher => " + worker_response))
	})
	http.ListenAndServe(":8000", nil)
}
