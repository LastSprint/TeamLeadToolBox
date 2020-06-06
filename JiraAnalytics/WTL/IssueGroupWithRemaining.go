package WTL

import "github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"

// IssueGroupWithRemaining provides information about what remaining time for this groups of issues
type IssueGroupWithRemaining struct {
	Group Common.IssueGroup
	RemainingSum int
}