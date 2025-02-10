package banco

import (
	"fmt"
	"os"
	"strconv"

	cripto "github.com/danielsilveiradevbr/helpercripto/pkg"
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
		host:     cripto.Cripto("D", os.Getenv("HOST"), os.Getenv("KEY")),
		porta:    os.Getenv("PORTA"),
		username: cripto.Cripto("D", os.Getenv("USERPOST"), os.Getenv("KEY")),
		password: cripto.Cripto("D", os.Getenv("PASSWORD"), os.Getenv("KEY")),
		base:     cripto.Cripto("D", os.Getenv("BASE"), os.Getenv("KEY")),
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
