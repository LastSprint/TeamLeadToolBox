# TeamLeadToolBox
An item lead tool box that may be really helpful in job

## WhatTimeLeft

This tool collects all bugs and tasks from specific board in jira then it will group each issue by assignee and then just sum all remaning inside grop. And then it will just print it. 

So by this way you can find out what busy time each team member has

For using this tool you need to create `config.json` file near (in the same dir) the ToolBox binary:

```JSON
{
  "user": "Your Jira Account Nmae",
  "pass": "Your Jira Password"
}
```

Then you need to run command `./TeamLeadToolBox wtl -projid=SMP board=iOS -sprint="SMP-IOS 01 Init" -showIssuesRefs` for example and it may print:

```
board = iOS and project = SMP and status in ("To Do","In Progress") and issuetype in ("10002","10004")
YourTeamMember1: 	6h 0m
  https://jira.surfstudio.ru//browse/SMP-123
	https://jira.surfstudio.ru//browse/SMP-32
YourTeamMember2: 	0h 0m
YourTeamMember3: 	5h 0m
  https://jira.surfstudio.ru//browse/SMP-34
YourTeamMember4: 	4h 19m
  https://jira.surfstudio.ru//browse/SMP-1
	https://jira.surfstudio.ru//browse/SMP-2
  https://jira.surfstudio.ru//browse/SMP-3
	https://jira.surfstudio.ru//browse/SMP-4
....
```

If you don't want to use huge cmd command all time you may use this script:

```Shell
echo "./TeamLeadToolBox wtl -projid=<YOUR PROJECT> -board=<BOARD>" >> run_wtl.sh
chmod +x run_wtl.sh
```

and then just

```Shell
./run_wtl.sh
```

To see what flags exist for `wtl` just execute `./TeamLeadToolBox wtl -h`
