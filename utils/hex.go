package utils

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func Has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func Remove0xPrefix(s string) string {
	if Has0xPrefix(s) {
		return s[2:]
	}
	return s
}

// HexChainIDToDecimal converts a hex chain ID (e.g. "0x2105") to its decimal
// string representation (e.g. "8453").
func HexChainIDToDecimal(hexID string) string {
	clean := Remove0xPrefix(hexID)
	n, err := strconv.ParseUint(clean, 16, 64)
	if err != nil {
		return clean // fallback: return as-is
	}
	return strconv.FormatUint(n, 10)
}

func HexToBytes(s string) ([]byte, error) {
	if Has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}

	return hex.DecodeString(s)
}

func IsValidEthAddress(s string) error {
	bytes, err := HexToBytes(s)
	if err != nil || len(bytes) > common.AddressLength {
		return fmt.Errorf("invalid address: %w", err)
	}
	return nil
}
