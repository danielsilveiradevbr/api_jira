package ddsservice

import (
	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/task"
)

func SalvaDDS(DDSJson *jsonDDS.JsonDDS) error {
	db, err := b.ConnectToPG()
	if err != nil {
		return err
	}
	for _, issue := range DDSJson.Issues {
		if err := task.SalvaTask(db, &issue); err != nil {
			return err
		}
	}

	return nil
}
