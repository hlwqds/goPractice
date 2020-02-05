package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type Post struct {
	User string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request){
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func wrtieHeaderExample(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door.")
}

func headerExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://127.0.0.1:8080/write")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User: "Huang Lin",
		Threads: []string{"first", "second", "third"},
	}

	json, _ := json.Marshal(post)
	w.Write(json)
}

func main(){
	server := http.Server{
		Addr:"127.0.0.1:8080",
	}	

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", wrtieHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}