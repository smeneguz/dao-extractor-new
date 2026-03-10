package tokenevents

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// ModuleConfig is the YAML configuration for the token-events module.
type ModuleConfig struct {
	DAOs []DAOConfig `yaml:"daos"`
}

// DAOConfig describes a single DAO whose token contracts are monitored.
type DAOConfig struct {
	Symbol    string              `yaml:"symbol"`
	Name      string              `yaml:"name"`
	ChainID   string              `yaml:"chain_id"`
	Contracts []DAOContractConfig `yaml:"contracts"`
}

// DAOContractConfig describes a single token contract belonging to a DAO.
type DAOContractConfig struct {
	Address    string `yaml:"address"`
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	DeployedAt uint64 `yaml:"deployed_at"`
}

// ParseConfig unmarshals rawConfig bytes into a ModuleConfig.
func ParseConfig(rawConfig []byte) (*ModuleConfig, error) {
	var cfg ModuleConfig
	if err := yaml.Unmarshal(rawConfig, &cfg); err != nil {
		return nil, fmt.Errorf("parse token-events config: %w", err)
	}
	if len(cfg.DAOs) == 0 {
		return nil, fmt.Errorf("token-events config: at least one DAO must be configured")
	}
	for i, dao := range cfg.DAOs {
		if dao.Symbol == "" {
			return nil, fmt.Errorf("token-events config: DAO[%d] missing symbol", i)
		}
		if dao.ChainID == "" {
			return nil, fmt.Errorf("token-events config: DAO[%d] (%s) missing chain_id", i, dao.Symbol)
		}
		if len(dao.Contracts) == 0 {
			return nil, fmt.Errorf("token-events config: DAO %s has no contracts", dao.Symbol)
		}
		for _, c := range dao.Contracts {
			if c.Address == "" {
				return nil, fmt.Errorf("token-events config: DAO %s has contract with empty address", dao.Symbol)
			}
		}
	}
	return &cfg, nil
}
