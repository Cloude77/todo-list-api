// Новый пакет для работы с базой данных

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" //Драйвер для PostgreSQL
)

//
//var DB *sql.DB //Глобальная переменная для соединения с БД
//
//// ConnectToDB устанавливает соединение с PostgreSQL
//
//func ConnectDB() {
//	connStr := "host=localhost port=5432 user=postgres dbname=todo_db sslmode=disable"
//	var err error
//	DB, err = sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = DB.Ping()
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Successfully connected!")

// }
var DB *sql.DB //Глобальная переменная для соединения с БД

// init загружает переменные окружения и устанавливает соединение с БД
func init() {
	// Загрузка переменных из окружения .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Получение параметров из переменный окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := ""
	dbName := os.Getenv("DB_NAME")

	//Формирование строки подключения
	var connStr string
	if dbPassword == "" {
		// Если пароль не указан, формируем строку без него
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName,
		)

	} else {
		//Если пароль указан, добавляем его в строку подключения
		connStr = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName,
		)
	}
	// Установка с БД
	var errDB error
	DB, errDB = sql.Open("postgres", connStr)
	if errDB != nil {
		log.Fatal(errDB)
	}
	errPing := DB.Ping()
	if errPing != nil {
		log.Fatal(errPing)
	}
	log.Println("Database connected")
}
