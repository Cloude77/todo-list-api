package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo-list/handlers"
)

var dbReady = make(chan struct{})

func main() {
	//Подключение к БД

	r := mux.NewRouter()

	//Регистрация обработчиков
	r.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")        //POST /tasks
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")           //Get /tasks
	r.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")   //Get /tasks?id=1
	r.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")    //PUT /tasks?=1
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE") //DELETE /tasks?=1

	//Запуск сервера.
	log.Println("Starting server on  :8080")
	if err := http.ListenAndServe(":8080", r); err != nil { //r - передаем роутер
		log.Fatal(err)
	}
}
