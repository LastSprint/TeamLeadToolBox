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

	allIssues := writeIssueTypes(data.AllIssuesMetrics)
	byPersonIssues := writeByAssignee(data.ByAssigneeMetrics)

	allIssues.WriteString("\n")
	allIssues.WriteString("***************************************************")
	allIssues.WriteString("\n")
	allIssues.WriteString(byPersonIssues.String())

	result := allIssues.String()

	return &result, nil
}

func writeByAssignee(data map[string][]IssueGroupWithTimeMetrics) *strings.Builder {
	bld := strings.Builder{}

	for key, group := range data {
		bld.WriteString(key)
		bld.WriteString("\n")
		tempBld := writeIssueTypes(group)
		bld.WriteString(tempBld.String())
		bld.WriteString("\n")
	}

	return &bld
}

func writeIssueTypes(data []IssueGroupWithTimeMetrics) *strings.Builder {
	bld := strings.Builder{}

	table := tablewriter.NewWriter(&bld)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

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