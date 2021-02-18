package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type n1class struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Roll int    `json:"Roll"`
	Age  int    `json:"Age"`
}

var n2 []n1class

func getclass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "json file: ")
	json.NewEncoder(w).Encode(n2)
}
func updatestudent(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	var x n1class
	json.NewDecoder(r.Body).Decode(&x)
	for index, i := range n2 {
		if i.Id == key {
			i.Id = x.Id
			i.Name = x.Name
			i.Roll = x.Roll
			i.Age = x.Age
			n2 := append(n2[:index], n2[index+1:]...)
			n2 = append(n2, i)
			json.NewEncoder(w).Encode(n2)
		}
	}
}
func updateparticular(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	var x n1class
	json.NewDecoder(r.Body).Decode(&x)
	for index, i := range n2 {
		if i.Name == name {
			i.Name = x.Name
			n2 := append(n2[:index], n2[index+1:]...)
			n2 = append(n2, i)
			json.NewEncoder(w).Encode(n2)
		}
	}
}
func deletestudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for index, i := range n2 {
		if i.Id == id {
			n2 = append(n2[:index], n2[index+1:]...)
		}
	}
}
func createStudent(w http.ResponseWriter, r *http.Request) {

	var x n1class
	json.NewDecoder(r.Body).Decode(&x)
	n2 = append(n2, x)
	json.NewEncoder(w).Encode(n2)
}
func webpage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to our company\n")
	fmt.Fprintf(w, "Hello World")
}
func return1article(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["name"]
	for _, i := range n2 {
		if i.Name == key {
			json.NewEncoder(w).Encode(i)
		}
	}
}
func handleRequest() {
	myrouter := mux.NewRouter().StrictSlash(true)
	myrouter.HandleFunc("/", webpage)
	myrouter.HandleFunc("/class", getclass).Methods("GET")
	myrouter.HandleFunc("/class", createStudent).Methods("POST")
	myrouter.HandleFunc("/class/{id}", updatestudent).Methods("PUT")
	myrouter.HandleFunc("/class/{name}", updateparticular).Methods("PATCH")
	myrouter.HandleFunc("/class/{id}", deletestudent).Methods("DELETE")
	myrouter.HandleFunc("/class/{name}", return1article)
	http.ListenAndServe(":7070", myrouter)

}
func main() {

	n2 = []n1class{{Id: "1", Name: "Rahul", Roll: 1, Age: 24},
		{Id: "2", Name: "Deepak", Roll: 2, Age: 24}}
	handleRequest()
}
