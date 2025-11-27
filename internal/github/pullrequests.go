package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

func (c *Client) GetPullRequests(ctx context.Context, author string, since time.Time) ([]IssueSearchResultItem, error) {
	query := fmt.Sprintf("is:pr author:%s created:>%s", author, since.Format(time.RFC3339))

	// Build URL with properly encoded query parameters
	baseURL := fmt.Sprintf("%s/search/issues", c.BaseURL)
	params := url.Values{}
	params.Add("q", query)
	params.Add("per_page", "100")
	params.Add("sort", "created")
	params.Add("order", "desc")

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	body, err := c.makeRequest(ctx, requestURL)
	if err != nil {
		return nil, err
	}

	var response IssueSearchResult
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Items, nil
}
