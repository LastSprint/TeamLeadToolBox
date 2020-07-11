package STA

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/Common"
	"github.com/LastSprint/TeamLeadToolBox/Tools"
	jmod "github.com/LastSprint/JiraGoIssues/models"
	jsrv "github.com/LastSprint/JiraGoIssues/services"
)

const jiraApiUrl = "https://jira.surfstudio.ru/rest/api/2/search"

func StartSpendTimeAnalytics(user JiraAnalytics.JiraUserModel, board string, epicLink, projectId, sprint string) (SpendTimeAnalyticsResult, error) {
	loader := jsrv.NewJiraIssueLoader(jiraApiUrl, user.Username, user.Password)

	request := Common.JiraSearchRequest{
		Wrapped: jsrv.SearchRequest{
			Board:                   board,
			IncludedTypes:           []string{jmod.IssueTypeTask, jmod.IssueTypeBug, jmod.IssueTypeServiceTask},
			ProjectID:               projectId,
			EpicLink:				 epicLink,
			AdditionFields: []jsrv.JiraField{},
		},
		Sprint:  sprint,
	}

	issues, err := loader.LoadIssues(request)

	if err != nil { return SpendTimeAnalyticsResult{}, err }

	// Global groups that didn't split by assignee
	allIssuesCh := make(chan interface{})
	byAssigneeCh := make(chan interface{})

	go Tools.PerformWithChan(allIssuesCh, func() interface{} {
		return makeIssueTypeProcessing(issues.Issues)
	})

	go Tools.PerformWithChan(byAssigneeCh, func() interface{} {
		return makeProcessingByAssignee(issues.Issues)
	})

	allIssuesMetrics := <- allIssuesCh
	byAssigneeMetrics := <- byAssigneeCh
	
	return SpendTimeAnalyticsResult{
		AllIssuesMetrics:  allIssuesMetrics.([]IssueGroupWithTimeMetrics),
		ByAssigneeMetrics: byAssigneeMetrics.(map[string][]IssueGroupWithTimeMetrics),
	}, nil
}

func makeProcessingByAssignee(issues []jmod.IssueEntity) map[string][]IssueGroupWithTimeMetrics {
	grouped := Common.GroupByAssignee(issues)

	result := make(map[string][]IssueGroupWithTimeMetrics, len(grouped))

	for key, group := range grouped {
		result[key] = makeIssueTypeProcessing(group)
	}

	return result
}

// makeIssueTypeProcessing awaiting an array of issues
// then it will split this array on numbers of array by their `IssueType.ID`
// And after that it will collect some analytics
func makeIssueTypeProcessing(issues []jmod.IssueEntity) []IssueGroupWithTimeMetrics {
	grouped := Common.ToIssuesGroups(Common.GroupIssuesByType(issues))

	result := make([]IssueGroupWithTimeMetrics, len(grouped))

	for i, group := range grouped {
		result[i] = makeGroupCalculation(group)
	}

	return result
}

// makeGroupCalculation will calculate this pipeline metrics for passed issues
func makeGroupCalculation(group Common.IssueGroup) IssueGroupWithTimeMetrics {
	allTimeSpend := Common.Sum(group.Issues, func(issue jmod.IssueEntity) int { return issue.Fields.TimeSpend })
	allEstimate := Common.Sum(group.Issues, func(issue jmod.IssueEntity) int { return issue.Fields.Estimate })
	var accuracy float64

	if allTimeSpend == 0 {
		accuracy = -1
	} else {
		accuracy = float64(allEstimate) / float64(allTimeSpend)
	}

	return IssueGroupWithTimeMetrics{
		Group:       group,
		SpendSum:    allTimeSpend,
		EstimateSum: allEstimate,
		Accuracy:    accuracy,
	}
}