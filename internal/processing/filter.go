package processing

import (
	"git-log/internal/github"
	"time"
)

// FilterPullRequests extracts essential information from GitHub PR search results
func FilterPullRequests(items []github.IssueSearchResultItem) []PullRequest {
	filtered := make([]PullRequest, 0, len(items))

	for _, item := range items {
		pr := PullRequest{
			Number:    item.Number,
			Title:     item.Title,
			Body:      item.Body,
			State:     item.State,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			ClosedAt:  item.ClosedAt,
			URL:       item.HTMLURL,
			Comments:  item.Comments,
			IsDraft:   item.Draft,
		}

		// Extract merged_at from pull_request reference if available
		if item.PullRequest != nil && item.PullRequest.MergedAt != nil {
			pr.MergedAt = item.PullRequest.MergedAt
		}

		// Extract label names
		pr.Labels = make([]string, 0, len(item.Labels))
		for _, label := range item.Labels {
			pr.Labels = append(pr.Labels, label.Name)
		}

		filtered = append(filtered, pr)
	}

	return filtered
}

// FilterCommits extracts essential information from GitHub commit search results
func FilterCommits(items []github.CommitSearchResultItem) []Commit {
	filtered := make([]Commit, 0, len(items))

	for _, item := range items {
		commit := Commit{
			SHA:     item.SHA,
			Message: item.Commit.Message,
			URL:     item.HTMLURL,
		}

		// Parse the date from the commit author
		if item.Commit.Author.Date != "" {
			// GitHub returns dates in RFC3339 format
			parsedDate, err := time.Parse(time.RFC3339, item.Commit.Author.Date)
			if err == nil {
				commit.Date = parsedDate
			}
		}

		filtered = append(filtered, commit)
	}

	return filtered
}

// ExtractRepositoryInfo extracts essential repository information
func ExtractRepositoryInfo(repo github.Repository) (name, fullName, description, url, language string) {
	name = repo.Name
	fullName = repo.FullName
	url = repo.HTMLURL

	if repo.Description != nil {
		description = *repo.Description
	}

	if repo.Language != nil {
		language = *repo.Language
	}

	return
}
