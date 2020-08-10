package main

import (
	"fmt"
	"log"
	"net/http"
)


func StartServer() {
	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	// api router here
	apiMux.HandleFunc("/", apiIndexHandelFunc)

	// router here
	mux.HandleFunc("/", indexHandelFunc)
	mux.Handle("/api/", apiMux)

	var s = &http.Server {
		Addr: ":8039",
		Handler: mux,
	}
	log.Fatal(s.ListenAndServe())
}

func apiIndexHandelFunc(w http.ResponseWriter, req *http.Request) {
	// because / match everything
	// so we need to check url here
	if req.URL.Path != "/api/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "API index welcome")
}

func indexHandelFunc(w http.ResponseWriter, req *http.Request) {
	// because / match everything
	// so we need to check url here
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Welcome to home page")
}
