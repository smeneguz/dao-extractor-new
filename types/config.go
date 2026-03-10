package types

import (
	"fmt"
	"time"
)

type Config struct {
	EtherscanConfig EterscanConfig `yaml:"etherscan"`
}

func (c *Config) Validate() error {
	if err := c.EtherscanConfig.Validate(); err != nil {
		return fmt.Errorf("etherscan config: %w", err)
	}

	return nil
}

// EterscanConfig represents the configuration for the Etherscan client.
type EterscanConfig struct {
	// The URL of the Etherscan API, if not set, the default URL will be used
	Url *string `yaml:"url"`
	// The API key that will be used to authenticate the requests
	ApiKey string `yaml:"api_key"`
	// Requests timeout in seconds
	Timeout time.Duration `yaml:"timeout"`
}

func (c *EterscanConfig) Validate() error {
	if c.ApiKey == "" {
		return fmt.Errorf("etherscan api key is required")
	}

	if c.Url != nil && *c.Url == "" {
		return fmt.Errorf("invalid etherscan url")
	}

	return nil
}
