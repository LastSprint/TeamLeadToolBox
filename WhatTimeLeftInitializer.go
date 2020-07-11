package main

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/WTL"
)

func safeStr(val *string) string {
	if val != nil { return *val}
	return ""
}

func safeBool(val *bool) bool {
	if val != nil { return *val}
	return false
}

type WTLStringFormatter interface {
	Handle(data []WTL.IssueGroupWithRemaining, err error, needsToPrintIssuesLink bool) (*string, error)
}

func CreateWhatTimeLeft(user JiraAnalytics.JiraUserModel, formatter WTLStringFormatter) (*string, error) {
	errMsg := ""

	if wtlProjectIdArg == nil {
		errMsg = "You must specify a Jira Project's ID"
		return &errMsg, nil
	}

	if wtlBoardIdArg == nil {
		errMsg = "You must specify the board's ID or name"
		return &errMsg, nil
	}

	data, err := WTL.StartWhatTimeLeft(user, *wtlBoardIdArg, safeStr(wtlEpicLink), safeStr(wtlProjectIdArg), safeStr(wtlSprintArg))

	return formatter.Handle(data, err, safeBool(wtlPrintIssuesRefs))
}