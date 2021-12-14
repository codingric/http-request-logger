package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func main() {

	http.HandleFunc("/", Handler)
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	body, _ := httputil.DumpRequest(r, true)
	fmt.Printf("--- %s %s ---\n", time.Now().Format(time.RFC3339), r.RemoteAddr)
	fmt.Print(string(body))
	w.Write(body)
}
