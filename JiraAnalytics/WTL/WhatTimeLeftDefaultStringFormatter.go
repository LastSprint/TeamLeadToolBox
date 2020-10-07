package WTL

import (
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	"strings"
)

/*
	WhatTimeLeftDefaultStringFormatter can handle result of WTL analytics and convert it to readable string representation
	This formatter used by default and its results are:
	```
	assignee_name1:	20h 20m
		https://jira.surfstudio.ru/browse/SPL-123
	assignee_name2:	8h 39m
	assignee_name3: 15m
		https://jira.surfstudio.ru/browse/SPL-124
		https://jira.surfstudio.ru/browse/SPL-125
		https://jira.surfstudio.ru/browse/SPL-126
 */
type WhatTimeLeftDefaultStringFormatter struct { }

// Handle works as already said in comments to `WhatTimeLeftDefaultStringFormatter`
func (formatter WhatTimeLeftDefaultStringFormatter) Handle(data []IssueGroupWithRemaining, err error, needsToPrintIssuesLink bool) (*string, error) {

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