package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8002/")
		if err != nil {
			workerResponse := "Minon Gopher slain"
			w.Write([]byte("Worker Gopher => " + workerResponse))
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			workerResponse := string(body)
			w.Write([]byte("Worker Gopher => " + workerResponse))
			defer resp.Body.Close()
		}
	})

	http.ListenAndServe(":8001", nil)
}
