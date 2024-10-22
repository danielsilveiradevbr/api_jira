package dds

import (
	"time"
)

type DDS struct {
	ID                  uint64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA             uint64
	KEY_JIRA            string
	DESCRICAO           string
	DATA_CRIACAO        time.Time
	DATA_UPDATE         time.Time
	ID_USER             uint64 `gorm:"foreignKey:UserID"`
	ID_REPORT           uint64 `gorm:"foreignKey:UserID"`
	PROGRES_TOTAL       uint64
	PROGRESS            uint64
	PERC_PROGRESS       float64
	SUMMARY             string
	ID_PROJECT          uint64 `gorm:"foreignKey:ProjetoID"`
	ID_PRIORITY         uint64 `gorm:"foreignKey:PriorityID"`
	ID_STATUS           uint64 `gorm:"foreignKey:StatusID"`
	TIPO                string
	ID_ISSUE_TYPE       uint64 `gorm:"foreignKey:IssueTypeID"`
	CREATEDAT           time.Time
	ID_SPRINT           uint64 `gorm:"foreignKey:SprintID"`
	QTDE_RETRABALHO     int
	ID_CLIENTE          uint64 `gorm:"foreignKey:ClienteID"`
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
