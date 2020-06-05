package main

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	"log"
)

func CreateWhatTimeLeft(user JiraAnalytics.JiraUserModel) (*string, error) {
	errMsg := ""

	if wtlProjectIdArg == nil {
		errMsg = "You must specify a Jira Project's ID"
		return &errMsg, nil
	}

	if wtlBoardIdArg == nil {
		errMsg = "You must specify the board's ID or name"
		return &errMsg, nil
	}

	sprint := ""

	if wtlSprintArg != nil {
		sprint = *wtlSprintArg
	}

	showIssuesRef := false

	if wtlPrintIssuesRefs != nil {
		showIssuesRef = *wtlPrintIssuesRefs
	}

	log.Print(*wtlPrintIssuesRefs)

	return JiraAnalytics.StartWhatTimeLeft(user, models.BoardType(*wtlBoardIdArg), *wtlProjectIdArg, sprint, showIssuesRef)
}