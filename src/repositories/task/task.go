package taskRep

import (
	issuesDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/issues"
	taskModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/task"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/assignee"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/auxiliares/progressTask"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/auxiliares/statusTask"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/classificacaoRelevancia"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/cliente"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/complexidade"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/creator"
	issuetype "github.com/danielsilveiradevbr/api_jira/src/repositories/issueType"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/label"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/priority"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/progress"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/project"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/reporter"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/requerAnaliseTecnica"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/requerDocumentacao"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/resolution"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sku"
	"github.com/danielsilveiradevbr/api_jira/src/repositories/sprint"
	status "github.com/danielsilveiradevbr/api_jira/src/repositories/status"
	tipoalteracao "github.com/danielsilveiradevbr/api_jira/src/repositories/tipoAlteracao"
	"gorm.io/gorm"
)

func SalvaTask(db *gorm.DB, taskDTO *issuesDto.Issues) (*taskModel.Task, error) {
	project, err := project.SalvaProject(db, &taskDTO.Fields.Project)
	if err != nil {
		return nil, err
	}

	sprint, err := sprint.SalvaSprint(db, taskDTO.Fields.Sprint)
	if err != nil {
		return nil, err
	}

	issueType, err := issuetype.SalvaIssueType(db, &taskDTO.Fields.Issuetype)
	if err != nil {
		return nil, err
	}

	assignee, err := assignee.SalvaAssigned(db, &taskDTO.Fields.Assinee)
	if err != nil {
		return nil, err
	}

	reporter, err := reporter.SalvaReporter(db, &taskDTO.Fields.Reporter)
	if err != nil {
		return nil, err
	}

	creator, err := creator.SalvaCreator(db, &taskDTO.Fields.Creator)
	if err != nil {
		return nil, err
	}

	priority, err := priority.SalvaPriority(db, &taskDTO.Fields.Priority)
	if err != nil {
		return nil, err
	}

	resolution, err := resolution.SalvaResolution(db, &taskDTO.Fields.Resolution)
	if err != nil {
		return nil, err
	}

	tipoAlteracao, err := tipoalteracao.SalvaTipoAlteracao(db, &taskDTO.Fields.TipoAlteracao)
	if err != nil {
		return nil, err
	}

	classificacaoRelevancia, err := classificacaoRelevancia.SalvaClassificacaoRelevancia(db, &taskDTO.Fields.ClassificacaoRelevancia)
	if err != nil {
		return nil, err
	}

	requerDocumentacao, err := requerDocumentacao.SalvaRequerDocumentacao(db, &taskDTO.Fields.RequerDocumentacao)
	if err != nil {
		return nil, err
	}

	requerAnaliseTecnica, err := requerAnaliseTecnica.SalvaRequerAnaliseTecnica(db, &taskDTO.Fields.RequerAnaliseTecnica)
	if err != nil {
		return nil, err
	}

	cliente, err := cliente.SalvaCliente(db, taskDTO.Fields.Cliente)
	if err != nil {
		return nil, err
	}

	sku, err := sku.SalvaSku(db, taskDTO.Fields.SKU)
	if err != nil {
		return nil, err
	}

	label, err := label.Salvalabel(db, taskDTO.Fields.Labels)
	if err != nil {
		return nil, err
	}

	complexidade, err := complexidade.SalvaComplexidade(db, &taskDTO.Fields.Complexidade)
	if err != nil {
		return nil, err
	}
	task := taskModel.NewTask(taskDTO)
	task.ProjectId = &project.ID
	task.SprintId = &sprint.ID
	task.TIPO = project.KEY_JIRA
	task.ComplexidadeId = &complexidade.ID
	task.AssineeId = &assignee.ID
	task.ReporterId = &reporter.ID
	task.IssueTypeId = &issueType.ID
	task.PriorityId = &priority.ID
	task.CreatorId = &creator.ID
	task.RequerAnaliseTecnicaId = &requerAnaliseTecnica.Id
	task.RequerDocumentacaoId = &requerDocumentacao.Id
	task.ClassificacaoRelevanciaId = &classificacaoRelevancia.Id
	task.TipoAlteracaoId = &tipoAlteracao.Id
	task.ResolutionId = &resolution.ID
	task.ClienteId = &cliente.ID
	task.SkuId = &sku.ID
	task.LabelId = &label.ID
	result := db.First(&task, "key_jira = ?", task.KEY_JIRA)
	if result.RowsAffected == 0 {
		res := db.Create(&task)
		if res.Error != nil {
			return nil, res.Error
		}
	} else {
		db.Save(&task)
	}

	status, err := status.SalvaStatus(db, &taskDTO.Fields.Status)
	if err != nil {
		return nil, err
	}

	err = statusTask.SalvaStatusTask(db, status, task)
	if err != nil {
		return nil, err
	}

	progress, err := progress.SalvaProgress(db, &taskDTO.Fields.Progress)
	if err != nil {
		return nil, err
	}

	err = progressTask.SalvaProgressTask(db, progress, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
