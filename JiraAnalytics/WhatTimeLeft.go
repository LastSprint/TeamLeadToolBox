package JiraAnalytics

import (
	"fmt"
	jdbmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	"strings"

	// At this time this is my private package but I will change it soon :D
	jmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
	jsrv "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/services"
)

const JIRA_URL = "https://jira.surfstudio.ru/rest/api/2/search"

// Start will collect information about each member of team in Jira.
// So this analytics will collect remaining time of each Bug/Task in `TODO` or `In Progress states`
// Then it will return this information as string (already formatted) or error
//
// Params:
//	- JiraUserModel: Your auth credentials for Jira
//	- board: The name of the board you want to collect information from (iOS/Android for example)
// 	- projectId: The jira key of the project. In EXM-123 this is the `EXM`
func StartWhatTimeLeft(user JiraUserModel, board jdbmod.BoardType, projectId string) (*string, error) {
	loader := jsrv.NewJiraIssueLoader(JIRA_URL, user.Username, user.Password)

	request := jsrv.SearchRequest{
		Board:                   board,
		IncludedStatuses:        []string{jmod.ToDo, jmod.InProgress},
		IncludedTypes:           []string{jmod.IssueTypeTask, jmod.IssueTypeBug},
		ProjectID:               projectId,
	}

	issues, err := loader.LoadIssues(request)

	if err != nil {
		return nil, err
	}

	data := getUserRemaining(issues.Issues)

	result := formUserTimeMap(data)

	return &result, nil
}

func formUserTimeMap(val map[string]string) string {
	builder := strings.Builder{}

	for key, value := range val {
		builder.WriteString(fmt.Sprintf("%s: \t%s\n", key, value))
	}

	return builder.String()
}

func getUserRemaining(issues []jmod.IssueEntity) map[string]string {
	groups := groupByAssignee(issues)

	result := make(map[string]string, len(groups))

	for key, group := range groups {
		timeRes := calculateRemaining(group)
		result[key] = timeToStringView(timeRes)
	}

	return result
}

func groupByAssignee(issues []jmod.IssueEntity) map[string][]jmod.IssueEntity {
	result := make(map[string][]jmod.IssueEntity, 0)

	for _, issue := range issues {

		name := issue.Fields.Assignee.Name

		if len(name) == 0 { continue }

		if _, ok := result[name]; !ok {
			result[name] = make([]jmod.IssueEntity, 1)
			result[name][0] = issue
		} else {
			result[name] = append(result[name], issue)
		}
	}

	return result
}

func calculateRemaining(issues []jmod.IssueEntity) int {
	result := 0

	for _, item := range issues {
		result += item.Fields.Remaining
	}

	return result
}

func timeToStringView(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%d sec", seconds)
	} else if seconds < (60 * 60) {
		return fmt.Sprintf("%dm", seconds/60)
	} else {
		return fmt.Sprintf("%dh %dm", seconds/60/60, seconds/60 % 60)
	}
}