package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// host := os.Getenv("POSTGRES_HOST")
	// port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	// user := os.Getenv("POSTGRES_USER")
	// dbname := os.Getenv("POSTGRES_DB")
	// password := os.Getenv("POSTGRES_PASSWORD")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println("Connecting to database..." + dbUrl)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")

	return db, nil
}
