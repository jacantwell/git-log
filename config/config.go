package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	GoogleToken      string
	GitHubToken      string
	Username         string
	LookbackDays             int
	ReportPath       string
	Model            string
}

func Load() (*Config, error) {

	githubToken := os.Getenv("ACCESS_TOKEN")
	if githubToken == "" {
		return nil, fmt.Errorf("ACCESS_TOKEN environment variable not set")
	}
	
	username := os.Getenv("USERNAME")
	if username == "" {
		return nil, fmt.Errorf("USERNAME environment variable not set")
	}
	
	googleToken := os.Getenv("GOOGLE_API_KEY")
	if googleToken == "" {
		return nil, fmt.Errorf("GOOGLE_API_KEY environment variable not set")
	}

	reportPath := os.Getenv("REPORT_PATH")
	if reportPath == "" {
		reportPath = "report.md"
	}
	
	// Make report path absolute if it's relative and we're in GitHub Actions
	if !filepath.IsAbs(reportPath) {
		workspace := os.Getenv("GITHUB_WORKSPACE")
		if workspace != "" {
			reportPath = filepath.Join(workspace, reportPath)
		}
	}

	model := os.Getenv("MODEL")
	if model == "" {
		model = "gemini-2.5-flash"
	}

	days := (os.Getenv("LOOKBACK_DAYS"))
	daysInt, err := strconv.Atoi(days)
		if err != nil {
			return nil, fmt.Errorf("invalid DAYS value: %v", err)
		}


	return &Config{
		GoogleToken:      googleToken,
		GitHubToken:      githubToken,
		Username:         username,
		LookbackDays:     daysInt,
		ReportPath:       reportPath,
		Model:            model,
	}, nil
}
