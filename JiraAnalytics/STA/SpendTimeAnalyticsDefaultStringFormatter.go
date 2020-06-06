package STA

import (
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	"strings"
)

type SpendTimeAnalyticsDefaultStringFormatter struct {}

func (formatter SpendTimeAnalyticsDefaultStringFormatter) Handle(data SpendTimeAnalyticsResult, err error) (*string, error) {

	const jiraWebUrl = "https://jira.surfstudio.ru"

	if err != nil { return nil, err}

	builder := strings.Builder{}

	for _, group := range data {
		builder.WriteString(fmt.Sprintf("%s:\t%s\n", group.Group.ID, Common.TimeToStringView(group.RemainingSum)))

		if !needsToPrintIssuesLink { continue }

		for _, issue := range group.Group.Issues {
			builder.WriteString(fmt.Sprintf("  %s\n", jiraWebUrl + "/browse/" + issue.Key))
		}
	}

	result := builder.String()

	return &result, nil
}