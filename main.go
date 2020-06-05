package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"io/ioutil"
	"log"
)

const PROJECT_ID = "projid"
const BOARD = "board"

var projectIdArg *string
var boardArg *string

func main() {
	initializeCmdArgsParser()
	res, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
	}

	config := map[string]string{}

	err = json.Unmarshal(res, &config)

	if err != nil { log.Fatal(err) }

	str, err := startAnalytics(config, flag.Args())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*str)
}

func initializeCmdArgsParser() {
	projectIdArg = flag.String(PROJECT_ID, "TRI", "Jira Project's ID")
	boardArg = flag.String(BOARD, "iOS", "Jira Board's Name or ID")

	flag.Parse()
}

func startAnalytics(config map[string]string, args []string) (*string, error) {

	switch args[0] {
	case "wtl":
		userNodel := JiraAnalytics.JiraUserModel{
			Username: config["user"],
			Password: config["pass"],
		}
		return CreateWhatTimeLeft(userNodel)
	}

	undefinedCmd := "Hmm.. seems like it's undefined command - " + args[0]
	return &undefinedCmd, nil
}