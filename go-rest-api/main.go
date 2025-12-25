package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Struct to hold data
type Person struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Age       int    `db:"age" json:"age"`
}

var db *sqlx.DB

func main() {
	var err error

	dsn := "root:@tcp(localhost)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln("Database connection error:", err)
	}
	fmt.Println("Connected to MySQL!")

	// Initialize the router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/people", getPeople).Methods("GET")
	r.HandleFunc("/people/{id}", getPerson).Methods("GET")
	r.HandleFunc("/people", createPerson).Methods("POST")
	r.HandleFunc("/people/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")

	// Start the server
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// Function to get all people
func getPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person

	err := db.Select(&people, "SELECT * FROM people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(people)
}

// Function to get a single person by ID
func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	err := db.Get(&person, "SELECT * FROM people WHERE id = ?", params["id"])
	if err != nil {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(person)
}

// Function to create a new person
func createPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := db.Exec(
		"INSERT INTO people (first_name, last_name, age) VALUES (?, ?, ?)",
		person.FirstName, person.LastName, person.Age,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertID, _ := result.LastInsertId()
	person.ID = int(insertID)

	json.NewEncoder(w).Encode(person)
}

// Function to update an existing person by ID
func updatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"UPDATE people SET first_name=?, last_name=?, age=? WHERE id=?",
		person.FirstName, person.LastName, person.Age, params["id"],
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(person)
}

// Function to delete a person by ID
func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := db.Exec("DELETE FROM people WHERE id = ?", params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Person deleted",
	})
}
