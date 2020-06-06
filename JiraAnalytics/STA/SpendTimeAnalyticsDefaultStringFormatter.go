package STA

import (
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	"github.com/olekukonko/tablewriter"
	"strings"
)

type SpendTimeAnalyticsDefaultStringFormatter struct {}

func (formatter SpendTimeAnalyticsDefaultStringFormatter) Handle(data SpendTimeAnalyticsResult, err error) (*string, error) {

	if err != nil { return nil, err}

	allIssues := strings.Builder{}

	allIssues.WriteString("\n")
	allIssues.WriteString("FOR THE WHOLE TEAM")
	allIssues.WriteString("\n")

	allIssues.WriteString(formatter.writeIssueTypes(data.AllIssuesMetrics).String())
	byPersonIssues := formatter.writeByAssignee(data.ByAssigneeMetrics)

	allIssues.WriteString("\n")
	allIssues.WriteString("***************************************************")
	allIssues.WriteString("\n")
	allIssues.WriteString(byPersonIssues.String())

	result := allIssues.String()

	return &result, nil
}

func (formatter SpendTimeAnalyticsDefaultStringFormatter) writeByAssignee(data map[string][]IssueGroupWithTimeMetrics) *strings.Builder {
	bld := strings.Builder{}

	for key, group := range data {
		bld.WriteString(strings.ToUpper(key))
		bld.WriteString("\n")
		tempBld := formatter.writeIssueTypes(group)
		bld.WriteString(tempBld.String())
		bld.WriteString("\n")
	}

	return &bld
}

func (formatter SpendTimeAnalyticsDefaultStringFormatter) writeIssueTypes(data []IssueGroupWithTimeMetrics) *strings.Builder {
	bld := strings.Builder{}

	table := tablewriter.NewWriter(&bld)

	for _, item := range data {
		row := []string{
			item.Group.ID,
			Common.TimeToStringView(item.SpendSum),
			Common.TimeToStringView(item.EstimateSum),
			fmt.Sprintf("%f", item.Accuracy),
		}
		table.Append(row)
	}

	table.SetHeader([]string{"Issue Type", "Sum Spent", "Sum Estimate", "Accuracy"})

	table.Render()
	return &bld
}