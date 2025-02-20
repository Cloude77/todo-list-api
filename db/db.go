// Новый пакет для работы с базой данных

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //Драйвер для PostgreSQL
)

var DB *sql.DB //Глобальная переменная для соединения с БД

// ConnectToDB устанавливает соединение с PostgreSQL

func ConnectDB() {
	connStr := "host=localhost port=5432 user=postgres dbname=todo_db sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

}
