package banco

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type paramConexao struct {
	host     string
	username string
	password string
	porta    string
	base     string
}

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	paramConexao := paramConexao{
		host:     os.Getenv("HOST"),
		porta:    os.Getenv("PORTA"),
		username: os.Getenv("USERPOST"),
		password: os.Getenv("PASSWORD"),
		base:     os.Getenv("BASE"),
	}
	porta, err := strconv.Atoi(paramConexao.porta)
	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		paramConexao.host, porta, paramConexao.username, paramConexao.password, paramConexao.base)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ExecSqls(db *sql.DB) error {
	defer db.Close()
	var sqls = ExexActionsDb()

	for _, v := range sqls {
		stmt, err := db.Prepare(v)
		defer stmt.Close()
		if err != nil {
			return err
		}
		_, err = stmt.Exec()
		// fmt.Println(v)
		if err != nil {
			return err
		}
	}
	return nil
}
