package config

import (
    "fmt"
    "os"
)

type Config struct {
    GitHubToken   string
    Username      string
    Days          int
    // Strategy      collector.CollectionStrategy
    // Organizations []string
    OutputPath    string
}

func Load() (*Config, error) {
    // For now, load from environment variables
    // Later you can add YAML/JSON config file support
    
    token := os.Getenv("GITHUB_TOKEN")
    if token == "" {
        return nil, fmt.Errorf("GITHUB_TOKEN environment variable not set")
    }

    username := os.Getenv("GITHUB_USERNAME")
    if username == "" {
        return nil, fmt.Errorf("GITHUB_USERNAME environment variable not set")
    }

    return &Config{
        GitHubToken:   token,
        Username:      username,
        Days:          90,
        OutputPath:    "github_activity.json",
    }, nil
}