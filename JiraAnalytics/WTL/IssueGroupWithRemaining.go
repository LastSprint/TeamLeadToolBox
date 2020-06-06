package WTL

import "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"

// IssueGroupWithRemaining just a type to storing domain model and calculations about this model
// By meaning this model should be used only as DTO - so we can use it to send processed data
// to some entity that can make it readable
type IssueGroupWithRemaining struct {
	// PersonName is person who is assignee for this issues
	PersonName string
	// Issues of person
	Issues []models.IssueEntity
	// This is the sum of issues remaining
	RemainingSum int
}