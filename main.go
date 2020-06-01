package main

import (
	"encoding/json"
	"fmt"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"
	"io/ioutil"
	"log"
	"os"

	jdbmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
)

func main() {
	args := os.Args[1:]

	res, err := ioutil.ReadFile("config.json")

	if err != nil {
		log.Fatal(err)
	}

	config := map[string]string{}

	err = json.Unmarshal(res, &config)

	if err != nil { log.Fatal(err) }

	if args[0] == "what_time_left" {
		jiraId := args[1]
		board := args[2]

		userNodel := JiraAnalytics.JiraUserModel{
			Username: config["user"],
			Password: config["pass"],
		}

		res, err := JiraAnalytics.StartWhatTimeLeft(userNodel, *jdbmod.NewBoardType(board), jiraId)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(*res)
	}
}