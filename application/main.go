package main

import (
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/sample", sampleHandler)
	http.HandleFunc("/health", func(rw http.ResponseWriter, r *http.Request) { rw.WriteHeader(200) })

	if err := http.ListenAndServe(":9000", nil); err != nil {
		panic(err)
	}
}

func sampleHandler(rw http.ResponseWriter, r *http.Request) {
	if rand.Float32() > 0.1 {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}
