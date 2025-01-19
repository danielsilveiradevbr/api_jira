package statusDto

import statusCategoryDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/statusCategory"

type Status struct {
	Self           string                           `json:"self"`
	Description    string                           `json:"description"`
	IconURL        string                           `json:"iconUrl"`
	Name           string                           `json:"name"`
	ID             string                           `json:"id"`
	StatusCategory statusCategoryDto.StatusCategory `json:"statusCategory"`
}
