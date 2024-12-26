package ddsservice

import (
	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"gorm.io/gorm"
)

func SalvaDDS(prCon *gorm.DB, DDSJson *jsonDDS.JsonDDS) error {
	for _, issue := range DDSJson.Issues {
		err := project.SalvaProject(prCon, &issue.Fields.Project)
		if err != nil {
			return err
		}
	}
	return nil
}
