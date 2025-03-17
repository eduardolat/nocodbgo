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

// Client provides access to the NocoDB API
type Client struct {
	// baseURL is the base URL for the NocoDB API
	baseURL string

	// apiToken is the API token for authentication
	apiToken string

	// httpClient is the HTTP client used to make requests
	httpClient *http.Client
}

// NewClient creates a new client builder for configuring and creating a NocoDB client
func NewClient() *clientBuilder {
	return &clientBuilder{
		httpClient: &http.Client{Timeout: defaultTimeout},
	}
}

// clientBuilder is used to build a new Client with a fluent API
type clientBuilder struct {
	baseURL    string
	apiToken   string
	httpClient *http.Client
}

// WithBaseURL sets the base URL for the NocoDB API.
// The URL should be the root URL of your NocoDB instance (e.g., "https://nocodb.example.com").
// Any trailing slashes will be automatically removed.
// Returns the clientBuilder for method chaining.
func (b *clientBuilder) WithBaseURL(baseURL string) *clientBuilder {
	b.baseURL = strings.TrimSuffix(baseURL, "/")
	return b
}

// WithAPIToken sets the API token for authentication with the NocoDB API.
// You can generate an API token from the NocoDB user interface.
// Returns the clientBuilder for method chaining.
func (b *clientBuilder) WithAPIToken(apiToken string) *clientBuilder {
	b.apiToken = apiToken
	return b
}

// WithHTTPClient sets a custom HTTP client for making requests to the NocoDB API.
// This allows for custom configuration such as proxies, custom transports, etc.
// Returns the clientBuilder for method chaining.
func (b *clientBuilder) WithHTTPClient(httpClient *http.Client) *clientBuilder {
	b.httpClient = httpClient
	return b
}

// WithHTTPTimeout sets the timeout duration for HTTP requests.
// If no HTTP client has been set, a new one will be created.
// Returns the clientBuilder for method chaining.
func (b *clientBuilder) WithHTTPTimeout(timeout time.Duration) *clientBuilder {
	if b.httpClient == nil {
		b.httpClient = &http.Client{}
	}
	b.httpClient.Timeout = timeout
	return b
}

// Create builds and returns a new NocoDB client with the configured options.
// Returns an error if any required configuration is missing.
func (b *clientBuilder) Create() (*Client, error) {
	if b.baseURL == "" {
		return nil, ErrBaseURLRequired
	}

	if b.apiToken == "" {
		return nil, ErrAPITokenRequired
	}

	if b.httpClient == nil {
		return nil, ErrHTTPClientRequired
	}

	return &Client{
		baseURL:    b.baseURL,
		apiToken:   b.apiToken,
		httpClient: b.httpClient,
	}, nil
}

// apiError represents an error returned by the NocoDB API
type apiError struct {
	Msg     string `json:"msg"`
	Message string `json:"message"`
	ErrMsg  string `json:"error"`
	Code    string `json:"code"`
}

// Error implements the error interface for apiError
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

// request makes an HTTP request to the NocoDB API with the provided method, path, body, and query parameters.
// It automatically includes the API token in the request header.
// Returns the response body as a byte slice or an error if the request fails.
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

	if ctx == nil {
		ctx = context.Background()
	}
	req, err := http.NewRequestWithContext(ctx, method, parsedUrl.String(), reqBody)

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

// Table returns a new Table instance for the specified table ID.
// This instance provides methods for performing CRUD operations on the table's records.
func (c *Client) Table(tableID string) *Table {
	return &Table{
		client:  c,
		tableID: tableID,
	}
}

// Table represents a table in NocoDB and provides methods for interacting with its records
type Table struct {
	client  *Client
	tableID string
}
