package model

import "github.com/danielzinhors/api_jira/src/model"

type Reporter struct {
	Self         string            `json:"self"`
	Name         string            `json:"name"`
	Key          string            `json:"key"`
	EmailAddress string            `json:"emailAddress"`
	AvatarUrls   *model.AvatarUrls `json:"avatarUrls"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
	TimeZone     string            `json:"timeZone"`
}
