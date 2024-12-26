package main

import (
	"fmt"

	"github.com/danielsilveiradevbr/api_jira/src/application/controllers"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	service "github.com/danielsilveiradevbr/api_jira/src/service/ddsService"
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
	jsonJira, err := controllers.AtualizaDDS()
	if err != nil {
		panic(err)
	}
	service.SalvaDDS(prCon, jsonJira)
}
