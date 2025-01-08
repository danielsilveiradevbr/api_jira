package dds

import (
	"time"
)

type TASK struct {
	ID                  int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA             string
	KEY_JIRA            string `gorm:"unique"`
	DESCRICAO           string
	ID_USER             int64 `gorm:"foreignKey:UserID"`
	ID_REPORT           int64 `gorm:"foreignKey:UserID"`
	PROGRES_TOTAL       int64
	PROGRESS            int64
	PERC_PROGRESS       float64
	SUMMARY             string
	ID_PROJECT          int64 `gorm:"foreignKey:ProjetoID"`
	ID_PRIORITY         int64 `gorm:"foreignKey:PriorityID"`
	ID_STATUS           int64 `gorm:"foreignKey:StatusID"`
	TIPO                string
	ID_ISSUE_TYPE       int64 `gorm:"foreignKey:IssueTypeID"`
	CREATEDAT           time.Time
	ID_SPRINT           int64 `gorm:"foreignKey:SprintID"`
	QTDE_RETRABALHO     int
	ID_CLIENTE          int64 `gorm:"foreignKey:ClienteID"`
	DUPLICACAO_CODIGO   int
	PADRONIZACAO_CODIGO int
	DOCUMENTACAO_CODIGO int
	LEGIBILIDADE_CODIGO int
	SIMPLICIDADE_CODIGO int
	MODULARIDADE_CODIGO int
	TARGET_END          time.Time
	RESOLUTION_DATE     time.Time
	HORAS_PREC          float64
	COMPLEXIDADE        int
}
