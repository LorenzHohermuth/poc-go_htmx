package routes

import (
	"fmt"
	"html/template"
	"log"
	"lorenz/go-htmx/internal/model"
	"net/http"
	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
			fmt.Println("Could not get all todos from db", err)
			return
	}

	const filePath = "templates/static/index.html"

	tmpl := template.Must(template.New(filePath).ParseFiles(filePath))
	
	err = tmpl.ExecuteTemplate(w, "T", "lorenz")
	fmt.Print(todos)
	//err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
			fmt.Println("Could not execute template", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sendTodos(w)
}

func markTodo(w http.ResponseWriter, r *http.Request) {

}

func creatTodo(w http.ResponseWriter, r *http.Request) {

}

func deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func SetupAndRun() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	mux.HandleFunc("/create", markTodo).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", mux))
}
