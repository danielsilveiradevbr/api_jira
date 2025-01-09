package ddsservice

import (
	jsonDDS "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	b "github.com/danielsilveiradevbr/api_jira/src/infra/banco"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/assignee"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/classificacaoRelevancia"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/creator"
	issuetype "github.com/danielsilveiradevbr/api_jira/src/repositories/issueType"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/priority"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/reporter"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/resolution"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sprint"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/status"
	tipoalteracao "github.com/danielsilveiradevbr/api_jira/src/repositories/tipoAlteracao"
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

		if err = creator.SalvaCreator(db, &issue.Fields.Creator); err != nil {
			return err
		}

		if err = priority.SalvaPriority(db, &issue.Fields.Priority); err != nil {
			return err
		}

		if err = status.SalvaStatus(db, &issue.Fields.Status); err != nil {
			return err
		}

		if err = resolution.SalvaResolution(db, &issue.Fields.Resolution); err != nil {
			return err
		}

		if err = tipoalteracao.SalvatipoAlteracao(db, &issue.Fields.TipoAlteracao); err != nil {
			return err
		}

		if err = classificacaoRelevancia.SalvaclassificacaoRelevancia(db, &issue.Fields.ClassificacaoRelevancia); err != nil {
			return err
		}
	}

	return nil
}
