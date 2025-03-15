// Package nocodbgo provides a client for the NocoDB v2 API
//
// More info and docs:
//
//   - Overview: https://docs.nocodb.com/developer-resources/rest-APIs/overview
//   - API Explorer: https://data-apis-v2.nocodb.com
//   - OpenAPI spec: https://data-apis-v2.nocodb.com/swagger-v2.json
package nocodbgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultTimeout is the default timeout for HTTP requests
	defaultTimeout = 30 * time.Second
)

// Client is the NocoDB API client
type Client struct {
	// baseURL is the base URL for the NocoDB API
	baseURL string

	// apiToken is the API token for authentication
	apiToken string

	// httpClient is the HTTP client used to make requests
	httpClient *http.Client
}

// NewClient creates a new client builder
func NewClient() *clientBuilder {
	return &clientBuilder{
		httpClient: &http.Client{Timeout: defaultTimeout},
	}
}

// apiError represents an error returned by the NocoDB API
type apiError struct {
	Msg     string `json:"msg"`
	Message string `json:"message"`
	ErrMsg  string `json:"error"`
	Code    string `json:"code"`
}

// Error implements the error interface
func (e apiError) Error() string {
	if e.Code != "" {
		if e.Msg != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.Msg)
		}
		if e.Message != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.Message)
		}
		if e.ErrMsg != "" {
			return fmt.Sprintf("%s: %s", e.Code, e.ErrMsg)
		}
	}

	if e.Msg != "" {
		return e.Msg
	}
	if e.Message != "" {
		return e.Message
	}
	if e.ErrMsg != "" {
		return e.ErrMsg
	}

	return "Unknown error"
}

// request makes an HTTP request to the NocoDB API, it includes the api token in the request header
func (c *Client) request(ctx context.Context, method string, path string, body any, query url.Values) ([]byte, error) {
	parsedUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(path, "/")))
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	if query != nil {
		parsedUrl.RawQuery = query.Encode()
	}

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	var req *http.Request
	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, parsedUrl.String(), reqBody)
	} else {
		req, err = http.NewRequest(method, parsedUrl.String(), reqBody)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("xc-token", c.apiToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr apiError
		if err := json.Unmarshal(respBody, &apiErr); err != nil {
			return nil, fmt.Errorf("status code %d: failed to unmarshal API error: %w", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("status code %d: API error: %s", resp.StatusCode, apiErr.Error())
	}

	return respBody, nil
}

// Table returns a new table with the given ID
func (c *Client) Table(tableID string) *Table {
	return &Table{
		client:  c,
		tableID: tableID,
	}
}

// Table represents a table in NocoDB
type Table struct {
	client  *Client
	tableID string
}
