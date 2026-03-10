package rpc

import (
	"context"
	"fmt"

	"github.com/dao-portal/flux/node"
	"gopkg.in/yaml.v3"

	fluxutils "github.com/dao-portal/flux/utils"
)

const NodeType = "evm-rpc"

func NodeBuilder(
	ctx context.Context,
	id string,
	rawConfig []byte,
) (node.Node, error) {
	// Parse the configurations
	var config Config
	err := yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		return nil, fmt.Errorf("unmarshal %s node config: %w", NodeType, err)
	}

	// Validate the configurations
	err = config.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid %s node config: %w", NodeType, err)
	}

	logger := fluxutils.LoggerFromContext(ctx).With().
		Str("node-type", NodeType).
		Str("node-id", id).
		Logger()
	return NewNode(ctx, logger, config)
}
