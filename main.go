package main

import (
	"fmt"

	_ "github.com/danielzinhors/api_jira/src/infra"
)

func main() {
	PreparaDb()
}

func PreparaDb() {
	_, err := *banco.connectToDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	if err != nil {
		fmt.Println("Erro ao executar a consulta:", err)
		return
	}

}
