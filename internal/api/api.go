package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bradford-hamilton/cipher-bin-server/db"
)

// Client makes network API calls to cipherb.in
type Client struct {
	client         CipherBinAPIClient
	APIBaseURL     string
	BrowserBaseURL string
}

// CipherBinAPIClient is used with http.Client and MockClient to allow mocking of services
type CipherBinAPIClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient is a constructor for the ApiClient
func NewClient(browserBaseURL, apiBaseURL string, client CipherBinAPIClient) (*Client, error) {
	apiClient := &Client{
		client:         client,
		BrowserBaseURL: browserBaseURL,
		APIBaseURL:     apiBaseURL,
	}

	return apiClient, nil
}

// PostMessage takes a msg of type *db.Message (this is what the server
// uses and will expect) and posts it to the cipherbin api
func (c *Client) PostMessage(msg *db.Message) error {
	payloadBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", c.APIBaseURL, "/msg"),
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Error: response status: %d", res.StatusCode)
	}

	return nil
}
