package WTL

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	jmod "github.com/LastSprint/JiraGoIssues/models"
	jsrv "github.com/LastSprint/JiraGoIssues/services"
)

const jiraApiUrl = "https://jira.surfstudio.ru/rest/api/2/search"

// Start will collect information about each member of team in Jira.
// So this analytics will collect remaining time of each Bug/Task in `TODO` or `In Progress' states
// Then it will return this information as string (already formatted) or error
//
// Params:
//	- JiraUserModel: Your auth credentials for Jira
//	- board: The name of the board you want to collect information from (iOS/Android for example)
// 	- projectId: The jira key of the project. In EXM-123 this is the `EXM`
//	- sprint: The jira sprint id (or just a name). If you don't want to use it just pass it as empty string
//	- epicLink: This jira epic name or key (like SPL-100). If you don't want to use it just pass it as empty string
func StartWhatTimeLeft(user JiraAnalytics.JiraUserModel, board string, epicLink, projectId, sprint string) ([]IssueGroupWithRemaining, error) {
	loader := jsrv.NewJiraIssueLoader(jiraApiUrl, user.Username, user.Password)

	request := Common.JiraSearchRequest{
		Wrapped: jsrv.SearchRequest{
			Board:                   board,
			IncludedStatuses:        []string{jmod.ToDo, jmod.InProgress},
			IncludedTypes:           []string{jmod.IssueTypeTask, jmod.IssueTypeBug},
			ProjectID:               projectId,
			EpicLink:				 epicLink,
			AdditionFields: []jsrv.JiraField{},
		},
		Sprint:  sprint,
	}

	issues, err := loader.LoadIssues(request)

	if err != nil {
		return nil, err
	}

	return getUserRemaining(Common.GroupByAssignee(issues.Issues)), nil
}

// getUserRemaining just calculate remaining for each issues group
// So you should pass only grouped by assignee issues
func getUserRemaining(groups map[string][]jmod.IssueEntity) []IssueGroupWithRemaining {

	grouped := Common.ToIssuesGroups(groups)
	result := make([]IssueGroupWithRemaining, len(grouped))

	index := 0

	for _, group := range grouped {

		remaining := Common.Sum(group.Issues, func(issue jmod.IssueEntity) int { return issue.Fields.Remaining })

		result[index] = IssueGroupWithRemaining{
			Group: group,
			RemainingSum: remaining,
		}
		index++
	}

	return result
}