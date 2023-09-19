package dds

import "github.com/danielzinhors/api_jira/src/domain/model/dds"

type Project struct {
	Self           string          `json:"self"`
	ID             string          `json:"id"`
	Key            string          `json:"key"`
	Name           string          `json:"name"`
	ProjectTypeKey string          `json:"projectTypeKey"`
	AvatarUrls     *dds.AvatarUrls `json:"avatarUrls"`
}
