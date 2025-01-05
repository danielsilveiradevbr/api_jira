package dds

import (
	"time"

	userDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	utils "github.com/danielsilveiradevbr/api_jira/src/utils"
)

type User struct {
	ID           int64 `gorm:"primaryKey;autoIncrement:true"`
	KEY_JIRA     string
	Email        string
	Avatar_16x16 string
	Avatar_24x24 string
	Avatar_32x32 string
	Avatar_48x48 string
	Nome         string
	DisplayName  string
	DATA_CRIACAO time.Time
	DATA_UPDATE  time.Time
}

func NewAssignee(assignee *userDto.Assignee) *User {
	return &User{
		KEY_JIRA:     assignee.Key,
		Email:        assignee.EmailAddress,
		Avatar_16x16: assignee.AvatarUrls.One6X16,
		Avatar_24x24: assignee.AvatarUrls.Two4X24,
		Avatar_32x32: assignee.AvatarUrls.Three2X32,
		Avatar_48x48: assignee.AvatarUrls.Four8X48,
		DisplayName:  assignee.DisplayName,
		Nome:         assignee.Name,
		DATA_CRIACAO: utils.GetADateTimeSaoPaulo(),
	}
}
