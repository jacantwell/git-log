package github

import (
    "encoding/json"
    "net/url"
    "fmt"
    "time"
)

func (c *Client) GetPullRequests(author string, since time.Time) ([]IssueSearchResultItem, error) {
    query := fmt.Sprintf("type:pr author:%s author-date:>%s", author, since.Format(time.RFC3339))
    // query := fmt.Sprintf("type:pr author-date:>%s", since.Format(time.RFC3339))

    // Build URL with properly encoded query parameters
    baseURL := fmt.Sprintf("%s/search/issues", c.BaseURL)
    params := url.Values{}
    params.Add("q", query)
    params.Add("per_page", "100")
    params.Add("sort", "author-date")
    params.Add("order", "desc")

    requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

    body, err := c.makeRequest(requestURL)
    if err != nil {
        return nil, err
    }

    var response IssueSearchResult
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, err
    }

    return response.Items, nil
}