package main

import (
	"fmt"

	"github.com/danielzinhors/api_jira/src/application/controllers"
	"github.com/danielzinhors/api_jira/src/infra/banco"
)

func main() {
	PreparaDb()
	AtualizaDDS()
}

func PreparaDb() {
	db, err := banco.ConnectToDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	err = banco.ExecSqls(db)
	if err != nil {
		fmt.Println("Erro ao executar a consulta:", err)
		return
	}

}

func AtualizaDDS() {
	err := controllers.AtualizaDDS()
	if err != nil {
		panic(err)
	}
}
