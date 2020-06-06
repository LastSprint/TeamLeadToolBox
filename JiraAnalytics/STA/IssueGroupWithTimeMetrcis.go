package STA

import "github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"

// IssueGroupWithTimeMetrics contains group and metrics about time spend on this issues
type IssueGroupWithTimeMetrics struct {
	Group Common.IssueGroup
	SpendSum int
	EstimateSum int
	Accuracy float64
}

// Result of TimeSpend analysis
type SpendTimeAnalyticsResult struct {
	AllIssuesMetrics []IssueGroupWithTimeMetrics
	// Key is assignee name you can print
	ByAssigneeMetrics map[string][]IssueGroupWithTimeMetrics
}