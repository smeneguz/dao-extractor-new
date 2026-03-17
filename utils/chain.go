package utils

// ChainNameFromID maps well-known EVM chain IDs to human-readable names.
func ChainNameFromID(id string) string {
	switch id {
	case "1":
		return "Ethereum"
	case "42161":
		return "Arbitrum"
	case "10":
		return "Optimism"
	case "137":
		return "Polygon"
	case "8453":
		return "Base"
	case "100":
		return "Gnosis"
	case "1135":
		return "Lisk"
	default:
		return "Chain-" + id
	}
}
