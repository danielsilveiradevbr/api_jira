package resolutionModel

import (
	resolutionDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/resolution"
	"gorm.io/gorm"
)

type Resolution struct {
	gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	Url       string
	Descricao string
	Name      string `gorm:"unique;not null"`
}

func NewResolution(resolutionDTO *resolutionDto.Resolution) *Resolution {
	return &Resolution{
		Url:       resolutionDTO.Self,
		Descricao: resolutionDTO.Description,
		Name:      resolutionDTO.Name,
	}
}
