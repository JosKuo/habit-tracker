package main

import (
	//"encoding/json"
	"log"
	"net/http"
)

func getHabits(w http.ResponseWriter, r *http.Request) {
	log.Println("im in service")
}

func main() {
	http.HandleFunc("/habits", getHabits)
	log.Println("Server is up and running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
