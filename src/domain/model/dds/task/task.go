package taskModel

import (
	"time"

	classificacaoRelevanciaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/classificacaoRelevancia"
	complexidadeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/complexidade"
	issueTypeModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/issueType"
	priorityModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/priority"
	projectModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/project"
	requerAnaliseTecnicaModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerAnaliseTecnica"
	requerDocumentacaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/requerDocumentacao"
	sprintModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/sprint"
	tipoAlteracaoModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/tipoAlteracao"
	userModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/user"
)

type TASK struct {
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
	IssuetypeId                      int64
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
}
