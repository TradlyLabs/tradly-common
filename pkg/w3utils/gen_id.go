package w3utils

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// GenerateID generates a unique token ID using the format "tk_" + keccak(<chain id>:<pair address>:<token address>)
// chainID: The blockchain ID
// pairAddress: The pair address
// tokenAddress: The token address
// Returns the generated token ID string
func GenerateID(prefix string, chainID int64, address ...string) string {
	// Create the input string in the required format
	input := fmt.Sprintf("%s:%d:%s", prefix, chainID, strings.Join(address, ":"))

	// Calculate Keccak-256 hash
	hashedBytes := crypto.Keccak256Hash([]byte(input)).Bytes()

	// Convert the hashed bytes to a hex string
	hashedString := hex.EncodeToString(hashedBytes)

	// Prepend <prefix> to the hashed string
	return hashedString
}

func GenerateTokenID(chainID int64, pairAddress, tokenAddress string) string {
	return GenerateID("tk_", chainID, pairAddress, tokenAddress)
}

func GeneratePairID(chainID int64, pairAddress string) string {
	return GenerateID("pair_", chainID, pairAddress, "")
}
