package Common

import (
	"fmt"
	jmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-jira-packages/models"
)

// GroupByAssignee will combine issues by value in `IssueEntity.Fields.Assignee.Name`
func GroupByAssignee(issues []jmod.IssueEntity) map[string][]jmod.IssueEntity {
	result := make(map[string][]jmod.IssueEntity, 0)

	for _, issue := range issues {

		name := issue.Fields.Assignee.Name

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