package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"git-log/config"
	"git-log/internal/github"
	"git-log/internal/processing"
	"git-log/internal/report"
)

func main() {
	// Load configuration
	config, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// Convert int days to time.Time
	since := time.Now().AddDate(0, 0, -config.Days)

	var client = github.NewClient(config.GitHubToken)

	fmt.Println("Fetching GitHub activity...")

	// Fetch commits
	commits, err := client.GetCommits(config.Username, since)
	if err != nil {
		fmt.Printf("Warning: Error getting commits: %v\n", err)
		commits = []github.CommitSearchResultItem{}
	}

	// Fetch pull requests
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

	// Save processed JSON
	fmt.Println("\nSaving processed data...")
	processedJSON, err := json.MarshalIndent(workLog, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling processed JSON: %v\n", err)
		return
	}

	err = os.WriteFile("work_log.json", processedJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing work_log.json: %v\n", err)
		return
	}
	fmt.Println("Processed data saved to work_log.json")

	// Save raw data for reference
	rawJSON, err := json.MarshalIndent(pullRequests, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling raw JSON: %v\n", err)
		return
	}

	err = os.WriteFile(config.OutputPath, rawJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing raw data: %v\n", err)
		return
	}

	// Analyse and generate report
	fmt.Println("Generating accomplishment report...")
	err = report.GenerateReport(config.Model)

	fmt.Println("\n✓ All files generated successfully!")
	fmt.Println("\nGenerated files:")
	fmt.Println("  - work_log.json (processed data)")
	fmt.Println("  - github_activity.json (raw data)")
	fmt.Println("  - report.md (accomplishment report)")

	if err != nil {
		fmt.Printf("Error generating report: %v\n", err)
		return
	}
}