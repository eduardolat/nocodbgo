package nocodbgo

import (
	"net/http"
	"strings"
	"time"
)

// clientBuilder is used to build a new Client with a fluent API
type clientBuilder struct {
	baseURL    string
	apiToken   string
	httpClient *http.Client
}

// WithBaseURL sets the base URL for the client
func (b *clientBuilder) WithBaseURL(baseURL string) *clientBuilder {
	b.baseURL = strings.TrimSuffix(baseURL, "/")
	return b
}

// WithAPIToken sets the API token for the client
func (b *clientBuilder) WithAPIToken(apiToken string) *clientBuilder {
	b.apiToken = apiToken
	return b
}

// WithHTTPClient sets the HTTP client for the client
func (b *clientBuilder) WithHTTPClient(httpClient *http.Client) *clientBuilder {
	b.httpClient = httpClient
	return b
}

// WithHTTPTimeout sets the timeout for the HTTP client
func (b *clientBuilder) WithHTTPTimeout(timeout time.Duration) *clientBuilder {
	if b.httpClient == nil {
		b.httpClient = &http.Client{}
	}
	b.httpClient.Timeout = timeout
	return b
}

// Create creates the NocoDB client with the configured options
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
