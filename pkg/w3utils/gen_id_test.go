package w3utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokenID(t *testing.T) {
	// Test case 1: Basic inputs
	chainID := int64(1)
	pairAddress := "0x1234567890123456789012345678901234567890"
	tokenAddress := "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"

	tokenID := GenerateID("tk_", chainID, pairAddress, tokenAddress)

	// Verify the format starts with "tk_"
	assert.Contains(t, tokenID, "tk_")

	// Verify the length is correct ("tk_" + 64 hex characters)
	assert.Len(t, tokenID, 66)

	// Test case 2: Different chain ID
	chainID2 := int64(1000)
	tokenID2 := GenerateID("tk_", chainID2, pairAddress, tokenAddress)
	assert.NotEqual(t, tokenID, tokenID2)

	// Test case 3: Different pair address
	pairAddress2 := "0x9876543210987654321098765432109876543210"
	tokenID3 := GenerateID("tk_", chainID, pairAddress2, tokenAddress)
	assert.NotEqual(t, tokenID, tokenID3)

	// Test case 4: Different token address
	tokenAddress2 := "0xfedcba fedcba fedcba fedcba fedcba fedcba fedcba fe"
	tokenID4 := GenerateID("tk_", chainID, pairAddress, tokenAddress2)
	assert.NotEqual(t, tokenID, tokenID4)
}
