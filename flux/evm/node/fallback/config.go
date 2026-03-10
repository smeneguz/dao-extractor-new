package fallback

import (
	"fmt"
	"time"

	rpcnode "github.com/dao-portal/extractor/flux/evm/node/rpc"
)

// Config for the fallback node.
type Config struct {
	Primary          rpcnode.Config `yaml:"primary"`
	Fallback         rpcnode.Config `yaml:"fallback"`
	FallbackCooldown time.Duration  `yaml:"fallback_cooldown"`
}

func (c *Config) Validate() error {
	if err := c.Primary.Validate(); err != nil {
		return fmt.Errorf("primary: %w", err)
	}
	if err := c.Fallback.Validate(); err != nil {
		return fmt.Errorf("fallback: %w", err)
	}
	return nil
}

// UnmarshalYAML implements the Unmarshaler interface.
func (c *Config) UnmarshalYAML(unmarshal func(any) error) error {
	type privateCfg Config
	config := privateCfg{
		FallbackCooldown: 500 * time.Millisecond,
	}
	err := unmarshal(&config)
	if err != nil {
		return err
	}
	*c = Config(config)
	return nil
}
