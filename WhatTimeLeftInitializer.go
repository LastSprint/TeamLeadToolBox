package main

import (
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/WTL"
	"gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	"log"
)

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

	sprint := ""

	if wtlSprintArg != nil {
		sprint = *wtlSprintArg
	}

	showIssuesRef := false

	if wtlPrintIssuesRefs != nil {
		showIssuesRef = *wtlPrintIssuesRefs
	}

	log.Print(*wtlPrintIssuesRefs)

	data, err := WTL.StartWhatTimeLeft(user, models.BoardType(*wtlBoardIdArg), *wtlProjectIdArg, sprint)

	return formatter.Handle(data, err, showIssuesRef)
}