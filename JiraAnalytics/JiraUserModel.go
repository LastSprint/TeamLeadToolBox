package JiraAnalytics

// This is jira model for basic auth.
// Yes, i just don't give a shit (:
// If you want to use Auth* methods so i just can contribute (:
//
// I promise I don't send your information to any place expect of Jira
//
// This model will use for get access to Jira
type JiraUserModel struct {
	// Username Jira user's name
	Username string
	// Password Jira user's password
	Password string
}
