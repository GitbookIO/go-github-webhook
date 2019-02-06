package github

type GitCommit struct {
	Author    string `json:"author"`
	Email     string `json:"email"`
	Repo      string `json:"repo"`
	RepoOwner string `json:"repo_owner"`
	Message   string `json:"message"`
	Date      string `json:"date"`
	Hash      string `json:"hash"`
}

type GitHubCommit struct {
	Distinct  bool            `json:"distinct"`
	URL       string          `json:"url"`
	Id        string          `json:"id"`
	Timestamp GitHubTimestamp `json:"timestamp"`
	Added     []string        `json:"added"`
	Message   string          `json:"message"`
	Committer GitHubPerson    `json:"committer"`
	Author    GitHubPerson    `json:"author"`
	Modified  []string        `json:"modified"`
	Removed   []string        `json:"removed"`
}

type GitHubRepo struct {
	HasIssues     bool            `json:"has_issues"`
	HasWiki       bool            `json:"has_wiki"`
	Size          int             `json:"size"`
	Description   string          `json:"description"`
	Owner         GitHubPerson    `json:"owner"`
	Homepage      string          `json:"homepage"`
	Watchers      int             `json:"watchers"`
	Language      string          `json:"language"`
	PushedAt      GitHubTimestamp `json:"pushed_at"`
	Name          string          `json:"name"`
	FullName      string          `json:"full_name"`
	Organization  string          `json:"organization"`
	HasDownloads  bool            `json:"has_downloads"`
	CreatedAt     GitHubTimestamp `json:"created_at"`
	URL           string          `json:"url"`
	CloneURL      string          `json:"clone_url"`
	OpenIssues    int             `json:"open_issues"`
	Forks         int             `json:"forks"`
	Private       bool            `json:"private"`
	Fork          bool            `json:"fork"`
	Stargazers    int             `json:"stargazers"`
	DefaultBranch string          `json:"default_branch"`
}

type GitHubPerson struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	Username *string `json:"username"`
}

type GitHubPayload struct {
	Before     string         `json:"before"`
	Created    bool           `json:"created"`
	Ref        string         `json:"ref"`
	Deleted    bool           `json:"deleted"`
	After      string         `json:"after"`
	HeadCommit GitHubCommit   `json:"head_commit"`
	Commits    []GitHubCommit `json:"commits"`
	Repository GitHubRepo     `json:"repository"`
	Forced     bool           `json:"forced"`
	Compare    string         `json:"compare"`
	Pusher     GitHubPerson   `json:"pusher"`

	Action       string                  `json:"action"`
	CheckSuite   GitHubV3CheckSuite      `json:"check_suite"`
	Installation GitHubV3AppInstallation `json:"installation"`
	Sender       GitHubV3Entity          `json:"sender"`
	Organization GitHubV3Entity          `json:"organization"`
}

type PayloadPong struct {
	Ok    bool   `json:"ok"`
	Event string `json:"event"`
	Error string `json:"error,omitempty"`
}

type GitHubV3AppInstallation struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
}

type GitHubV3Entity struct {
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	Type              string `json:"type"`
	Login             string `json:"login"`
	SideAdmin         bool   `json:"site_admin"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTLMURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionURL   string `json:"subscriptions_url"`
	OrganizationURL   string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
}

type GitHubV3CheckSuite struct {
	ID                 int    `json:"id"`
	NodeID             string `json:"node_id"`
	HeadBranch         string `json:"head_branch"`
	HeadSHA            string `json:"head_sha"`
	Status             string `json:"status"`
	LastCheckRunsCount int    `json:"latest_check_runs_count"`
	CheckRunsURL       string `json:"check_runs_url"`
}
