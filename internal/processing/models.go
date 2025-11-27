package processing

import "time"

// WorkLog represents a collection of activities grouped by repository
type WorkLog struct {
	Repositories []RepositoryActivity `json:"repositories"`
	Summary      Summary              `json:"summary"`
}

// RepositoryActivity contains all activity for a single repository
type RepositoryActivity struct {
	Name         string        `json:"name"`
	FullName     string        `json:"full_name"`
	Description  string        `json:"description,omitempty"`
	URL          string        `json:"url"`
	PullRequests []PullRequest `json:"pull_requests,omitempty"`
	Commits      []Commit      `json:"commits,omitempty"`
	Language     string        `json:"language,omitempty"`
}

// PullRequest represents essential PR information
type PullRequest struct {
	Number    int        `json:"number"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	State     string     `json:"state"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	MergedAt  *time.Time `json:"merged_at,omitempty"`
	URL       string     `json:"url"`
	Comments  int        `json:"comments"`
	Labels    []string   `json:"labels,omitempty"`
	IsDraft   bool       `json:"is_draft,omitempty"`
}

// Commit represents essential commit information
type Commit struct {
	SHA     string    `json:"sha"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	URL     string    `json:"url"`
}

// Summary provides high-level statistics
type Summary struct {
	TotalRepositories int       `json:"total_repositories"`
	TotalPullRequests int       `json:"total_pull_requests"`
	TotalCommits      int       `json:"total_commits"`
	DateRange         DateRange `json:"date_range"`
}

// DateRange represents the time period covered
type DateRange struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
