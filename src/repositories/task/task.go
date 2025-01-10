package task

import (
	taskDTO "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	model "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/assignee"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/classificacaoRelevancia"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/cliente"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/complexidade"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/creator"
	issuetype "github.com/danielsilveiradevbr/api_jira/src/repositories/issueType"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/label"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/priority"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/reporter"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/requerAnaliseTecnica"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/requerDocumentacao"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/resolution"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sku"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sprint"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/status"
	tipoalteracao "github.com/danielsilveiradevbr/api_jira/src/repositories/tipoAlteracao"
	"gorm.io/gorm"
	
)

func SalvaTask(db *gorm.DB, taskDTO *taskDTO.JsonDDS.Issue) (error) {
	if model, err := project.SalvaProject(db, &issue.Fields.Project); err != nil {
		return err
	}

	if model, err = sprint.SalvaSprint(db, issue.Fields.Sprint); err != nil {
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

	if err = requerDocumentacao.SalvarequerDocumentacao(db, &issue.Fields.RequerDocumentacao); err != nil {
		return err
	}

	if err = requerAnaliseTecnica.SalvaRequerAnaliseTecnica(db, &issue.Fields.RequerAnaliseTecnica); err != nil {
		return err
	}

	if err = cliente.SalvaCliente(db, issue.Fields.Cliente); err != nil {
		return err
	}

	if err = sku.SalvaSku(db, issue.Fields.SKU); err != nil {
		return err
	}

	if err = label.Salvalabel(db, issue.Fields.Labels); err != nil {
		return err
	}

	if err = complexidade.SalvaComplexidade(db, &issue.Fields.Complexidade); err != nil {
		return err
	}
	task := model.Newtask(taskDTO)
	result := db.First(&task, "codigo = ?", task.Codigo)
	if result.RowsAffected == 0 {
		res := db.Create(&task)
		if res.Error != nil {
			return res.Error
		}
	} else {
		db.Save(&task)
	}
	return nil
}