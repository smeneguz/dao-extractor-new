package fallback

import (
	"fmt"
	"time"

	rpcnode "github.com/dao-portal/extractor/flux/evm/node/rpc"
)

// FallbackList is a list of node configs that can be unmarshaled from
// either a single object or a YAML sequence, for backward compatibility.
//
//	fallback:                     # old format – single object
//	  url: "https://..."
//
//	fallback:                     # new format – list
//	  - url: "https://first..."
//	  - url: "https://second..."
type FallbackList []rpcnode.Config

func (f *FallbackList) UnmarshalYAML(unmarshal func(any) error) error {
	// Try list first.
	var list []rpcnode.Config
	if err := unmarshal(&list); err == nil {
		*f = list
		return nil
	}
	// Fall back to single object.
	var single rpcnode.Config
	if err := unmarshal(&single); err != nil {
		return fmt.Errorf("fallback: expected a node config or list of node configs: %w", err)
	}
	*f = []rpcnode.Config{single}
	return nil
}

// Config for the fallback node.
type Config struct {
	Primary          rpcnode.Config `yaml:"primary"`
	Fallbacks        FallbackList   `yaml:"fallback"`
	FallbackCooldown time.Duration  `yaml:"fallback_cooldown"`
}

func (c *Config) Validate() error {
	if err := c.Primary.Validate(); err != nil {
		return fmt.Errorf("primary: %w", err)
	}
	if len(c.Fallbacks) == 0 {
		return fmt.Errorf("fallback: at least one fallback node is required")
	}
	for i, fb := range c.Fallbacks {
		if err := fb.Validate(); err != nil {
			return fmt.Errorf("fallback[%d]: %w", i, err)
		}
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
