package model

import "github.com/danielzinhors/api_jira/src/model"

type Project struct {
	Self           string            `json:"self"`
	ID             string            `json:"id"`
	Key            string            `json:"key"`
	Name           string            `json:"name"`
	ProjectTypeKey string            `json:"projectTypeKey"`
	AvatarUrls     *model.AvatarUrls `json:"avatarUrls"`
}
