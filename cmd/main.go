package main

import (
    "encoding/json"
    "fmt"
	"os"
	"time"

	"git-log/internal/github"
	"git-log/config"
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

    // commits, err := client.GetCommits(config.Username, since)
    // if err != nil {
    //     fmt.Printf("Error getting commits: %v\n", err)
    //     return
    // }

	pullRequests, err := client.GetPullRequests(config.Username, since)
	if err != nil {
		fmt.Printf("Error getting pull requests: %v\n", err)
		return
	}


    // // Output summary
    // fmt.Printf("\nFound activity in %d repositories\n", len(commits))
    // for _, commit := range commits {
	// 	fmt.Printf("\n%s:\n", commit.Repository.FullName)
    //     fmt.Printf("  Message: %s\n", commit.Commit.Message)
    // }

    // // Save to file
    // output, err := json.MarshalIndent(commits, "", "  ")
    // if err != nil {
    //     fmt.Printf("Error marshaling JSON: %v\n", err)
    //     return
    // }

    // err = os.WriteFile(config.OutputPath, output, 0644)
    // if err != nil {
    //     fmt.Printf("Error writing file: %v\n", err)
    //     return
    // }

    // Output summary
    fmt.Printf("\nFound activity in %d repositories\n", len(pullRequests))
    for _, pr := range pullRequests {
		fmt.Printf("\n%s:\n", pr.Repository.FullName)
        fmt.Printf("  Title: %s\n", pr.Title)
		fmt.Printf("  Body: %s\n", pr.Body)
    }

    // Save to file
    output, err := json.MarshalIndent(pullRequests, "", "  ")
    if err != nil {
        fmt.Printf("Error marshaling JSON: %v\n", err)
        return
    }

    err = os.WriteFile(config.OutputPath, output, 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }

    fmt.Println("\nActivity saved to github_activity.json")
}