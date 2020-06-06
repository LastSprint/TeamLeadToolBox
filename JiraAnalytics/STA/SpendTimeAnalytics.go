package STA

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	jdbmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	jmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
	jsrv "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/services"
)

const jiraApiUrl = "https://jira.surfstudio.ru/rest/api/2/search"

func StartSpendTimeAnalytics(user JiraAnalytics.JiraUserModel, board jdbmod.BoardType, epicLink, projectId, sprint string) error {
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

	grouped := Common.GroupIssuesByType(issues.Issues)
	return nil
}