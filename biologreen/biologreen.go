// biologreen/biologreen.go

package biologreen

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// This is the default production URL.
const defaultBaseURL = "https://api.biologreen.com/v1"

// Client is the main client for interacting with the BioLogreen API.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new BioLogreen API client.
// The base URL is an optional second argument; if not provided, it defaults to the production API.
func NewClient(apiKey string, baseURL ...string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}

	// Determine the base URL to use.
	url := defaultBaseURL
	if len(baseURL) > 0 && baseURL[0] != "" {
		url = baseURL[0]
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    url,
		httpClient: &http.Client{},
	}, nil
}

// _post is a helper function to handle making POST requests to the API.
func (c *Client) _post(endpoint string, payload interface{}, responseTarget interface{}) error {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal request payload: %w", err)
	}

	// Create the request
	url := c.baseURL + endpoint
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-KEY", c.apiKey)

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Handle non-successful status codes
	if resp.StatusCode >= 400 {
		var errorResponse apiErrorResponse
		if json.Unmarshal(body, &errorResponse) == nil && errorResponse.Detail != "" {
			return fmt.Errorf("API error (%d): %s", resp.StatusCode, errorResponse.Detail)
		}
		// Fallback error if the error response itself is not valid JSON
		return fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	// Decode successful response
	if err := json.Unmarshal(body, responseTarget); err != nil {
		return fmt.Errorf("failed to unmarshal successful response: %w", err)
	}

	return nil
}

// SignupFace registers a new user by their face.
func (c *Client) SignupFace(request SignupRequest) (*FaceAuthResponse, error) {
	var response FaceAuthResponse
	err := c._post("/auth/signup-face", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// LoginFace authenticates an existing user by their face.
func (c *Client) LoginFace(request LoginRequest) (*FaceAuthResponse, error) {
	var response FaceAuthResponse
	err := c._post("/auth/login-face", request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ImageFileToBase64 is a helper utility to convert an image file to a Base64 string.
func ImageFileToBase64(filePath string) (string, error) {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read image file at '%s': %w", filePath, err)
	}
	return base64.StdEncoding.EncodeToString(fileBytes), nil
}
