package TeamLeadToolBox

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"github.com/LastSprint/TeamLeadToolBox/JiraAnalytics"

	jdbmod "gitlab.com/surfstudio/infrastructure/spa/spa-backend-com-packages/dbservices/models"
	)

func Main() {
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

		JiraAnalytics.StartWhatTimeLeft(userNodel, *jdbmod.NewBoardType(board), jiraId)
	}
}