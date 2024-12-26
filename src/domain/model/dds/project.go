package dds

import (
	"time"

	project "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
)

type Project struct {
	ID           int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA      string
	KEY_JIRA     string
	DESCRICAO    string
	DATA_CRIACAO time.Time
	DATA_UPDATE  time.Time
}

func NewProject(projetodto *project.Project) *Project {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	return &Project{
		ID_JIRA:      projetodto.ID,
		KEY_JIRA:     projetodto.Key,
		DESCRICAO:    projetodto.Name,
		DATA_CRIACAO: time.Now().In(loc),
	}
}
