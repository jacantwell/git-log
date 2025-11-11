package config

import (
	"fmt"
	"os"
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

	googleToken := os.Getenv("GOOGLE_API_KEY")
	if googleToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	username := os.Getenv("GITHUB_USERNAME")
	if username == "" {
		return nil, fmt.Errorf("GITHUB_USERNAME environment variable not set")
	}

	return &Config{
		GoogleToken:      googleToken,
		GitHubToken:      githubToken,
		Username:         username,
		Days:             3,
		ReportPath:       "report.md",
		SystemPromptPath: "internal/report/system_prompt.md",
		Model:            "gemini-2.5-flash",
	}, nil
}
