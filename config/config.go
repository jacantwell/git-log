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
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	return &Config{
		GoogleToken:      googleToken,
		GitHubToken:      githubToken,
		Username:         username,
		Days:             30,
		ReportPath:       "report.md",
		SystemPromptPath: "internal/report/system_prompt.md",
		Model:            "gemini-2.5-flash",
	}, nil
}
