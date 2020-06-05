package main

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
)

func CreateWhatTimeLeft(user JiraAnalytics.JiraUserModel) (*string, error) {
	errMsg := ""

	if projectIdArg == nil {
		errMsg = "You must specify a Jira Project's ID"
		return &errMsg, nil
	}

	if boardArg == nil {
		errMsg = "You must specify the board's ID or name"
		return &errMsg, nil
	}

	return JiraAnalytics.StartWhatTimeLeft(user, models.BoardType(*boardArg), *projectIdArg)
}