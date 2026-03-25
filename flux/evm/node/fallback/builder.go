package fallback

import (
	"context"
	"fmt"

	"github.com/dao-portal/flux/node"
	fluxutils "github.com/dao-portal/flux/utils"
	"gopkg.in/yaml.v3"

	rpcnode "github.com/dao-portal/extractor/flux/evm/node/rpc"
)

// NodeType is the identifier for the fallback node in config.yaml.
const NodeType = "evm-rpc-fallback"

// NodeBuilder constructs a fallback node from YAML config.
func NodeBuilder(
	ctx context.Context,
	id string,
	rawConfig []byte,
) (node.Node, error) {
	var config Config
	err := yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal %s node config: %w", NodeType, err)
	}

	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid %s node config: %w", NodeType, err)
	}

	logger := fluxutils.LoggerFromContext(ctx).With().
		Str("node-type", NodeType).
		Str("node-id", id).
		Logger()

	primary, err := rpcnode.NewNode(ctx, logger.With().Str("role", "primary").Logger(), config.Primary)
	if err != nil {
		return nil, fmt.Errorf("create primary node: %w", err)
	}

	fallbackNodes := make([]*rpcnode.Node, 0, len(config.Fallbacks))
	fallbackURLs := make([]string, 0, len(config.Fallbacks))
	for i, fbCfg := range config.Fallbacks {
		fbNode, err := rpcnode.NewNode(
			ctx,
			logger.With().Str("role", fmt.Sprintf("fallback-%d", i)).Logger(),
			fbCfg,
		)
		if err != nil {
			return nil, fmt.Errorf("create fallback node [%d]: %w", i, err)
		}
		fallbackNodes = append(fallbackNodes, fbNode)
		fallbackURLs = append(fallbackURLs, fbCfg.URL)
	}

	logger.Info().
		Str("primary", config.Primary.URL).
		Strs("fallbacks", fallbackURLs).
		Dur("cooldown", config.FallbackCooldown).
		Msg("fallback node initialized")

	return NewNode(primary, fallbackNodes, logger, config.FallbackCooldown), nil
}
