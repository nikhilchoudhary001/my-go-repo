package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/", handler1)
	http.HandleFunc("/api/users", handler2)
	http.HandleFunc("/api/login", handler3)
	http.HandleFunc("/api/student", handler4)

	log.Fatal(http.ListenAndServe(":8081", nil))

	fmt.Println("After server")

}

func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside server 1")

	fmt.Fprintf(w, "Helloe from Server 1")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside server 2")

	fmt.Fprintf(w, "Helloe from Server 2")
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside server 3")

	fmt.Fprintf(w, "Helloe from Server 3")
}

func handler4(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside server 4")

	fmt.Fprintf(w, "Helloe from Server 4")
}
