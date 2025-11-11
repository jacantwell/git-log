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
	URL               string          `json:"url"`
	RepositoryURL     string          `json:"repository_url"`
	LabelsURL         string          `json:"labels_url"`
	CommentsURL       string          `json:"comments_url"`
	EventsURL         string          `json:"events_url"`
	HTMLURL           string          `json:"html_url"`
	ID                int64           `json:"id"`
	NodeID            string          `json:"node_id"`
	Number            int             `json:"number"`
	Title             string          `json:"title"`
	User              *SimpleUser     `json:"user"`
	Labels            []Label         `json:"labels"`
	State             string          `json:"state"`
	Draft             bool            `json:"draft,omitempty"`
	Comments          int             `json:"comments"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	ClosedAt          *time.Time      `json:"closed_at"`
	PullRequest       *PullRequestRef `json:"pull_request,omitempty"`
	Body              string          `json:"body"`
	Score             float64         `json:"score"`
	AuthorAssociation string          `json:"author_association"`
	Repository        Repository      `json:"repository"`
	BodyText          string          `json:"body_text,omitempty"`
	Type              *IssueType      `json:"type"`
}

// SimpleUser represents a GitHub user
type SimpleUser struct {
	Name   *string `json:"name"`
	Email  *string `json:"email"`
	Login  string  `json:"login"`
	ID     int64   `json:"id"`
	NodeID string  `json:"node_id"`
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
	ID          int64      `json:"id"`
	NodeID      string     `json:"node_id"`
	Name        string     `json:"name"`
	FullName    string     `json:"full_name"`
	Owner       SimpleUser `json:"owner"`
	Private     bool       `json:"private"`
	HTMLURL     string     `json:"html_url"`
	Description *string    `json:"description"`
	URL         string     `json:"url"`
	Homepage    *string    `json:"homepage"`
	Language    *string    `json:"language"`
	Topics      []string   `json:"topics,omitempty"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
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

//-----------------//
// COMMIT SEARCHES //
//-----------------//

// CommitSearchResult represents the response from GitHub's commit search API
type CommitSearchResult struct {
	TotalCount        int                      `json:"total_count"`
	IncompleteResults bool                     `json:"incomplete_results"`
	Items             []CommitSearchResultItem `json:"items"`
}

// CommitSearchResultItem represents a single commit in the search results
type CommitSearchResultItem struct {
	URL         string       `json:"url"`
	SHA         string       `json:"sha"`
	NodeID      string       `json:"node_id"`
	HTMLURL     string       `json:"html_url"`
	CommentsURL string       `json:"comments_url"`
	Commit      CommitDetail `json:"commit"`
	Author      *SimpleUser  `json:"author"`
	Committer   *GitUser     `json:"committer"`
	Repository  Repository   `json:"repository"`
}

// CommitDetail contains the detailed commit information
type CommitDetail struct {
	Author    GitUser  `json:"author"`
	Committer *GitUser `json:"committer"`
	Message   string   `json:"message"`
	URL       string   `json:"url"`
}

// GitUser represents Git author/committer information
type GitUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}
