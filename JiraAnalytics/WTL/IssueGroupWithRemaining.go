package WTL

import "github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"

// IssueGroupWithRemaining provides information about remaining time for these groups of issues
type IssueGroupWithRemaining struct {
	Group Common.IssueGroup
	RemainingSum int
}