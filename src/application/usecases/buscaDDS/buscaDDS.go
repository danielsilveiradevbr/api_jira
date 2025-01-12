package buscaDDS

import (
	dds "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/jsonDDS"
	"github.com/danielsilveiradevbr/api_jira/src/infra/jira"
)

func BuscaDDS(sprintFiltro string) (*dds.JsonDDS, error) {
	jsonDDS, err := jira.BuscaDDS(sprintFiltro)
	if err != nil {
		return nil, err
	}
	return jsonDDS, nil
}
