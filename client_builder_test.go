package nocodbgo

import (
	"net/http"
	"testing"
	"time"
)

func TestClientBuilder(t *testing.T) {
	// Test successful build
	client, err := NewClient().
		WithBaseURL("https://example.com").
		WithAPIToken("test-token").
		WithHTTPTimeout(30 * time.Second).
		Build()

	if err != nil {
		t.Errorf("Build() error = %v, want nil", err)
	}

	//nolint:all
	if client == nil {
		t.Error("Build() client is nil, want non-nil")
	}

	//nolint:all
	if client.baseURL != "https://example.com" {
		t.Errorf("Build() baseURL = %v, want %v", client.baseURL, "https://example.com")
	}

	if client.apiToken != "test-token" {
		t.Errorf("Build() apiToken = %v, want %v", client.apiToken, "test-token")
	}

	if client.httpClient.Timeout != 30*time.Second {
		t.Errorf("Build() httpClient.Timeout = %v, want %v", client.httpClient.Timeout, 30*time.Second)
	}

	// Test with custom HTTP client
	customClient := &http.Client{Timeout: 60 * time.Second}
	client, err = NewClient().
		WithBaseURL("https://example.com").
		WithAPIToken("test-token").
		WithHTTPClient(customClient).
		Build()

	if err != nil {
		t.Errorf("Build() error = %v, want nil", err)
	}

	if client.httpClient != customClient {
		t.Error("Build() httpClient is not the custom client")
	}

	// Test with trailing slash in base URL
	client, err = NewClient().
		WithBaseURL("https://example.com/").
		WithAPIToken("test-token").
		WithHTTPTimeout(30 * time.Second).
		Build()

	if err != nil {
		t.Errorf("Build() error = %v, want nil", err)
	}

	if client.baseURL != "https://example.com" {
		t.Errorf("Build() baseURL = %v, want %v", client.baseURL, "https://example.com")
	}

	// Test error cases
	_, err = NewClient().
		WithAPIToken("test-token").
		WithHTTPTimeout(30 * time.Second).
		Build()

	if err != ErrBaseURLRequired {
		t.Errorf("Build() error = %v, want %v", err, ErrBaseURLRequired)
	}

	_, err = NewClient().
		WithBaseURL("https://example.com").
		WithHTTPTimeout(30 * time.Second).
		Build()

	if err != ErrAPITokenRequired {
		t.Errorf("Build() error = %v, want %v", err, ErrAPITokenRequired)
	}

	// Create a builder with nil HTTP client
	builder := &clientBuilder{
		baseURL:  "https://example.com",
		apiToken: "test-token",
	}

	_, err = builder.Build()

	if err != ErrHTTPClientRequired {
		t.Errorf("Build() error = %v, want %v", err, ErrHTTPClientRequired)
	}
}
