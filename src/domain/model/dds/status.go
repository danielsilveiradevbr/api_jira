package dds

import "time"

type Status struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA      uint64
	DESCRICAO    string
	DATA_CRIACAO time.Time
	DATA_UPDATE  time.Time
}
