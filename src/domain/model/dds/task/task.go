package taskModel

import (
	"time"

	issuesDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/issues"
	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/classificacaoRelevancia"
	clienteModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/cliente"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/complexidade"
	issueTypeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/issueType"
	labelModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/label"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/priority"
	projectModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/project"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerAnaliseTecnica"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerDocumentacao"
	resolutionModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/resolution"
	skuModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sku"
	sprintModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sprint"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/tipoAlteracao"
	userModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/user"
	helper "github.com/danielsilveiradevbr/api_jira/src/helpers"
	"gorm.io/gorm"
)

type Task struct {
	ID                               uint `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA                          string
	KEY_JIRA                         string `gorm:"unique;not null"`
	DESCRICAO                        string
	PROGRES_TOTAL                    int64
	PROGRESS                         int64
	PERC_PROGRESS                    float32
	SUMMARY                          string
	TIPO                             string
	CREATE                           time.Time
	QTDE_RETRABALHO                  float32
	DUPLICACAO_CODIGO                float32
	PADRONIZACAO_CODIGO              float32
	DOCUMENTACAO_CODIGO              float32
	LEGIBILIDADE_CODIGO              float32
	SIMPLICIDADE_CODIGO              float32
	MODULARIDADE_CODIGO              float32
	TARGET_END                       time.Time
	RESOLUTION_DATE                  time.Time
	HORAS_DEV                        float32
	ComplexidadeId                   *uint
	Complexidade                     complexidadeModel.Complexidade
	AGGREGATE_TIME_ORIGINAL_ESTIMATE float32
	AssineeId                        *uint
	Assinee                          userModel.User
	ReporterId                       *uint
	Reporter                         userModel.User
	IssueTypeId                      *uint
	IssueType                        issueTypeModel.IssueType
	ProjectId                        *uint
	Project                          projectModel.Project
	Update                           time.Time
	CUSTOMFIELD_10105                string
	PriorityId                       *uint
	Priority                         priorityModel.Priority
	SprintId                         *uint
	Sprint                           sprintModel.Sprint
	SOLUCAO                          string
	TIME_ESTIMATE                    float32
	AGGREGATE_TIME_ESTIMATE          float32
	CreatorId                        *uint
	Creator                          userModel.User
	RequerAnaliseTecnicaId           *uint
	RequerAnaliseTecnica             requerAnaliseTecnicaModel.RequerAnaliseTecnica
	CUSTOMFIELD_10434                string
	TIMESPENT                        float32
	AGGREGATE_TIMESPENT              float32
	CUSTOMFIELD_10432                string
	CUSTOMFIELD_10433                string
	WORKRATIO                        float32
	RESUMO_ALTERACAO                 string
	RequerDocumentacaoId             *uint
	RequerDocumentacao               requerDocumentacaoModel.RequerDocumentacao
	ClassificacaoRelevanciaId        *uint
	ClassificacaoRelevancia          classificacaoRelevanciaModel.ClassificacaoRelevancia
	TipoAlteracaoId                  *uint
	TipoAlteracao                    tipoAlteracaoModel.TipoAlteracao
	HORAS_TEST                       float32
	Self                             string
	ResolutionId                     *uint
	Resolution                       resolutionModel.Resolution
	ClienteId                        *uint
	Cliente                          clienteModel.Cliente
	SkuId                            *uint
	Sku                              skuModel.Sku
	LabelId                          *uint
	Label                            labelModel.Label
}

func (Task) TableName() string {
	return "tasks"
}

func (t *Task) BeforeCreate(db *gorm.DB) (err error) {
	if *t.ProjectId == 0 {
		t.ProjectId = nil
	}
	if *t.SprintId == 0 {
		t.SprintId = nil
	}
	if *t.ComplexidadeId == 0 {
		t.ComplexidadeId = nil
	}
	if *t.AssineeId == 0 {
		t.AssineeId = nil
	}
	if *t.ReporterId == 0 {
		t.ReporterId = nil
	}
	if *t.IssueTypeId == 0 {
		t.IssueTypeId = nil
	}
	if *t.PriorityId == 0 {
		t.PriorityId = nil
	}
	if *t.CreatorId == 0 {
		t.CreatorId = nil
	}
	if *t.RequerAnaliseTecnicaId == 0 {
		t.RequerAnaliseTecnicaId = nil
	}
	if *t.RequerDocumentacaoId == 0 {
		t.RequerDocumentacaoId = nil
	}
	if *t.ClassificacaoRelevanciaId == 0 {
		t.ClassificacaoRelevanciaId = nil
	}
	if *t.TipoAlteracaoId == 0 {
		t.TipoAlteracaoId = nil
	}
	if *t.ClienteId == 0 {
		t.ClienteId = nil
	}
	if *t.ResolutionId == 0 {
		t.ResolutionId = nil
	}
	if *t.SkuId == 0 {
		t.SkuId = nil
	}
	if *t.LabelId == 0 {
		t.LabelId = nil
	}
	return nil
}

func NewTask(issuesDto *issuesDto.Issues) *Task {
	return &Task{
		AGGREGATE_TIMESPENT:              issuesDto.Fields.AggregateTimeOriginalEstimate,
		AGGREGATE_TIME_ESTIMATE:          issuesDto.Fields.AggregateTimeEstimate,
		AGGREGATE_TIME_ORIGINAL_ESTIMATE: issuesDto.Fields.AggregateTimeOriginalEstimate,
		CREATE:                           helper.StrToTimeTime(issuesDto.Fields.Created),
		CUSTOMFIELD_10105:                issuesDto.Fields.Customfield10105,
		CUSTOMFIELD_10432:                issuesDto.Fields.Customfield10432,
		ID_JIRA:                          issuesDto.ID,
		KEY_JIRA:                         issuesDto.Key,
		Self:                             issuesDto.Self,
		DESCRICAO:                        issuesDto.Fields.Description,
		SUMMARY:                          issuesDto.Fields.Summary,
		QTDE_RETRABALHO:                  issuesDto.Fields.QtdeRetrabalho,
		DUPLICACAO_CODIGO:                issuesDto.Fields.DuplicacaoCodigo,
		PADRONIZACAO_CODIGO:              issuesDto.Fields.PadronizacaoCodigo,
		DOCUMENTACAO_CODIGO:              issuesDto.Fields.DocumentacaoCodigo,
		LEGIBILIDADE_CODIGO:              issuesDto.Fields.LegibilidadeCodigo,
		SIMPLICIDADE_CODIGO:              issuesDto.Fields.SimplicidadeCodigo,
		MODULARIDADE_CODIGO:              issuesDto.Fields.ModularidadeCodigo,
		TARGET_END:                       helper.StrToTimeTime(issuesDto.Fields.TargetEnd),
		RESOLUTION_DATE:                  helper.StrToTimeTime(issuesDto.Fields.ResolutionDate),
		HORAS_DEV:                        issuesDto.Fields.HorasDev,
		Update:                           helper.StrToTimeTime(issuesDto.Fields.Updated),
		SOLUCAO:                          issuesDto.Fields.Solucao,
		TIME_ESTIMATE:                    issuesDto.Fields.TimeEstimate,
		CUSTOMFIELD_10434:                issuesDto.Fields.Customfield10434,
		TIMESPENT:                        issuesDto.Fields.Timespent,
		CUSTOMFIELD_10433:                issuesDto.Fields.Customfield10433,
		WORKRATIO:                        issuesDto.Fields.Workratio,
		HORAS_TEST:                       issuesDto.Fields.HorasTest,
		RESUMO_ALTERACAO:                 issuesDto.Fields.ResumoAlteracao,
		ProjectId:                        nil,
		SprintId:                         nil,
		ComplexidadeId:                   nil,
		AssineeId:                        nil,
		ReporterId:                       nil,
		IssueTypeId:                      nil,
		PriorityId:                       nil,
		CreatorId:                        nil,
		RequerAnaliseTecnicaId:           nil,
		RequerDocumentacaoId:             nil,
		ClassificacaoRelevanciaId:        nil,
		TipoAlteracaoId:                  nil,
		ResolutionId:                     nil,
		ClienteId:                        nil,
		SkuId:                            nil,
		LabelId:                          nil,
	}
}
