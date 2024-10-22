package banco

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type paramConexao struct {
	host     string
	username string
	password string
	porta    string
	base     string
}

func ConnectToPG() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
