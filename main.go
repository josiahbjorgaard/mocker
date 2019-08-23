package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	//"strconv"
)

//Note is a structure to hold notes
type Note struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

//Init specification as a slice
var notes []Note

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, elem := range notes {
		if elem.ID == params["id"] {
			json.NewEncoder(w).Encode(elem)
			return
		}
	}
	json.NewEncoder(w).Encode(notes)
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newNote Note
	//	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		fmt.Println("Error")
	}
	newNote.ID = strconv.Itoa(len(notes) + 1)
	//	newNote.Content = params["content"]
	//	newNote.Name = params["name"]
	notes = append(notes, newNote)
	json.NewEncoder(w).Encode(newNote)
}

func updateNote(w http.ResponseWriter, r *http.Request) {

}

func deleteNote(w http.ResponseWriter, r *http.Request) {

}

func main() {

	portPtr := flag.Int("port", 5000, "an integer specifying the receiving port")
	flag.Parse()

	notes = append(notes, Note{ID: "1", Name: "NULL", Content: "NULL"})

	router := mux.NewRouter()
	router.HandleFunc("/note", getNotes).Methods("GET")
	router.HandleFunc("/note/{id}", getNote).Methods("GET")
	router.HandleFunc("/note", createNote).Methods("POST")
	router.HandleFunc("/note/{id}", updateNote).Methods("POST")
	router.HandleFunc("/note/{id}", deleteNote).Methods("DELETE")
	portString := ":" + strconv.Itoa(*portPtr)
	fmt.Println(portString)
	log.Fatal(http.ListenAndServe(portString, router))
}
