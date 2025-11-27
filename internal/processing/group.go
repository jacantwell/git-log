package processing

import (
	"git-log/internal/github"
	"sort"
	"strings"
	"time"
)

// GroupByRepository organizes pull requests and commits by their repositories
func GroupByRepository(prs []github.IssueSearchResultItem, commits []github.CommitSearchResultItem) *WorkLog {
	repoMap := make(map[string]*RepositoryActivity)

	// Process pull requests
	for _, pr := range prs {
		repoFullName := pr.Repository.FullName

		// If repository info is empty, try to extract from PR URL
		if repoFullName == "" && pr.HTMLURL != "" {
			repoFullName = extractRepoFromURL(pr.HTMLURL)
		}

		// Skip if we still can't determine the repository
		if repoFullName == "" {
			continue
		}

		// Initialize repository if not exists
		if _, exists := repoMap[repoFullName]; !exists {
			name, fullName, description, url, language := ExtractRepositoryInfo(pr.Repository)

			// If extraction failed, use info from URL
			if fullName == "" {
				fullName = repoFullName
				name = extractRepoNameFromFullName(repoFullName)
				url = "https://github.com/" + repoFullName
			}

			repoMap[repoFullName] = &RepositoryActivity{
				Name:         name,
				FullName:     fullName,
				Description:  description,
				URL:          url,
				Language:     language,
				PullRequests: []PullRequest{},
				Commits:      []Commit{},
			}
		}

		// Add filtered PR to repository
		filteredPRs := FilterPullRequests([]github.IssueSearchResultItem{pr})
		if len(filteredPRs) > 0 {
			repoMap[repoFullName].PullRequests = append(repoMap[repoFullName].PullRequests, filteredPRs[0])
		}
	}

	// Process commits
	for _, commit := range commits {
		repoFullName := commit.Repository.FullName

		// Initialize repository if not exists
		if _, exists := repoMap[repoFullName]; !exists {
			// Convert MinimalRepository to Repository for extraction
			repo := github.Repository{
				Name:        commit.Repository.Name,
				FullName:    commit.Repository.FullName,
				Description: commit.Repository.Description,
				HTMLURL:     commit.Repository.HTMLURL,
				Language:    commit.Repository.Language,
			}
			name, fullName, description, url, language := ExtractRepositoryInfo(repo)
			repoMap[repoFullName] = &RepositoryActivity{
				Name:         name,
				FullName:     fullName,
				Description:  description,
				URL:          url,
				Language:     language,
				PullRequests: []PullRequest{},
				Commits:      []Commit{},
			}
		}

		// Add filtered commit to repository
		filteredCommits := FilterCommits([]github.CommitSearchResultItem{commit})
		if len(filteredCommits) > 0 {
			repoMap[repoFullName].Commits = append(repoMap[repoFullName].Commits, filteredCommits[0])
		}
	}

	// Convert map to slice and sort by repository name
	repositories := make([]RepositoryActivity, 0, len(repoMap))
	for _, repo := range repoMap {
		// Sort PRs by created date (newest first)
		sort.Slice(repo.PullRequests, func(i, j int) bool {
			return repo.PullRequests[i].CreatedAt.After(repo.PullRequests[j].CreatedAt)
		})

		// Sort commits by date (newest first)
		sort.Slice(repo.Commits, func(i, j int) bool {
			return repo.Commits[i].Date.After(repo.Commits[j].Date)
		})

		repositories = append(repositories, *repo)
	}

	// Sort repositories alphabetically
	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].FullName < repositories[j].FullName
	})

	// Generate summary
	summary := generateSummary(repositories)

	return &WorkLog{
		Repositories: repositories,
		Summary:      summary,
	}
}

// generateSummary creates summary statistics for the work log
func generateSummary(repos []RepositoryActivity) Summary {
	summary := Summary{
		TotalRepositories: len(repos),
	}

	var earliestDate, latestDate time.Time
	firstDate := true

	for _, repo := range repos {
		summary.TotalPullRequests += len(repo.PullRequests)
		summary.TotalCommits += len(repo.Commits)

		// Track date range from PRs
		for _, pr := range repo.PullRequests {
			if firstDate || pr.CreatedAt.Before(earliestDate) {
				earliestDate = pr.CreatedAt
			}
			if firstDate || pr.CreatedAt.After(latestDate) {
				latestDate = pr.CreatedAt
				firstDate = false
			}
		}

		// Track date range from commits
		for _, commit := range repo.Commits {
			if firstDate || commit.Date.Before(earliestDate) {
				earliestDate = commit.Date
			}
			if firstDate || commit.Date.After(latestDate) {
				latestDate = commit.Date
				firstDate = false
			}
		}
	}

	summary.DateRange = DateRange{
		Start: earliestDate,
		End:   latestDate,
	}

	return summary
}

// extractRepoFromURL extracts the repository full name from a GitHub URL
// Example: "https://github.com/jacantwell/git-log/pull/1" -> "jacantwell/git-log"
func extractRepoFromURL(url string) string {
	// Remove protocol and domain
	parts := strings.Split(url, "github.com/")
	if len(parts) < 2 {
		return ""
	}

	// Split path and take first two parts (owner/repo)
	pathParts := strings.Split(parts[1], "/")
	if len(pathParts) < 2 {
		return ""
	}

	return pathParts[0] + "/" + pathParts[1]
}

// extractRepoNameFromFullName extracts just the repo name from full name
// Example: "jacantwell/git-log" -> "git-log"
func extractRepoNameFromFullName(fullName string) string {
	parts := strings.Split(fullName, "/")
	if len(parts) >= 2 {
		return parts[1]
	}
	return fullName
}
