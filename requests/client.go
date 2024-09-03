package requests

import (
	"GPT_cli/global"
	gptCliErrors "GPT_cli/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Config struct {
	// Base URL for API requests.
	BaseURL string

	// API Key (Required)
	APIKey string

	// Organization ID (Optional)
	OrganizationID string
}

type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Config
	config *Config
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.APIKey))
	if c.config.OrganizationID != "" {
		req.Header.Set("OpenAI-Organization", c.config.OrganizationID)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		// Parse body
		var errMessage interface{}
		if err := json.NewDecoder(res.Body).Decode(&errMessage); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("api request failed: status Code: %d %s %s Message: %+v", res.StatusCode, res.Status, res.Request.URL, errMessage)
	}

	return res, nil
}

func NewClient() (*Client, error) {
	apikey, apiUrl := "", ""
	for _, x := range global.TokenSetting.Models {
		if x.Enable {
			apikey = x.Token
			apiUrl = x.ApiURL
			break
		}
	}
	if apikey == "" || apiUrl == "" {
		return nil, gptCliErrors.ErrAPIKeyRequired
	}

	return &Client{
		client: &http.Client{},
		config: &Config{
			BaseURL: apiUrl,
			APIKey:  apikey,
		},
	}, nil
}

func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" || config.BaseURL == "" {
		return nil, gptCliErrors.ErrAPIKeyRequired
	}

	return &Client{
		client: &http.Client{},
		config: config,
	}, nil
}
