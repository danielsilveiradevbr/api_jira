package ddsservice

import (
	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/assignee"
	issuetype "github.com/danielsilveiradevbr/api_jira/src/repositories/issueType"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/priority"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/reporter"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sprint"
)

func SalvaDDS(DDSJson *jsonDDS.JsonDDS) error {
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

		if err = reporter.SalvaReporter(db, &issue.Fields.Reporter); err != nil {
			return err
		}

		if err = priority.SalvaPriority(db, &issue.Fields.Priority); err != nil {
			return err
		}
	}

	return nil
}
