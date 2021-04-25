package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("listen on port: 8080")
	fmt.Println("accessing \"http://localhost:8080/\" will response hit hostname")

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":8080", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// get hostname
	log.Println(r.Method)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("get host name failed:", err.Error())
	}

	responseStr := fmt.Sprintf("\n******************************\n")
	responseStr += fmt.Sprintf("You've hit %s:8080\n", hostname)
	responseStr += fmt.Sprintf("******************************\n\n")

	fmt.Fprintf(w, responseStr)
}
