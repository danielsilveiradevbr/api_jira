package dds

import (
	userDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           int64 `gorm:"primaryKey;autoIncrement:true"`
	KEY_JIRA     string
	Email        string `gorm:"unique"`
	Avatar_16x16 string
	Avatar_24x24 string
	Avatar_32x32 string
	Avatar_48x48 string
	Nome         string
	DisplayName  string
}

func newUser(key, emailAddress, one6X16, two4x24, three2X32, four8X48, displayName, name string) *User {
	return &User{
		KEY_JIRA:     key,
		Email:        emailAddress,
		Avatar_16x16: one6X16,
		Avatar_24x24: two4x24,
		Avatar_32x32: three2X32,
		Avatar_48x48: four8X48,
		DisplayName:  displayName,
		Nome:         name,
	}
}

func NewAssignee(assignee *userDto.Assignee) *User {
	return newUser(
		assignee.Key,
		assignee.EmailAddress,
		assignee.AvatarUrls.One6X16,
		assignee.AvatarUrls.Two4X24,
		assignee.AvatarUrls.Three2X32,
		assignee.AvatarUrls.Four8X48,
		assignee.DisplayName,
		assignee.Name,
	)
}

func NewReporter(reporter *userDto.Reporter) *User {
	return newUser(
		reporter.Key,
		reporter.EmailAddress,
		reporter.AvatarUrls.One6X16,
		reporter.AvatarUrls.Two4X24,
		reporter.AvatarUrls.Three2X32,
		reporter.AvatarUrls.Four8X48,
		reporter.DisplayName,
		reporter.Name,
	)
}
