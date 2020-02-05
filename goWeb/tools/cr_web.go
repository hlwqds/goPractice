package main

import (
	"fmt"
	"net/http"
)

func HandleCrWebReuqest(w http.ResponseWriter, r *http.Request){
	
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8777",
	}

	http.HandleFunc("/cr_web", HandleCrWebReuqest)

	server.ListenAndServe()
}