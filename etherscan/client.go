package etherscan

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dao-portal/extractor/types"
)

const DefaultURL = "https://api.etherscan.io/v2/api"

// Client represents an Etherscan client.
type Client struct {
	url    string
	apiKey string

	httpClient *http.Client
}

// NewClient returns a new Etherscan client.
func NewClient(cfg *types.EterscanConfig) *Client {
	url := DefaultURL
	if cfg.Url != nil {
		url = *cfg.Url
	}
	var timeaOut time.Duration
	if cfg.Timeout != 0 {
		timeaOut = cfg.Timeout
	} else {
		timeaOut = 10 * time.Second
	}

	return &Client{
		url:    url,
		apiKey: cfg.ApiKey,
		httpClient: &http.Client{
			Timeout: timeaOut,
		},
	}
}

// GetContractAbi returns the ABI of the contract with the given address and chain ID.
func (c *Client) GetContractAbi(ctx context.Context, chainID string, contractAddress string) ([]byte, error) {
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, fmt.Errorf("invalid client url: %w", err)
	}

	// Remove the 0x prefix
	chainID = strings.ReplaceAll(chainID, "0x", "")

	q := u.Query()
	q.Set("chainid", chainID)
	q.Set("module", "contract")
	q.Set("action", "getabi")
	q.Set("address", contractAddress)
	q.Set("apikey", c.apiKey)
	u.RawQuery = q.Encode()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get contract ABI: status code %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var response EtherscanResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Etherscan response: %w", err)
	}

	if err := response.Error(); err != nil {
		return nil, fmt.Errorf("failed to get contract ABI: %w", err)
	}

	return []byte(response.Result), nil
}
