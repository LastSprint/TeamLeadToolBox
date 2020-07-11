package main

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/STA"
)

type STAFormatter interface {
	Handle(data STA.SpendTimeAnalyticsResult, err error) (*string, error)
}

func CreateSpendTimeAnalytics(user JiraAnalytics.JiraUserModel, formatter STAFormatter) (*string, error) {
	errMsg := ""

	if wtlProjectIdArg == nil {
		errMsg = "You must specify a Jira Project's ID"
		return &errMsg, nil
	}

	if wtlBoardIdArg == nil {
		errMsg = "You must specify the board's ID or name"
		return &errMsg, nil
	}

	data, err := STA.StartSpendTimeAnalytics(user, *wtlBoardIdArg, safeStr(wtlEpicLink), safeStr(wtlProjectIdArg), safeStr(wtlSprintArg))

	return formatter.Handle(data, err)
}