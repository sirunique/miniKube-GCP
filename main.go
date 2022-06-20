package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to minikube gcp")
	fmt.Println("Endpoint Hit: homepage")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	who := "Muritadoh Sodeeq"

	fmt.Fprintf(w, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func whoAmI(w http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Muritadoh Sodeeq", Title: "Kubernetes Engoneer", State: "Lagos"},
	}

	json.NewEncoder(w).Encode(who)
	fmt.Println("Endpoint Hit : ", who)
}

type whoami struct {
	Name  string
	Title string
	State string
}
