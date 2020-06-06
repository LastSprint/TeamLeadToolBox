package JiraAnalytics

import (
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	jdbmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	"strings"

	// At this time this is my private package but I will change it soon :D
	jmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
	jsrv "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/services"
)

const jiraApiUrl = "https://jira.surfstudio.ru/rest/api/2/search"
const jiraWebUrl = "https://jira.surfstudio.ru"

// Start will collect information about each member of team in Jira.
// So this analytics will collect remaining time of each Bug/Task in `TODO` or `In Progress states`
// Then it will return this information as string (already formatted) or error
//
// Params:
//	- JiraUserModel: Your auth credentials for Jira
//	- board: The name of the board you want to collect information from (iOS/Android for example)
// 	- projectId: The jira key of the project. In EXM-123 this is the `EXM`
//	- sprint: The jira sprint id (or just a name)
//	- showIssuesRefs: if true then in result will printed issue links like `jiraWebUrl/browse/SPL-100`
func StartWhatTimeLeft(user JiraUserModel, board jdbmod.BoardType, projectId, sprint string, showIssuesRefs bool) (*string, error) {
	loader := jsrv.NewJiraIssueLoader(jiraApiUrl, user.Username, user.Password)

	request := Common.JiraSearchRequest{
		Wrapped: jsrv.SearchRequest{
			Board:                   board,
			IncludedStatuses:        []string{jmod.ToDo, jmod.InProgress},
			IncludedTypes:           []string{jmod.IssueTypeTask, jmod.IssueTypeBug},
			ProjectID:               projectId,
			AdditionFields: []jsrv.JiraField{},
		},
		Sprint:  sprint,
	}

	issues, err := loader.LoadIssues(request)

	if err != nil {
		return nil, err
	}

	groups := Common.GroupByAssignee(issues.Issues)

	data := getUserRemaining(groups)

	groupsToMap := map[string][]jmod.IssueEntity{}

	if showIssuesRefs {
		groupsToMap = groups
	}

	result := formUserTimeMap(data, groupsToMap)

	return &result, nil
}

func formUserTimeMap(val map[string]string, groups map[string][]jmod.IssueEntity) string {
	builder := strings.Builder{}

	for key, value := range val {
		builder.WriteString(fmt.Sprintf("%s:\t%s\n", key, value))

		if _, ok := groups[key]; !ok {
			continue
		}

		for _, issue := range groups[key] {
			builder.WriteString(fmt.Sprintf("  %s\n", jiraWebUrl + "/browse/" + issue.Key))
		}
	}

	return builder.String()
}

func getUserRemaining(groups map[string][]jmod.IssueEntity) map[string]string {
	result := make(map[string]string, len(groups))

	for key, group := range groups {
		timeRes := calculateRemaining(group)
		result[key] = Common.TimeToStringView(timeRes)
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