package dds

import "time"

type Project struct {
	ID           uint64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA      uint64
	KEY_JIRA     string
	DESCRICAO    string
	DATA_CRIACAO time.Time
	DATA_UPDATE  time.Time
}
