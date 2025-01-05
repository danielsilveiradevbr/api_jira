package ddsDto

type Project struct {
	Self            string     `json:"self"`
	ID              string     `json:"id"`
	Key             string     `json:"key"`
	Name            string     `json:"name"`
	ProjectTypeKey  string     `json:"projectTypeKey"`
	AvatarUrls      AvatarUrls `json:"avatarUrls"`
	ProjectCategory ProjectCategory
}
