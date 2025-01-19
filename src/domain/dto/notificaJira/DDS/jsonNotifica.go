package notificaJiraDto

import (
	changelogDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/changelog"
	issuesDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/issues"
	userDto "github.com/danielsilveiradevbr/api_jira/src/domain/dto/auxiliar/user"
)

type JsonNotifica struct {
	Timestamp          int64                  `json:"timestamp"`
	WebhookEvent       string                 `json:"webhookEvent"`
	IssueEventTypeName string                 `json:"issue_event_type_name"`
	User               userDto.User           `json:"user"`
	Issue              issuesDto.Issues       `json:"issue"`
	Changelog          changelogDto.Changelog `json:"changelog"`
}
