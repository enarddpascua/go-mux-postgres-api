package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

//struct
type Todo struct {
	ID   string `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

//init todo
var todos []Todo

//Get all todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

//Delete todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through Todos and find ID

	for index, todo := range todos {
		if todo.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}

//Create todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = strconv.Itoa(rand.Intn(10000000))
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

//Update todo
func updateTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")

	//init router
	r := mux.NewRouter()

	//Mock data
	todos = append(todos, Todo{
		ID:   "123",
		Todo: "Maghugas ng plato",
		Done: false,
	})
	todos = append(todos, Todo{
		ID:   "321",
		Todo: "Tumae",
		Done: false,
	})

	//route handlers/ Endpoints
	r.HandleFunc("/api/todos", getTodos).Methods("GET")
	r.HandleFunc("/api/todos/{id}", deleteTodo).Methods("DELETE")
	r.HandleFunc("/api/todo", createTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}", updateTodo).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", r))

}
