package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/STA"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics/WTL"
	"io/ioutil"
	"log"
	"os"
)

var wtlArgsSet *flag.FlagSet
var wtlProjectIdArg *string
var wtlBoardIdArg *string
var wtlSprintArg *string
var wtlPrintIssuesRefs *bool
var wtlEpicLink *string

func main() {
	args := os.Args[1:]

	res, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
	}

	wtlArgsSet = flag.NewFlagSet("wtl", flag.ExitOnError)

	wtlProjectIdArg = wtlArgsSet.String("projid", "", "Jira Project's ID")
	wtlBoardIdArg = wtlArgsSet.String("board", "", "Jira Board's Name or ID")
	wtlSprintArg = wtlArgsSet.String("sprint", "", "Jira Project's sprint ID or name")
	wtlEpicLink = wtlArgsSet.String("epic", "", "Jira Project's epic ID or name")
	wtlPrintIssuesRefs = wtlArgsSet.Bool("showIssuesRefs", false, "Print all issues references under the each assignee name")

	config := map[string]string{}

	err = json.Unmarshal(res, &config)

	if err != nil { log.Fatal(err) }

	str, err := startAnalytics(config, args)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*str)
}

func startAnalytics(config map[string]string, args []string) (*string, error) {

	userNodel := JiraAnalytics.JiraUserModel{
		Username: config["user"],
		Password: config["pass"],
	}

	switch args[0] {
	case "wtl":
		wtlArgsSet.Parse(args[1:])
		return CreateWhatTimeLeft(userNodel, WTL.WhatTimeLeftDefaultStringFormatter{})
	case "sta":
		wtlArgsSet.Parse(args[1:])
		return CreateSpendTimeAnalytics(userNodel, STA.SpendTimeAnalyticsDefaultStringFormatter{})
	}

	undefinedCmd := "Hmm.. seems like it's undefined command - " + args[0]
	return &undefinedCmd, nil
}