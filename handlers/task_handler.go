package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo-list/db"
	"todo-list/models"
)

//var tasks = make(map[int]*models.Task) // Хранилище задач
//var idCounter = 1                      // Счетчик ID

// CreateTask создает новую задачу
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.CreatedAt = time.Now()

	// Установка ID и времени создания
	//newTask.ID = idCounter
	//newTask.CreatedAt = time.Now()
	//tasks[idCounter] = &newTask
	//idCounter++

	query := `INSERT INTO tasks (title, done, created_at) VALUES ($1, $2, $3) RETURNING id`
	err = db.DB.QueryRow(query, newTask.Title, newTask.Done, newTask.CreatedAt).Scan(&newTask.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	// Отправка ответа клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

// GetTasks возвращает список всех задач
func GetTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, done, created_at FROM tasks")
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt) //Извлекаем значения из строки и записываем их в структуру
		if err != nil {
			log.Println(err)
			http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, &task)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		http.Error(w, "Failed to get tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID возвращает задачу по ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из параметров URL
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "ID must be a positive integer", http.StatusBadRequest)
		return
	}

	var task models.Task
	query := `SELECT id, title, done, created_at FROM tasks WHERE id = $1` //Создаём SQL-запрос для получения задачи по ID
	err = db.DB.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Done, &task.CreatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusInternalServerError)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}
	// Отправляем задачу в ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

}

// UpdateTask обновляет задачу
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из параметров URL
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "ID must be a positive integer", http.StatusBadRequest)
		return
	}

	// Декодируем новые данные задачи
	var updatedTask models.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := `UPDATE tasks SET title=$1, done=$2 WHERE id=$3 RETURNING id`         //Создаём SQL-запрос для обновления задачи.
	err = db.DB.QueryRow(query, updatedTask.Title, updatedTask.Done, id).Scan(&id) //Выполняем запрос и проверяем, успешно ли он выполнен
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get task", http.StatusInternalServerError)
		return
	}

	// Отправляем обновленную задачу в ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask удаляет задачу
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из параметров URL
	params := mux.Vars(r)
	idParam := params["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "ID must be a positive integer", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM tasks WHERE id = $1` //Создаём SQL-запрос для удаления задачи
	res, err := db.DB.Exec(query, id)          //Выполняем запрос без возврата данных.
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	count, err := res.RowsAffected()
	if err != nil || count == 0 { //Проверяем, сколько строк было удалено. Если 0, значит задача не найдена.
		http.Error(w, "Task not found", http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusNoContent)
}
