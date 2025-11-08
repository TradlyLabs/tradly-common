package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigUnmarshal(t *testing.T) {
	// Create a temporary config file based on the example
	tempDir := t.TempDir()
	tempConfigPath := filepath.Join(tempDir, "mist.yaml")

	// Copy the example config to the temp location
	exampleConfig, err := os.ReadFile("mist.example.yaml")
	assert.NoError(t, err)

	err = os.WriteFile(tempConfigPath, exampleConfig, 0644)
	assert.NoError(t, err)

	// Load config using viper (simulating how service.go does it)
	v := viper.New()
	v.SetConfigFile(tempConfigPath)

	err = v.ReadInConfig()
	assert.NoError(t, err)

	// Unmarshal into our config struct
	var config Config
	err = v.Unmarshal(&config)
	assert.NoError(t, err)

	// Check that we have the evm configs
	assert.NotNil(t, config.Evm)

	// Check baseSepolia config (Viper normalizes keys to lowercase)
	baseSepolia, exists := config.Evm["basesepolia"]
	assert.True(t, exists, "basesepolia config should exist")
	assert.Equal(t, int64(84532), baseSepolia.ChainID)
	assert.Len(t, baseSepolia.RPCs, 1)
	assert.Equal(t, "https://sepolia.base.org", baseSepolia.RPCs[0].URL)
	assert.Equal(t, 50, baseSepolia.RPCs[0].LimitPerSecond)

	// Check bsc config
	bsc, exists := config.Evm["bsc"]
	assert.True(t, exists, "bsc config should exist")
	assert.Equal(t, int64(56), bsc.ChainID)
	assert.Len(t, bsc.RPCs, 2)
	assert.Equal(t, "wss://bsc-rpc.publicnode.com", bsc.RPCs[0].URL)
	assert.Equal(t, 50, bsc.RPCs[0].LimitPerSecond)
	assert.Equal(t, "https://bsc-rpc.publicnode.com", bsc.RPCs[1].URL)
	assert.Equal(t, 50, bsc.RPCs[1].LimitPerSecond)
}
