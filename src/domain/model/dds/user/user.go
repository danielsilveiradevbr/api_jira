package userModel

import (
	tipoUserDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/tipoUser"
	assineeDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/assineee"
	creatorDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/creator"
	reporterDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/reporter"
	tipoUserModel "github.com/danielsilveiradevbr/api_jira/src/domain/model/auxiliar/tipoUser"
	tipoUserRep "github.com/danielsilveiradevbr/api_jira/src/repositories/auxiliares/tipoUser"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int64 `gorm:"primaryKey;autoIncrement:true"`
	KEY_JIRA     string
	Email        string `gorm:"unique;not null"`
	Avatar_16x16 string
	Avatar_24x24 string
	Avatar_32x32 string
	Avatar_48x48 string
	Nome         string
	DisplayName  string
	TipoUserId   int64
	TipoUser     tipoUserModel.TipoUser
}

func newUser(db *gorm.DB, key, emailAddress, one6X16, two4x24, three2X32, four8X48, displayName, name string, tipoDeUser int) *User {
	tipo := "Assignee"
	if tipoDeUser == 2 {
		tipo = "Reporter"
	} else if tipoDeUser == 3 {
		tipo = "Creator"
	}
	tipoUser, err := tipoUserRep.SalvaTipoUser(db, &tipoUserDto.TipoUser{
		Nome: tipo,
	})
	if err != nil {
		return nil
	}

	return &User{
		KEY_JIRA:     key,
		Email:        emailAddress,
		Avatar_16x16: one6X16,
		Avatar_24x24: two4x24,
		Avatar_32x32: three2X32,
		Avatar_48x48: four8X48,
		DisplayName:  displayName,
		Nome:         name,
		TipoUserId:   tipoUser.ID,
	}
}

func NewAssignee(db *gorm.DB, assignee *assineeDto.Assignee) *User {
	return newUser(
		db,
		assignee.Key,
		assignee.EmailAddress,
		assignee.AvatarUrls.One6X16,
		assignee.AvatarUrls.Two4X24,
		assignee.AvatarUrls.Three2X32,
		assignee.AvatarUrls.Four8X48,
		assignee.DisplayName,
		assignee.Name,
		1,
	)
}

func NewReporter(db *gorm.DB, reporter *reporterDto.Reporter) *User {
	return newUser(
		db,
		reporter.Key,
		reporter.EmailAddress,
		reporter.AvatarUrls.One6X16,
		reporter.AvatarUrls.Two4X24,
		reporter.AvatarUrls.Three2X32,
		reporter.AvatarUrls.Four8X48,
		reporter.DisplayName,
		reporter.Name,
		2,
	)
}

func NewCreator(db *gorm.DB, creator *creatorDto.Creator) *User {
	return newUser(
		db,
		creator.Key,
		creator.EmailAddress,
		creator.AvatarUrls.One6X16,
		creator.AvatarUrls.Two4X24,
		creator.AvatarUrls.Three2X32,
		creator.AvatarUrls.Four8X48,
		creator.DisplayName,
		creator.Name,
		3,
	)
}
