package ddsservice

import (
	"fmt"

	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sprint"
	"github.com/joho/godotenv"
)

func SalvaDDS(DDSJson *jsonDDS.JsonDDS) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	db, err := b.ConnectToPG()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		panic(err)
	}
	for _, issue := range DDSJson.Issues {
		err := project.SalvaProject(db, &issue.Fields.Project)
		if err != nil {
			return err
		}
		err = sprint.SalvaSprint(db, issue.Fields.Sprint)
		if err != nil {
			return err
		}
	}

	return nil
}
