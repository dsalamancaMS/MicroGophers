package main

import (
	"io/ioutil"
	"net/http"
)

func bossHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8001/")
	if err != nil {
		workerResponse := "Worker is lazy"
		w.Write([]byte("Boss Gopher => " + workerResponse))
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		workerResponse := string(body)
		w.Write([]byte("Boss Gopher => " + workerResponse))
		defer resp.Body.Close()
	}
}

func main() {

	http.HandleFunc("/", bossHandler)

	http.ListenAndServe(":8000", nil)
}
