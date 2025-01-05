package dds

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA   uint64
	DESCRICAO string `gorm:"unique"`
}
