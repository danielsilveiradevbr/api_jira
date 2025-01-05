package ddsservice

import (
	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/assignee"
	issuetype "github.com/danielsilveiradevbr/api_jira/src/repositories/issueType"
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
		return err
	}
	for _, issue := range DDSJson.Issues {

		if err := project.SalvaProject(db, &issue.Fields.Project); err != nil {
			return err
		}

		if err = sprint.SalvaSprint(db, issue.Fields.Sprint); err != nil {
			return err
		}

		if err = issuetype.SalvaIssueType(db, &issue.Fields.Issuetype); err != nil {
			return err
		}

		if err = assignee.SalvaAssigned(db, &issue.Fields.Assinee); err != nil {
			return err
		}
	}

	return nil
}
