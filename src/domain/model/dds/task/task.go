package taskModel

import (
	"time"

	issuesDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/issues"
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
	"github.com/danielsilveiradevbr/api_jira/src/utils"
)

type Task struct {
	ID                               int64 `gorm:"primaryKey;autoIncrement:true"`
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
	COMPLEXIDADEId                   int64
	COMPLEXIDADE                     complexidadeModel.Complexidade
	AGGREGATE_TIME_ORIGINAL_ESTIMATE float32
	AssineeId                        int64
	Assinee                          userModel.User
	ReporterId                       int64
	Reporter                         userModel.User
	IssueTypeId                      int64
	IssueType                        issueTypeModel.IssueType
	ProjectId                        int64
	Project                          projectModel.Project
	Update                           time.Time
	CUSTOMFIELD_10105                string
	PriorityId                       int64
	Priority                         priorityModel.Priority
	SprintId                         int64
	Sprint                           sprintModel.Sprint
	SOLUCAO                          string
	TIME_ESTIMATE                    float32
	AGGREGATE_TIME_ESTIMATE          float32
	CreatorId                        int64
	Creator                          userModel.User
	RequerAnaliseTecnicaId           int64
	RequerAnaliseTecnica             requerAnaliseTecnicaModel.RequerAnaliseTecnica
	CUSTOMFIELD_10434                string
	TIMESPENT                        float32
	AGGREGATE_TIMESPENT              float32
	CUSTOMFIELD_10432                string
	CUSTOMFIELD_10433                string
	WORKRATIO                        float32
	RESUMO_ALTERACAO                 string
	RequerDocumentacaoId             int64
	RequerDocumentacao               requerDocumentacaoModel.RequerDocumentacao
	ClassificacaoRelevanciaId        int64
	ClassificacaoRelevancia          classificacaoRelevanciaModel.ClassificacaoRelevancia
	TipoAlteracaoId                  int64
	TipoAlteracao                    tipoAlteracaoModel.TipoAlteracao
	HORAS_TEST                       float32
	Self                             string
	ResolutionId                     int64
	Resolution                       resolutionModel.Resolution
	ClienteId                        int64
	Cliente                          clienteModel.Cliente
	SkuId                            int64
	Sku                              skuModel.Sku
	LabelId                          int64
	Label                            labelModel.Label
}

func (Task) TableName() string {
	return "tasks"
}

func NewTask(issuesDto *issuesDto.Issues) *Task {
	return &Task{
		AGGREGATE_TIMESPENT:              issuesDto.Fields.AggregateTimeOriginalEstimate,
		AGGREGATE_TIME_ESTIMATE:          issuesDto.Fields.AggregateTimeEstimate,
		AGGREGATE_TIME_ORIGINAL_ESTIMATE: issuesDto.Fields.AggregateTimeOriginalEstimate,
		CREATE:                           utils.StrToTimeTime(issuesDto.Fields.Created),
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
		TARGET_END:                       utils.StrToTimeTime(issuesDto.Fields.TargetEnd),
		RESOLUTION_DATE:                  utils.StrToTimeTime(issuesDto.Fields.ResolutionDate),
		HORAS_DEV:                        issuesDto.Fields.HorasDev,
		Update:                           utils.StrToTimeTime(issuesDto.Fields.Updated),
		SOLUCAO:                          issuesDto.Fields.Solucao,
		TIME_ESTIMATE:                    issuesDto.Fields.TimeEstimate,
		CUSTOMFIELD_10434:                issuesDto.Fields.Customfield10434,
		TIMESPENT:                        issuesDto.Fields.Timespent,
		CUSTOMFIELD_10433:                issuesDto.Fields.Customfield10433,
		WORKRATIO:                        issuesDto.Fields.Workratio,
		HORAS_TEST:                       issuesDto.Fields.HorasTest,
		RESUMO_ALTERACAO:                 issuesDto.Fields.ResumoAlteracao,
	}
}
