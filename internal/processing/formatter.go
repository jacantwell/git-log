package processing

import (
	"fmt"
	"strings"
	"time"
)

// FormatAsMarkdown converts a WorkLog to a formatted markdown document
func FormatAsMarkdown(workLog *WorkLog) string {
	var sb strings.Builder
	
	// Title and summary
	sb.WriteString("# GitHub Activity Work Log\n\n")
	sb.WriteString(fmt.Sprintf("**Period:** %s to %s\n\n",
		workLog.Summary.DateRange.Start.Format("January 2, 2006"),
		workLog.Summary.DateRange.End.Format("January 2, 2006")))
	
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- **Repositories:** %d\n", workLog.Summary.TotalRepositories))
	sb.WriteString(fmt.Sprintf("- **Pull Requests:** %d\n", workLog.Summary.TotalPullRequests))
	sb.WriteString(fmt.Sprintf("- **Commits:** %d\n\n", workLog.Summary.TotalCommits))
	
	// Activity by repository
	sb.WriteString("## Activity by Repository\n\n")
	
	for _, repo := range workLog.Repositories {
		sb.WriteString(formatRepository(repo))
	}
	
	return sb.String()
}

func formatRepository(repo RepositoryActivity) string {
	var sb strings.Builder
	
	// Repository header
	sb.WriteString(fmt.Sprintf("### %s\n\n", repo.FullName))
	
	if repo.Description != "" {
		sb.WriteString(fmt.Sprintf("*%s*\n\n", repo.Description))
	}
	
	if repo.Language != "" {
		sb.WriteString(fmt.Sprintf("**Language:** %s\n\n", repo.Language))
	}
	
	sb.WriteString(fmt.Sprintf("[View Repository](%s)\n\n", repo.URL))
	
	// Pull requests section
	if len(repo.PullRequests) > 0 {
		sb.WriteString(fmt.Sprintf("#### Pull Requests (%d)\n\n", len(repo.PullRequests)))
		
		for _, pr := range repo.PullRequests {
			sb.WriteString(formatPullRequest(pr))
		}
	}
	
	// Commits section
	if len(repo.Commits) > 0 {
		sb.WriteString(fmt.Sprintf("#### Commits (%d)\n\n", len(repo.Commits)))
		
		for _, commit := range repo.Commits {
			sb.WriteString(formatCommit(commit))
		}
	}
	
	sb.WriteString("\n")
	return sb.String()
}

func formatPullRequest(pr PullRequest) string {
	var sb strings.Builder
	
	// PR title with link and status
	status := getStatusEmoji(pr.State, pr.MergedAt)
	sb.WriteString(fmt.Sprintf("**%s [#%d: %s](%s)**\n\n",
		status, pr.Number, pr.Title, pr.URL))
	
	// Date information
	dateInfo := fmt.Sprintf("*Created: %s", pr.CreatedAt.Format("Jan 2, 2006"))
	if pr.MergedAt != nil {
		dateInfo += fmt.Sprintf(" | Merged: %s", pr.MergedAt.Format("Jan 2, 2006"))
	} else if pr.ClosedAt != nil {
		dateInfo += fmt.Sprintf(" | Closed: %s", pr.ClosedAt.Format("Jan 2, 2006"))
	}
	dateInfo += "*\n\n"
	sb.WriteString(dateInfo)
	
	// Labels
	if len(pr.Labels) > 0 {
		sb.WriteString("**Labels:** ")
		for i, label := range pr.Labels {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(fmt.Sprintf("`%s`", label))
		}
		sb.WriteString("\n\n")
	}
	
	// PR body/description
	if pr.Body != "" {
		// Truncate long descriptions
		body := pr.Body
		if len(body) > 300 {
			body = body[:300] + "..."
		}
		sb.WriteString(fmt.Sprintf("%s\n\n", body))
	}
	
	// Additional info
	if pr.Comments > 0 {
		sb.WriteString(fmt.Sprintf("*%d comments*\n\n", pr.Comments))
	}
	
	sb.WriteString("---\n\n")
	return sb.String()
}

func formatCommit(commit Commit) string {
	var sb strings.Builder
	
	// Commit message with link
	verified := ""
	if commit.Verified {
		verified = " ✓"
	}
	
	// Get first line of commit message
	message := strings.Split(commit.Message, "\n")[0]
	if len(message) > 80 {
		message = message[:80] + "..."
	}
	
	sb.WriteString(fmt.Sprintf("- **[`%s`](%s)%s** %s",
		commit.SHA[:7], commit.URL, verified, message))
	
	if !commit.Date.IsZero() {
		sb.WriteString(fmt.Sprintf(" *(%s)*", commit.Date.Format("Jan 2, 2006")))
	}
	
	sb.WriteString("\n")
	return sb.String()
}

func getStatusEmoji(state string, mergedAt *time.Time) string {
	if mergedAt != nil {
		return "✅" // Merged
	}
	switch state {
	case "open":
		return "🔵" // Open
	case "closed":
		return "🔴" // Closed without merging
	default:
		return "⚪" // Unknown
	}
}

// FormatAsSummary creates a brief summary suitable for a CV
func FormatAsSummary(workLog *WorkLog) string {
	var sb strings.Builder
	
	sb.WriteString("# GitHub Activity Summary\n\n")
	sb.WriteString(fmt.Sprintf("**Period:** %s to %s\n\n",
		workLog.Summary.DateRange.Start.Format("Jan 2006"),
		workLog.Summary.DateRange.End.Format("Jan 2006")))
	
	sb.WriteString("## Key Contributions\n\n")
	
	// Highlight repositories with most activity
	type repoStats struct {
		name     string
		prs      int
		commits  int
		language string
	}
	
	stats := make([]repoStats, 0, len(workLog.Repositories))
	for _, repo := range workLog.Repositories {
		stats = append(stats, repoStats{
			name:     repo.FullName,
			prs:      len(repo.PullRequests),
			commits:  len(repo.Commits),
			language: repo.Language,
		})
	}
	
	for _, stat := range stats {
		activity := []string{}
		if stat.prs > 0 {
			activity = append(activity, fmt.Sprintf("%d pull requests", stat.prs))
		}
		if stat.commits > 0 {
			activity = append(activity, fmt.Sprintf("%d commits", stat.commits))
		}
		
		langInfo := ""
		if stat.language != "" {
			langInfo = fmt.Sprintf(" (%s)", stat.language)
		}
		
		sb.WriteString(fmt.Sprintf("- **%s**%s: %s\n",
			stat.name, langInfo, strings.Join(activity, ", ")))
	}
	
	sb.WriteString(fmt.Sprintf("\n**Total Activity:** %d pull requests and %d commits across %d repositories\n",
		workLog.Summary.TotalPullRequests,
		workLog.Summary.TotalCommits,
		workLog.Summary.TotalRepositories))
	
	return sb.String()
}