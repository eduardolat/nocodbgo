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
