package main

import (
	"fmt"
	"os"
	"time"

	"git-log/config"
	"git-log/internal/github"
	"git-log/internal/processing"
	"git-log/internal/report"
)

func main() {

	config, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// Convert int days to time.Time
	since := time.Now().AddDate(0, 0, -config.Days)

	fmt.Println("Fetching GitHub activity...")

	var client = github.NewClient(config.GitHubToken)

	commits, err := client.GetCommits(config.Username, since)
	if err != nil {
		fmt.Printf("Warning: Failed to fetch commits: %v\n", err)
		fmt.Println("Continuing with pull requests only...")
		commits = []github.CommitSearchResultItem{}
	}

	pullRequests, err := client.GetPullRequests(config.Username, since)
	if err != nil {
		fmt.Printf("Error getting pull requests: %v\n", err)
		return
	}

	fmt.Printf("Found %d pull requests and %d commits\n", len(pullRequests), len(commits))

	// Process and group data
	fmt.Println("Processing activity data...")
	workLog := processing.GroupByRepository(pullRequests, commits)

	// Display summary
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Repositories: %d\n", workLog.Summary.TotalRepositories)
	fmt.Printf("Pull Requests: %d\n", workLog.Summary.TotalPullRequests)
	fmt.Printf("Commits: %d\n", workLog.Summary.TotalCommits)
	fmt.Printf("Period: %s to %s\n",
		workLog.Summary.DateRange.Start.Format("Jan 2, 2006"),
		workLog.Summary.DateRange.End.Format("Jan 2, 2006"))

	// Analyse and generate report
	fmt.Println("Generating accomplishment report...")
	result, err := report.GenerateReport(config.Model, *workLog, config.SystemPromptPath, config.ReportPath)

	// Save report to file
	if err == nil {
		err = os.WriteFile(config.ReportPath, []byte(result), 0644)
		if err != nil {
			fmt.Printf("Error writing report to file: %v\n", err)
			return
		}
		fmt.Printf("Report saved to %s\n", config.ReportPath)
	}

	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		return
	}
}
