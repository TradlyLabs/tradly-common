package w3utils

import (
	"testing"
)

func TestGenerateTokenID(t *testing.T) {
	// Test case 1: Basic inputs
	chainID := int64(84532)
	// pairAddress := "0x62E0984a778857BAc86386B67FFf5d51ed4699D8"
	tokenAddress := "0x85dEfe4bA6a088b74ad65fe0D547Dc3b4c21D9dE"

	expectedID := GenerateID("token", chainID, tokenAddress)
	t.Log(expectedID)
}
