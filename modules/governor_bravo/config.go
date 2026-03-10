package governorbravo

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// ModuleConfig is the YAML configuration for the governor-bravo module.
type ModuleConfig struct {
	DAOs []DAOConfig `yaml:"daos"`
}

// DAOConfig describes a single DAO that uses a GovernorBravo-style contract.
type DAOConfig struct {
	Symbol    string              `yaml:"symbol"`
	Name      string              `yaml:"name"`
	ChainID   string              `yaml:"chain_id"`
	Contracts []DAOContractConfig `yaml:"contracts"`
}

// DAOContractConfig describes a single contract belonging to a DAO.
type DAOContractConfig struct {
	Address    string `yaml:"address"`
	Type       string `yaml:"type"`        // governor_proxy, governor_implementation, timelock, token
	Name       string `yaml:"name"`        // human-readable label
	DeployedAt uint64 `yaml:"deployed_at"` // block number
}

// ParseConfig unmarshals rawConfig bytes into a ModuleConfig.
func ParseConfig(rawConfig []byte) (*ModuleConfig, error) {
	var cfg ModuleConfig
	if err := yaml.Unmarshal(rawConfig, &cfg); err != nil {
		return nil, fmt.Errorf("parse governor-bravo config: %w", err)
	}
	if len(cfg.DAOs) == 0 {
		return nil, fmt.Errorf("governor-bravo config: at least one DAO must be configured")
	}
	for i, dao := range cfg.DAOs {
		if dao.Symbol == "" {
			return nil, fmt.Errorf("governor-bravo config: DAO[%d] missing symbol", i)
		}
		if dao.ChainID == "" {
			return nil, fmt.Errorf("governor-bravo config: DAO[%d] (%s) missing chain_id", i, dao.Symbol)
		}
		hasGovernor := false
		for _, c := range dao.Contracts {
			if c.Address == "" {
				return nil, fmt.Errorf("governor-bravo config: DAO %s has contract with empty address", dao.Symbol)
			}
			if c.Type == "governor_proxy" || c.Type == "governor" {
				hasGovernor = true
			}
		}
		if !hasGovernor {
			return nil, fmt.Errorf("governor-bravo config: DAO %s has no governor_proxy or governor contract", dao.Symbol)
		}
	}
	return &cfg, nil
}
