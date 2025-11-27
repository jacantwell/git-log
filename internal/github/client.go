package github

import (
    "fmt"
    "io"
    "context"
    "net/http"
    "time"
)

type Client struct {
    Token      string
    HTTPClient *http.Client
    BaseURL    string
}

func NewClient(token string) *Client {
    return &Client{
        Token:      token,
        HTTPClient: &http.Client{Timeout: 10 * time.Second},
        BaseURL:    "https://api.github.com",
    }
}

func (c *Client) makeRequest(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token))
    req.Header.Set("Accept", "application/vnd.github.v3+json")

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("API request failed with status: %d, body: %s", 
            resp.StatusCode, string(body))
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return body, nil
}