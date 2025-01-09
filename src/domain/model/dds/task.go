package dds

import (
	"time"
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
	COMPLEXIDADE                     Complexidade
	AGGREGATE_TIME_ORIGINAL_ESTIMATE float32
	AssineeId                        int64
	Assinee                          User
	ReporterId                       int64
	Reporter                         User
	IssuetypeId                      int64
	Issuetype                        Issuetype
	ProjectId                        int64
	Project                          Project
	Update                           time.Time
	CUSTOMFIELD_10105                string
	PriorityId                       int64
	Priority                         Priority
	SprintId                         int64
	Sprint                           Sprint
	SOLUCAO                          string
	TIME_ESTIMATE                    float32
	AGGREGATE_TIME_ESTIMATE          float32
	CreatorId                        int64
	Creator                          User
	RequerAnaliseTecnicaId           int64
	RequerAnaliseTecnica             RequerAnaliseTecnica
	CUSTOMFIELD_10434                string
	TIMESPENT                        float32
	AGGREGATE_TIMESPENT              float32
	CUSTOMFIELD_10432                string
	CUSTOMFIELD_10433                string
	WORKRATIO                        float32
	RESUMO_ALTERACAO                 string
	RequerDocumentacaoId             int64
	RequerDocumentacao               RequerDocumentacao
	ClassificacaoRelevanciaId        int64
	ClassificacaoRelevancia          ClassificacaoRelevancia
	TipoAlteracaoId                  int64
	TipoAlteracao                    TipoAlteracao
	HORAS_TEST                       float32
}
