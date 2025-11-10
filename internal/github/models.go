package github

import "time"

// IssueSearchResult represents the response from GitHub's issue search API
type IssueSearchResult struct {
	TotalCount        int                     `json:"total_count"`
	IncompleteResults bool                    `json:"incomplete_results"`
	Items             []IssueSearchResultItem `json:"items"`
}

// IssueSearchResultItem represents a single issue in the search results
type IssueSearchResultItem struct {
	URL                      string                  `json:"url"`
	RepositoryURL            string                  `json:"repository_url"`
	LabelsURL                string                  `json:"labels_url"`
	CommentsURL              string                  `json:"comments_url"`
	EventsURL                string                  `json:"events_url"`
	HTMLURL                  string                  `json:"html_url"`
	ID                       int64                   `json:"id"`
	NodeID                   string                  `json:"node_id"`
	Number                   int                     `json:"number"`
	Title                    string                  `json:"title"`
	Locked                   bool                    `json:"locked"`
	ActiveLockReason         *string                 `json:"active_lock_reason"`
	Assignees                []SimpleUser            `json:"assignees"`
	User                     *SimpleUser             `json:"user"`
	Labels                   []Label                 `json:"labels"`
	SubIssuesSummary         *SubIssuesSummary       `json:"sub_issues_summary,omitempty"`
	IssueDependenciesSummary *IssueDependenciesSummary `json:"issue_dependencies_summary,omitempty"`
	IssueFieldValues         []IssueFieldValue       `json:"issue_field_values,omitempty"`
	State                    string                  `json:"state"`
	StateReason              *string                 `json:"state_reason"`
	Assignee                 *SimpleUser             `json:"assignee"`
	Milestone                *Milestone              `json:"milestone"`
	Comments                 int                     `json:"comments"`
	CreatedAt                time.Time               `json:"created_at"`
	UpdatedAt                time.Time               `json:"updated_at"`
	ClosedAt                 *time.Time              `json:"closed_at"`
	TextMatches              []TextMatch             `json:"text_matches,omitempty"`
	PullRequest              *PullRequestRef         `json:"pull_request,omitempty"`
	Body                     string                  `json:"body"`
	Score                    float64                 `json:"score"`
	AuthorAssociation        string                  `json:"author_association"`
	Draft                    bool                    `json:"draft,omitempty"`
	Repository               Repository              `json:"repository"`
	BodyHTML                 string                  `json:"body_html,omitempty"`
	BodyText                 string                  `json:"body_text,omitempty"`
	TimelineURL              string                  `json:"timeline_url,omitempty"`
	Type                     *IssueType              `json:"type"`
	PerformedViaGithubApp    *GitHubApp              `json:"performed_via_github_app"`
	Reactions                *ReactionRollup         `json:"reactions,omitempty"`
}

// SimpleUser represents a GitHub user
type SimpleUser struct {
	Name              *string `json:"name"`
	Email             *string `json:"email"`
	Login             string  `json:"login"`
	ID                int64   `json:"id"`
	NodeID            string  `json:"node_id"`
	AvatarURL         string  `json:"avatar_url"`
	GravatarID        *string `json:"gravatar_id"`
	URL               string  `json:"url"`
	HTMLURL           string  `json:"html_url"`
	FollowersURL      string  `json:"followers_url"`
	FollowingURL      string  `json:"following_url"`
	GistsURL          string  `json:"gists_url"`
	StarredURL        string  `json:"starred_url"`
	SubscriptionsURL  string  `json:"subscriptions_url"`
	OrganizationsURL  string  `json:"organizations_url"`
	ReposURL          string  `json:"repos_url"`
	EventsURL         string  `json:"events_url"`
	ReceivedEventsURL string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	StarredAt         string  `json:"starred_at,omitempty"`
	UserViewType      string  `json:"user_view_type,omitempty"`
}

// Label represents a GitHub label
type Label struct {
	ID          int64   `json:"id"`
	NodeID      string  `json:"node_id"`
	URL         string  `json:"url"`
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Default     bool    `json:"default"`
	Description *string `json:"description"`
}

// SubIssuesSummary contains summary information about sub-issues
type SubIssuesSummary struct {
	Total            int `json:"total"`
	Completed        int `json:"completed"`
	PercentCompleted int `json:"percent_completed"`
}

// IssueDependenciesSummary contains summary information about issue dependencies
type IssueDependenciesSummary struct {
	BlockedBy      int `json:"blocked_by"`
	Blocking       int `json:"blocking"`
	TotalBlockedBy int `json:"total_blocked_by"`
	TotalBlocking  int `json:"total_blocking"`
}

// IssueFieldValue represents a value assigned to an issue field
type IssueFieldValue struct {
	IssueFieldID       int64                `json:"issue_field_id"`
	NodeID             string               `json:"node_id"`
	DataType           string               `json:"data_type"`
	Value              interface{}          `json:"value"`
	SingleSelectOption *SingleSelectOption  `json:"single_select_option,omitempty"`
}

// SingleSelectOption represents details about a selected option for single_select fields
type SingleSelectOption struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

// Milestone represents a GitHub milestone
type Milestone struct {
	URL          string      `json:"url"`
	HTMLURL      string      `json:"html_url"`
	LabelsURL    string      `json:"labels_url"`
	ID           int64       `json:"id"`
	NodeID       string      `json:"node_id"`
	Number       int         `json:"number"`
	State        string      `json:"state"`
	Title        string      `json:"title"`
	Description  *string     `json:"description"`
	Creator      *SimpleUser `json:"creator"`
	OpenIssues   int         `json:"open_issues"`
	ClosedIssues int         `json:"closed_issues"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	ClosedAt     *time.Time  `json:"closed_at"`
	DueOn        *time.Time  `json:"due_on"`
}

// TextMatch represents search result text matches
type TextMatch struct {
	ObjectURL  string        `json:"object_url"`
	ObjectType *string       `json:"object_type"`
	Property   string        `json:"property"`
	Fragment   string        `json:"fragment"`
	Matches    []TextMatchDetail `json:"matches"`
}

// TextMatchDetail represents the details of a text match
type TextMatchDetail struct {
	Text    string `json:"text"`
	Indices []int  `json:"indices"`
}

// PullRequestRef contains pull request reference information
type PullRequestRef struct {
	MergedAt *time.Time `json:"merged_at"`
	DiffURL  *string    `json:"diff_url"`
	HTMLURL  *string    `json:"html_url"`
	PatchURL *string    `json:"patch_url"`
	URL      *string    `json:"url"`
}

// Repository represents a GitHub repository
type Repository struct {
	ID                       int64           `json:"id"`
	NodeID                   string          `json:"node_id"`
	Name                     string          `json:"name"`
	FullName                 string          `json:"full_name"`
	License                  *LicenseSimple  `json:"license"`
	Forks                    int             `json:"forks"`
	Permissions              *Permissions    `json:"permissions,omitempty"`
	Owner                    SimpleUser      `json:"owner"`
	Private                  bool            `json:"private"`
	HTMLURL                  string          `json:"html_url"`
	Description              *string         `json:"description"`
	Fork                     bool            `json:"fork"`
	URL                      string          `json:"url"`
	ArchiveURL               string          `json:"archive_url"`
	AssigneesURL             string          `json:"assignees_url"`
	BlobsURL                 string          `json:"blobs_url"`
	BranchesURL              string          `json:"branches_url"`
	CollaboratorsURL         string          `json:"collaborators_url"`
	CommentsURL              string          `json:"comments_url"`
	CommitsURL               string          `json:"commits_url"`
	CompareURL               string          `json:"compare_url"`
	ContentsURL              string          `json:"contents_url"`
	ContributorsURL          string          `json:"contributors_url"`
	DeploymentsURL           string          `json:"deployments_url"`
	DownloadsURL             string          `json:"downloads_url"`
	EventsURL                string          `json:"events_url"`
	ForksURL                 string          `json:"forks_url"`
	GitCommitsURL            string          `json:"git_commits_url"`
	GitRefsURL               string          `json:"git_refs_url"`
	GitTagsURL               string          `json:"git_tags_url"`
	GitURL                   string          `json:"git_url"`
	IssueCommentURL          string          `json:"issue_comment_url"`
	IssueEventsURL           string          `json:"issue_events_url"`
	IssuesURL                string          `json:"issues_url"`
	KeysURL                  string          `json:"keys_url"`
	LabelsURL                string          `json:"labels_url"`
	LanguagesURL             string          `json:"languages_url"`
	MergesURL                string          `json:"merges_url"`
	MilestonesURL            string          `json:"milestones_url"`
	NotificationsURL         string          `json:"notifications_url"`
	PullsURL                 string          `json:"pulls_url"`
	ReleasesURL              string          `json:"releases_url"`
	SSHURL                   string          `json:"ssh_url"`
	StargazersURL            string          `json:"stargazers_url"`
	StatusesURL              string          `json:"statuses_url"`
	SubscribersURL           string          `json:"subscribers_url"`
	SubscriptionURL          string          `json:"subscription_url"`
	TagsURL                  string          `json:"tags_url"`
	TeamsURL                 string          `json:"teams_url"`
	TreesURL                 string          `json:"trees_url"`
	CloneURL                 string          `json:"clone_url"`
	MirrorURL                *string         `json:"mirror_url"`
	HooksURL                 string          `json:"hooks_url"`
	SvnURL                   string          `json:"svn_url"`
	Homepage                 *string         `json:"homepage"`
	Language                 *string         `json:"language"`
	ForksCount               int             `json:"forks_count"`
	StargazersCount          int             `json:"stargazers_count"`
	WatchersCount            int             `json:"watchers_count"`
	Size                     int             `json:"size"`
	DefaultBranch            string          `json:"default_branch"`
	OpenIssuesCount          int             `json:"open_issues_count"`
	IsTemplate               bool            `json:"is_template,omitempty"`
	Topics                   []string        `json:"topics,omitempty"`
	HasIssues                bool            `json:"has_issues"`
	HasProjects              bool            `json:"has_projects"`
	HasWiki                  bool            `json:"has_wiki"`
	HasPages                 bool            `json:"has_pages"`
	HasDownloads             bool            `json:"has_downloads"`
	HasDiscussions           bool            `json:"has_discussions,omitempty"`
	Archived                 bool            `json:"archived"`
	Disabled                 bool            `json:"disabled"`
	Visibility               string          `json:"visibility,omitempty"`
	PushedAt                 *time.Time      `json:"pushed_at"`
	CreatedAt                *time.Time      `json:"created_at"`
	UpdatedAt                *time.Time      `json:"updated_at"`
	AllowRebaseMerge         bool            `json:"allow_rebase_merge,omitempty"`
	TempCloneToken           string          `json:"temp_clone_token,omitempty"`
	AllowSquashMerge         bool            `json:"allow_squash_merge,omitempty"`
	AllowAutoMerge           bool            `json:"allow_auto_merge,omitempty"`
	DeleteBranchOnMerge      bool            `json:"delete_branch_on_merge,omitempty"`
	AllowUpdateBranch        bool            `json:"allow_update_branch,omitempty"`
	UseSquashPRTitleAsDefault bool           `json:"use_squash_pr_title_as_default,omitempty"`
	SquashMergeCommitTitle   string          `json:"squash_merge_commit_title,omitempty"`
	SquashMergeCommitMessage string          `json:"squash_merge_commit_message,omitempty"`
	MergeCommitTitle         string          `json:"merge_commit_title,omitempty"`
	MergeCommitMessage       string          `json:"merge_commit_message,omitempty"`
	AllowMergeCommit         bool            `json:"allow_merge_commit,omitempty"`
	AllowForking             bool            `json:"allow_forking,omitempty"`
	WebCommitSignoffRequired bool            `json:"web_commit_signoff_required,omitempty"`
	OpenIssues               int             `json:"open_issues"`
	Watchers                 int             `json:"watchers"`
	MasterBranch             string          `json:"master_branch,omitempty"`
	StarredAt                string          `json:"starred_at,omitempty"`
	AnonymousAccessEnabled   bool            `json:"anonymous_access_enabled,omitempty"`
	CodeSearchIndexStatus    *CodeSearchIndexStatus `json:"code_search_index_status,omitempty"`
}

// LicenseSimple represents a simple license
type LicenseSimple struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	URL     *string `json:"url"`
	SpdxID  *string `json:"spdx_id"`
	NodeID  string  `json:"node_id"`
	HTMLURL string  `json:"html_url,omitempty"`
}

// Permissions represents repository permissions
type Permissions struct {
	Admin    bool `json:"admin"`
	Pull     bool `json:"pull"`
	Triage   bool `json:"triage,omitempty"`
	Push     bool `json:"push"`
	Maintain bool `json:"maintain,omitempty"`
}

// CodeSearchIndexStatus represents the status of the code search index
type CodeSearchIndexStatus struct {
	LexicalSearchOK bool   `json:"lexical_search_ok"`
	LexicalCommitSHA string `json:"lexical_commit_sha"`
}

// IssueType represents the type of issue
type IssueType struct {
	ID          int       `json:"id"`
	NodeID      string    `json:"node_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Color       *string   `json:"color"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	IsEnabled   bool      `json:"is_enabled,omitempty"`
}

// GitHubApp represents a GitHub app
type GitHubApp struct {
	ID                 int64                  `json:"id"`
	Slug               string                 `json:"slug,omitempty"`
	NodeID             string                 `json:"node_id"`
	ClientID           string                 `json:"client_id,omitempty"`
	Owner              interface{}            `json:"owner"` // Can be SimpleUser or Enterprise
	Name               string                 `json:"name"`
	Description        *string                `json:"description"`
	ExternalURL        string                 `json:"external_url"`
	HTMLURL            string                 `json:"html_url"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
	Permissions        map[string]string      `json:"permissions"`
	Events             []string               `json:"events"`
	InstallationsCount int                    `json:"installations_count,omitempty"`
}

// ReactionRollup represents reaction counts
type ReactionRollup struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	PlusOne    int    `json:"+1"`
	MinusOne   int    `json:"-1"`
	Laugh      int    `json:"laugh"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Hooray     int    `json:"hooray"`
	Eyes       int    `json:"eyes"`
	Rocket     int    `json:"rocket"`
}

//-----------------//
// COMMIT SEARCHES //
//-----------------//


// CommitSearchResult represents the response from GitHub's commit search API
type CommitSearchResult struct {
	TotalCount        int                        `json:"total_count"`
	IncompleteResults bool                       `json:"incomplete_results"`
	Items             []CommitSearchResultItem   `json:"items"`
}

// CommitSearchResultItem represents a single commit in the search results
type CommitSearchResultItem struct {
	URL          string              `json:"url"`
	SHA          string              `json:"sha"`
	NodeID       string              `json:"node_id"`
	HTMLURL      string              `json:"html_url"`
	CommentsURL  string              `json:"comments_url"`
	Commit       CommitDetail        `json:"commit"`
	Author       *SimpleUser         `json:"author"`
	Committer    *GitUser            `json:"committer"`
	Parents      []CommitParent      `json:"parents"`
	Repository   MinimalRepository   `json:"repository"`
	Score        float64             `json:"score"`
	TextMatches  []TextMatch         `json:"text_matches,omitempty"`
}

// CommitDetail contains the detailed commit information
type CommitDetail struct {
	Author       GitUser       `json:"author"`
	Committer    *GitUser      `json:"committer"`
	CommentCount int           `json:"comment_count"`
	Message      string        `json:"message"`
	Tree         CommitTree    `json:"tree"`
	URL          string        `json:"url"`
	Verification *Verification `json:"verification,omitempty"`
}

// GitUser represents Git author/committer information
type GitUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

// CommitTree represents a git tree reference
type CommitTree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

// Verification represents commit verification information
type Verification struct {
	Verified   bool    `json:"verified"`
	Reason     string  `json:"reason"`
	Payload    *string `json:"payload"`
	Signature  *string `json:"signature"`
	VerifiedAt *string `json:"verified_at"`
}

// CommitParent represents a parent commit reference
type CommitParent struct {
	URL     string `json:"url,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
	SHA     string `json:"sha,omitempty"`
}

// MinimalRepository represents a minimal repository structure
type MinimalRepository struct {
	ID                      int64                   `json:"id"`
	NodeID                  string                  `json:"node_id"`
	Name                    string                  `json:"name"`
	FullName                string                  `json:"full_name"`
	Owner                   SimpleUser              `json:"owner"`
	Private                 bool                    `json:"private"`
	HTMLURL                 string                  `json:"html_url"`
	Description             *string                 `json:"description"`
	Fork                    bool                    `json:"fork"`
	URL                     string                  `json:"url"`
	ArchiveURL              string                  `json:"archive_url"`
	AssigneesURL            string                  `json:"assignees_url"`
	BlobsURL                string                  `json:"blobs_url"`
	BranchesURL             string                  `json:"branches_url"`
	CollaboratorsURL        string                  `json:"collaborators_url"`
	CommentsURL             string                  `json:"comments_url"`
	CommitsURL              string                  `json:"commits_url"`
	CompareURL              string                  `json:"compare_url"`
	ContentsURL             string                  `json:"contents_url"`
	ContributorsURL         string                  `json:"contributors_url"`
	DeploymentsURL          string                  `json:"deployments_url"`
	DownloadsURL            string                  `json:"downloads_url"`
	EventsURL               string                  `json:"events_url"`
	ForksURL                string                  `json:"forks_url"`
	GitCommitsURL           string                  `json:"git_commits_url"`
	GitRefsURL              string                  `json:"git_refs_url"`
	GitTagsURL              string                  `json:"git_tags_url"`
	GitURL                  string                  `json:"git_url,omitempty"`
	IssueCommentURL         string                  `json:"issue_comment_url"`
	IssueEventsURL          string                  `json:"issue_events_url"`
	IssuesURL               string                  `json:"issues_url"`
	KeysURL                 string                  `json:"keys_url"`
	LabelsURL               string                  `json:"labels_url"`
	LanguagesURL            string                  `json:"languages_url"`
	MergesURL               string                  `json:"merges_url"`
	MilestonesURL           string                  `json:"milestones_url"`
	NotificationsURL        string                  `json:"notifications_url"`
	PullsURL                string                  `json:"pulls_url"`
	ReleasesURL             string                  `json:"releases_url"`
	SSHURL                  string                  `json:"ssh_url,omitempty"`
	StargazersURL           string                  `json:"stargazers_url"`
	StatusesURL             string                  `json:"statuses_url"`
	SubscribersURL          string                  `json:"subscribers_url"`
	SubscriptionURL         string                  `json:"subscription_url"`
	TagsURL                 string                  `json:"tags_url"`
	TeamsURL                string                  `json:"teams_url"`
	TreesURL                string                  `json:"trees_url"`
	CloneURL                string                  `json:"clone_url,omitempty"`
	MirrorURL               *string                 `json:"mirror_url"`
	HooksURL                string                  `json:"hooks_url"`
	SvnURL                  string                  `json:"svn_url,omitempty"`
	Homepage                *string                 `json:"homepage"`
	Language                *string                 `json:"language"`
	ForksCount              int                     `json:"forks_count,omitempty"`
	StargazersCount         int                     `json:"stargazers_count,omitempty"`
	WatchersCount           int                     `json:"watchers_count,omitempty"`
	Size                    int                     `json:"size,omitempty"`
	DefaultBranch           string                  `json:"default_branch,omitempty"`
	OpenIssuesCount         int                     `json:"open_issues_count,omitempty"`
	IsTemplate              bool                    `json:"is_template,omitempty"`
	Topics                  []string                `json:"topics,omitempty"`
	HasIssues               bool                    `json:"has_issues,omitempty"`
	HasProjects             bool                    `json:"has_projects,omitempty"`
	HasWiki                 bool                    `json:"has_wiki,omitempty"`
	HasPages                bool                    `json:"has_pages,omitempty"`
	HasDownloads            bool                    `json:"has_downloads,omitempty"`
	HasDiscussions          bool                    `json:"has_discussions,omitempty"`
	Archived                bool                    `json:"archived,omitempty"`
	Disabled                bool                    `json:"disabled,omitempty"`
	Visibility              string                  `json:"visibility,omitempty"`
	PushedAt                *time.Time              `json:"pushed_at"`
	CreatedAt               *time.Time              `json:"created_at"`
	UpdatedAt               *time.Time              `json:"updated_at"`
	Permissions             *Permissions            `json:"permissions,omitempty"`
	RoleName                string                  `json:"role_name,omitempty"`
	TempCloneToken          string                  `json:"temp_clone_token,omitempty"`
	DeleteBranchOnMerge     bool                    `json:"delete_branch_on_merge,omitempty"`
	SubscribersCount        int                     `json:"subscribers_count,omitempty"`
	NetworkCount            int                     `json:"network_count,omitempty"`
	CodeOfConduct           *CodeOfConduct          `json:"code_of_conduct,omitempty"`
	License                 *RepositoryLicense      `json:"license"`
	Forks                   int                     `json:"forks,omitempty"`
	OpenIssues              int                     `json:"open_issues,omitempty"`
	Watchers                int                     `json:"watchers,omitempty"`
	AllowForking            bool                    `json:"allow_forking,omitempty"`
	WebCommitSignoffRequired bool                   `json:"web_commit_signoff_required,omitempty"`
	SecurityAndAnalysis     *SecurityAndAnalysis    `json:"security_and_analysis"`
	CustomProperties        map[string]interface{}  `json:"custom_properties,omitempty"`
}

// CodeOfConduct represents a code of conduct
type CodeOfConduct struct {
	Key     string  `json:"key"`
	Name    string  `json:"name"`
	URL     string  `json:"url"`
	Body    string  `json:"body,omitempty"`
	HTMLURL *string `json:"html_url"`
}

// RepositoryLicense represents a repository license
type RepositoryLicense struct {
	Key    string  `json:"key"`
	Name   string  `json:"name"`
	SpdxID string  `json:"spdx_id"`
	URL    string  `json:"url"`
	NodeID string  `json:"node_id"`
}

// SecurityAndAnalysis represents security and analysis settings
type SecurityAndAnalysis struct {
	AdvancedSecurity                 *SecurityFeature `json:"advanced_security,omitempty"`
	CodeSecurity                     *SecurityFeature `json:"code_security,omitempty"`
	DependabotSecurityUpdates        *SecurityFeature `json:"dependabot_security_updates,omitempty"`
	SecretScanning                   *SecurityFeature `json:"secret_scanning,omitempty"`
	SecretScanningPushProtection     *SecurityFeature `json:"secret_scanning_push_protection,omitempty"`
	SecretScanningNonProviderPatterns *SecurityFeature `json:"secret_scanning_non_provider_patterns,omitempty"`
	SecretScanningAIDetection        *SecurityFeature `json:"secret_scanning_ai_detection,omitempty"`
}

// SecurityFeature represents a security feature status
type SecurityFeature struct {
	Status string `json:"status"`
}