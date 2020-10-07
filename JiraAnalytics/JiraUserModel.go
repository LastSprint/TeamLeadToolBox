package JiraAnalytics

// This is jira model for basic auth.
// Yes, i just don't give a shit (:
// If you want to use Auth* methods so you can contribute (:
//
// I promise I won't send your information to any place except of Jira
//
// This model will be used for getting access to Jira
type JiraUserModel struct {
	// Username Jira user's name
	Username string
	// Password Jira user's password
	Password string
}
