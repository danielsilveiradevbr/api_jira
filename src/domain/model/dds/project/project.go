package projectModel

import (
	projectDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/project"

	projectCategoryModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/dds/projectCategory"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID                uint `gorm:"primaryKey;autoIncrement:true"`
	ID_JIRA           string
	KEY_JIRA          string
	DESCRICAO         string `gorm:"unique;not null"`
	Avatar_16x16      string
	Avatar_24x24      string
	Avatar_32x32      string
	Avatar_48x48      string
	projectTypeKey    string
	ProjectCategoryID *uint
	ProjectCategory   projectCategoryModel.ProjectCategory
}

func (p *Project) BeforeCreate(db *gorm.DB) (err error) {
	if *p.ProjectCategoryID == 0 {
		p.ProjectCategoryID = nil
	}
	return nil
}

func NewProject(projetodto *projectDto.Project) *Project {
	return &Project{
		ID_JIRA:           projetodto.ID,
		KEY_JIRA:          projetodto.Key,
		DESCRICAO:         projetodto.Name,
		Avatar_16x16:      projetodto.AvatarUrls.One6X16,
		Avatar_24x24:      projetodto.AvatarUrls.Two4X24,
		Avatar_32x32:      projetodto.AvatarUrls.Three2X32,
		Avatar_48x48:      projetodto.AvatarUrls.Four8X48,
		projectTypeKey:    projetodto.ProjectTypeKey,
		ProjectCategoryID: nil,
	}
}
