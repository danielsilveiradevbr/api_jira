package dds

import (
	project "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID                int64 `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA           string
	KEY_JIRA          string
	DESCRICAO         string `gorm:"unique;not null"`
	Avatar_16x16      string
	Avatar_24x24      string
	Avatar_32x32      string
	Avatar_48x48      string
	projectTypeKey    string
	ProjectCategoryID int64
	ProjectCategory   ProjectCategory
}

func NewProject(projetodto *project.Project) *Project {
	return &Project{
		ID_JIRA:        projetodto.ID,
		KEY_JIRA:       projetodto.Key,
		DESCRICAO:      projetodto.Name,
		Avatar_16x16:   projetodto.AvatarUrls.One6X16,
		Avatar_24x24:   projetodto.AvatarUrls.Two4X24,
		Avatar_32x32:   projetodto.AvatarUrls.Three2X32,
		Avatar_48x48:   projetodto.AvatarUrls.Four8X48,
		projectTypeKey: projetodto.ProjectTypeKey,
	}
}
