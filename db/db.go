//package db
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"github.com/joho/godotenv"
//	_ "github.com/lib/pq" // Драйвер для PostgreSQL
//)
//
//var DB *sql.DB // Глобальная переменная для соединения с БД
//
//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Println("Error loading .env file")
//	}
//
//	dbHost := os.Getenv("DB_HOST")
//	dbPort := os.Getenv("DB_PORT")
//	dbUser := os.Getenv("DB_USER")
//	dbPassword := os.Getenv("DB_PASSWORD") // Может быть пустым
//	dbName := os.Getenv("DB_NAME")
//
//	// Проверяем, что все переменные окружения установлены
//	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
//		log.Fatal("Missing required environment variables. Please check your .env file.")
//	}
//
//	var connStr string
//	if dbPassword == "" {
//		connStr = fmt.Sprintf(
//			"host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName,
//		)
//	} else {
//		connStr = fmt.Sprintf(
//			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//			dbHost, dbPort, dbUser, dbPassword, dbName,
//		)
//	}
//
//	// Добавляем задержку перед подключением к базе данных
//	time.Sleep(5 * time.Second)
//
//	DB, err = sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	errPing := DB.Ping()
//	if errPing != nil {
//		log.Fatal(errPing)
//	}
//
//	// Создание таблицы tasks, если она не существует
//	_, err = DB.Exec(`
//       CREATE TABLE IF NOT EXISTS tasks (
//           id SERIAL PRIMARY KEY,
//           title VARCHAR(255) NOT NULL,
//           done BOOLEAN DEFAULT FALSE,
//           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//       )
//   `)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	log.Println("Database connected")
//}

//package db
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//	"os"
//	"time"
//
//	"github.com/joho/godotenv"
//	_ "github.com/lib/pq" // Драйвер для PostgreSQL
//)
//
//var DB *sql.DB // Глобальная переменная для соединения с БД
//
//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Println("Error loading .env file")
//	}
//
//	// Получаем параметры БД
//	dbHost := os.Getenv("DB_HOST")
//	dbPort := os.Getenv("DB_PORT")
//	dbUser := os.Getenv("DB_USER")
//	dbPassword := os.Getenv("DB_PASSWORD") // Может быть пустым
//	dbName := os.Getenv("DB_NAME")
//
//	var connStr string
//	if dbPassword == "" {
//		connStr = fmt.Sprintf(
//			"host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName,
//		)
//	} else {
//		connStr = fmt.Sprintf(
//			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//			dbHost, dbPort, dbUser, dbPassword, dbName,
//		)
//	}
//
//	// Добавляем задержку перед подключением к базе данных
//	time.Sleep(5 * time.Second)
//
//	DB, err = sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	errPing := DB.Ping()
//	if errPing != nil {
//		log.Fatal(errPing)
//	}
//
//	log.Println("Database connected")
//}

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	// Получаем параметры БД
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Формируем строку подключения
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Подключаемся к БД с повторными попытками
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Проверяем подключение с задержкой
	for i := 1; i <= 5; i++ {
		err = DB.Ping()
		if err == nil {
			break
		}
		log.Printf("Попытка %d: ошибка проверки подключения: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Не удалось подключиться к БД: %v", err)
	}

	// Создаем таблицу с логированием ошибок
	log.Println("Создаем таблицу tasks...")
	_, err = DB.Exec(`
        CREATE TABLE IF NOT EXISTS public.tasks (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            done BOOLEAN DEFAULT FALSE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		log.Fatalf("Ошибка создания таблицы: %v", err)
	}
	log.Println("Таблица tasks создана успешно")
}
