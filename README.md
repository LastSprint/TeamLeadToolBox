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

Then you need to run command `./TeamLeadToolBox what_time_left SMP iOS` for example and it may print:

```
board = iOS and project = SMP and status in ("To Do","In Progress") and issuetype in ("10002","10004")
YourTeamMember1: 	6h 0m
YourTeamMember2: 	4h 4m
YourTeamMember3: 	5h 0m
YourTeamMember4: 	4h 19m
....
```

If you don't want to use huge cmd command all time you may use this script:

```Shell
echo "./TeamLeadToolBox what_time_left <YOUR PROJECT> <BOARD>" >> run_wtl.sh
chmod +x run_wtl.sh
```

and then just

```Shell
./run_wtl.sh
```
