package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type paramConexao struct {
	username string
	password string
	porta    string
	base     string
	sslmode  string
}

func connectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
		panic(err)
	}

	paramConexao := paramConexao{
		username: os.Getenv("USERNAME"),
		password: os.Getenv("PASSWORD"),
		porta:    os.Getenv("PORTA"),
		base:     os.Getenv("BASE"),
		sslmode:  os.Getenv("SSLMODE"),
	}

	dbConfig := "user=" + paramConexao.username + "dbname=" + paramConexao.base + "sslmode=" + paramConexao.sslmode

	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
