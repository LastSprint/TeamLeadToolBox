package Common

import (
	"fmt"
	jmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
)

// GroupBy contains all boilerplate code for grouping elements by one key
// - Parameters:
//		- issues: Elements to group
//		- fldToGroup: Provides value for grouping
//
// For example you want to group by assignee:
//
// ```Go
// GroupBy(issues, func(issue jmod.IssueEntity) string { return issue.Fields.Assignee.Name })
// ```
func GroupBy(issues []jmod.IssueEntity, fldToGroup func(jmod.IssueEntity)string) map[string][]jmod.IssueEntity {
	result := make(map[string][]jmod.IssueEntity, 0)

	for _, issue := range issues {

		name := fldToGroup(issue)

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

// GroupByAssignee will combine issues by value in `IssueEntity.Fields.Assignee.Name`
func GroupByAssignee(issues []jmod.IssueEntity) map[string][]jmod.IssueEntity {
	return GroupBy(issues, func(issue jmod.IssueEntity) string { return issue.Fields.Assignee.Name })
}

// GroupIssuesByType groups issues by their issue type ID
func GroupIssuesByType(issues []jmod.IssueEntity) map[string][]jmod.IssueEntity {
	return GroupBy(issues, func(issue jmod.IssueEntity) string { return issue.Fields.Issuetype.ID })
}

// ToIssuesGroups converts grouped issues to array of `IssueGroup`
// This method uses map key as ID
func ToIssuesGroups(groups map[string][]jmod.IssueEntity) []IssueGroup {

	result := make([]IssueGroup, len(groups))

	index := 0

	for key, value := range groups {
		result[index] = IssueGroup{
			ID:         key,
			Issues:     value,
		}
		index++
	}

	return result
}

// TimeToStringView converts time in seconds to string with next formats:
// 	- seconds < 1m -> `$seconds sec`
//	- seconds < 1h -> `($seconds/60)m`
//	- `($seconds/60/60)h ($seconds/60 % 60)m`
//
// For example:
//	- 10 seconds -> `10 sec`
//	- 75 seconds -> `1m`
//	- 5100 seconds -> `1h 25m`
func TimeToStringView(seconds int) string {
	if seconds < 60 {
		return fmt.Sprintf("%d sec", seconds)
	} else if seconds < (60 * 60) {
		return fmt.Sprintf("%dm", seconds/60)
	} else {
		return fmt.Sprintf("%dh %dm", seconds/60/60, seconds/60 % 60)
	}
}