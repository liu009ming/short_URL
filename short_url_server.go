package main

import "fmt"
import "log"
import "net/http"

func main() {
	http.HandleFunc("/",hanler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func hanler(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
}