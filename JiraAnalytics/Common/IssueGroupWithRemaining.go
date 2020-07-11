package Common

import "github.com/LastSprint/JiraGoIssues/models"

// IssueGroupWithRemaining just a type to storing domain model and calculations about this model
// By meaning this model should be used only as DTO - so we can use it to send processed data
// to some entity that can make it readable
type IssueGroup struct {
	// ID this is the by that we can union this issues
	// For example if we wont to group issues by their assignee then it would be Assignee.Name
	ID string
	// Issues of person
	Issues []models.IssueEntity
}