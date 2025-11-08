package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestEVMConfigParsing(t *testing.T) {
	// Create a temporary config file based on the example
	tempDir := t.TempDir()
	tempConfigPath := filepath.Join(tempDir, "mist.yaml")

	// Copy the example config to the temp location
	exampleConfig, err := os.ReadFile("mist.example.yaml")
	assert.NoError(t, err)

	err = os.WriteFile(tempConfigPath, exampleConfig, 0644)
	assert.NoError(t, err)

	// Load config using viper
	v := viper.New()
	v.SetConfigFile(tempConfigPath)

	err = v.ReadInConfig()
	assert.NoError(t, err)

	// Check that we can access the baseSepolia config
	baseSepolia := v.GetStringMap("evm.baseSepolia")
	assert.NotNil(t, baseSepolia)

	// Check that we can access the rpc array
	rpc := v.Get("evm.baseSepolia.rpc")
	assert.NotNil(t, rpc)

	// Check that chainId is correctly parsed
	chainId := v.GetInt("evm.baseSepolia.chainId")
	assert.Equal(t, 84532, chainId)
}
