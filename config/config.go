package config

import (
	"fmt"
	"os"
)

type Config struct {
	GoogleToken      string
	ResendToken      string
	GitHubToken      string
	Username         string
	emailAddress     string
	Days             int
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

	emailAddress := os.Getenv("EMAIL_ADDRESS")
	// Email address is optional; no error if not set

	googleToken := os.Getenv("GOOGLE_API_KEY")
	if googleToken == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}

	resendToken := os.Getenv("RESEND_API_KEY")
	// Resend API key is optional; no error if not set

	return &Config{
		GoogleToken:      googleToken,
		ResendToken:      resendToken,
		GitHubToken:      githubToken,
		Username:         username,
		emailAddress:     emailAddress,
		Days:             30,
		ReportPath:       "report.md",
		SystemPromptPath: "internal/report/system_prompt.md",
		Model:            "gemini-2.5-flash",
	}, nil
}
