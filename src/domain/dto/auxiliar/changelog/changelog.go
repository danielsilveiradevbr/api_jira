package changelogDto

import changelogItemsDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/changelogItems"

type Changelog struct {
	ID    string                    `json:"id"`
	Items []changelogItemsDto.Items `json:"items"`
}
