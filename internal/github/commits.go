package github

import (
    "encoding/json"
    "context"
    "fmt"
    "net/url"
    "time"
)

func (c *Client) GetCommits(ctx context.Context, author string, since time.Time) ([]CommitSearchResultItem, error) {
    query := fmt.Sprintf("author:%s author-date:>%s", author, since.Format(time.RFC3339))

    // Build URL with properly encoded query parameters
    baseURL := fmt.Sprintf("%s/search/commits", c.BaseURL)
    params := url.Values{}
    params.Add("q", query)
    params.Add("per_page", "100")
    params.Add("sort", "author-date")
    params.Add("order", "desc")

    requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    body, err := c.makeRequest(ctx, requestURL)
    if err != nil {
        return nil, err
    }

    var response CommitSearchResult
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    return response.Items, nil
}
