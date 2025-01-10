package projectDto

import (
	avatarUrlsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/avatarUrls"
	projectCategoryDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/ddsDto/projectCategory"
)

type Project struct {
	Self            string                             `json:"self"`
	ID              string                             `json:"id"`
	Key             string                             `json:"key"`
	Name            string                             `json:"name"`
	ProjectTypeKey  string                             `json:"projectTypeKey"`
	AvatarUrls      avatarUrlsDto.AvatarUrls           `json:"avatarUrls"`
	ProjectCategory projectCategoryDto.ProjectCategory `json:"projectCategory"`
}
