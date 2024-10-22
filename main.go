package main

import (
	"fmt"

	"github.com/danielsilveiradevbr/api_jira/src/application/controllers"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"gorm.io/gorm"
)

func main() {
	db, err := b.ConnectToPG()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		panic(err)
	}
	AtualizaDDS(db)
}

func AtualizaDDS(prCon *gorm.DB) {
	jsondds, err := controllers.AtualizaDDS()
	if err != nil {
		panic(err)
	}

}
