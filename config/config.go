package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	GoogleToken      string
	GitHubToken      string
	Username         string
	Days             int
	ReportPath       string
	SystemPromptPath string
	Model            string
}

func Load() (*Config, error) {

	githubToken := os.Getenv("ACCESS_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("ACCESS_TOKEN environment variable not set")
	}
	
	username := os.Getenv("USERNAME")
	if username == "" {
		return nil, fmt.Errorf("GITHUB_USERNAME environment variable not set")
	}
	
	googleToken := os.Getenv("GOOGLE_API_KEY")
	if googleToken == "" {
		return nil, fmt.Errorf("GOOGLE_API_KEY environment variable not set")
	}

	reportPath := os.Getenv("REPORT_PATH")
	if reportPath == "" {
		reportPath = "report.md"
	}

	model := os.Getenv("MODEL")
	if model == "" {
		model = "gemini-2.5-flash"
	}

	days := (os.Getenv("DAYS"))
	if days == "" {
		days = "30"
	}

	// Convert Days to int
	daysInt, err := strconv.Atoi(days)
	if err != nil {
		return nil, fmt.Errorf("invalid DAYS value: %v", err)
	}



	return &Config{
		GoogleToken:      googleToken,
		GitHubToken:      githubToken,
		Username:         username,
		Days:             daysInt,
		ReportPath:       reportPath,
		SystemPromptPath: "internal/report/system_prompt.md",
		Model:            model,
	}, nil
}
