package assineeDto

import avatarUrlsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/avatarUrls"

type Assignee struct {
	Self         string                   `json:"self"`
	Name         string                   `json:"name"`
	Key          string                   `json:"key"`
	EmailAddress string                   `json:"emailAddress"`
	AvatarUrls   avatarUrlsDto.AvatarUrls `json:"avatarUrls"`
	DisplayName  string                   `json:"displayName"`
	Active       bool                     `json:"active"`
	TimeZone     string                   `json:"timeZone"`
}