package dds

import "github.com/danielzinhors/api_jira/src/domain/model/dds"

type Creator struct {
	Self         string          `json:"self"`
	Name         string          `json:"name"`
	Key          string          `json:"key"`
	EmailAddress string          `json:"emailAddress"`
	AvatarUrls   *dds.AvatarUrls `json:"avatarUrls"`
	DisplayName  string          `json:"displayName"`
	Active       bool            `json:"active"`
	TimeZone     string          `json:"timeZone"`
}
