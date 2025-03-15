package nocodbgo

import (
	"net/http"
	"strings"
	"time"
)

// ClientBuilder is used to build a new Client with a fluent API
type ClientBuilder struct {
	baseURL    string
	apiToken   string
	httpClient *http.Client
}

// WithBaseURL sets the base URL for the client
func (b *ClientBuilder) WithBaseURL(baseURL string) *ClientBuilder {
	b.baseURL = strings.TrimSuffix(baseURL, "/")
	return b
}

// WithAPIToken sets the API token for the client
func (b *ClientBuilder) WithAPIToken(apiToken string) *ClientBuilder {
	b.apiToken = apiToken
	return b
}

// WithHTTPClient sets the HTTP client for the client
func (b *ClientBuilder) WithHTTPClient(httpClient *http.Client) *ClientBuilder {
	b.httpClient = httpClient
	return b
}

// WithHTTPTimeout sets the timeout for the HTTP client
func (b *ClientBuilder) WithHTTPTimeout(timeout time.Duration) *ClientBuilder {
	if b.httpClient == nil {
		b.httpClient = &http.Client{}
	}
	b.httpClient.Timeout = timeout
	return b
}

// Build creates a new NocoDB client with the configured options
func (b *ClientBuilder) Build() (*Client, error) {
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
