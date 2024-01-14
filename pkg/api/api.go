package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cipherbin/cipher-bin-server/app"
	"github.com/cipherbin/cipher-bin-server/db"
)

// Client makes network API calls to cipherb.in
type Client struct {
	cipherBinAPIClient
	APIBaseURL     string
	BrowserBaseURL string
}

// cipherBinAPIClient is used with http.Client and MockClient to allow mocking of services
type cipherBinAPIClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewClient is a constructor for the ApiClient and satisfies the CipherBinAPIClient interface
func NewClient(browserBaseURL, apiBaseURL string, client cipherBinAPIClient) (*Client, error) {
	return &Client{
		cipherBinAPIClient: client,
		BrowserBaseURL:     browserBaseURL,
		APIBaseURL:         apiBaseURL,
	}, nil
}

// PostMessage takes a msg of type *db.Message (this is what the server uses and will expect)
// and posts it to the live cipherbin api
func (c *Client) PostMessage(msg *db.Message) error {
	payloadBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, c.APIBaseURL+"/msg", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", "application/json")

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error: response status: %d", res.StatusCode)
	}
	return nil
}

// GetMessage simply takes a cipherb.in public URL string and returns the appropriate encrypted message
func (c *Client) GetMessage(url string) (*app.MessageResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New("sorry, this message has either already been viewed and destroyed or it never existed at all")
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: response status: %d", res.StatusCode)
	}

	var r app.MessageResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
