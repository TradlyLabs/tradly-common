package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestDebugConfig(t *testing.T) {
	// Create a temporary config file based on the example
	tempDir := t.TempDir()
	tempConfigPath := filepath.Join(tempDir, "mist.yaml")

	// Copy the example config to the temp location
	exampleConfig, err := os.ReadFile("mist.example.yaml")
	if err != nil {
		t.Fatal(err)
	}

	err = os.WriteFile(tempConfigPath, exampleConfig, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load config using viper
	v := viper.New()
	v.SetConfigFile(tempConfigPath)

	err = v.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}

	// Try unmarshaling with different options
	var config Config
	err = v.Unmarshal(&config)
	if err != nil {
		t.Fatal(err)
	}

	// Print the result
	fmt.Printf("Config EVM map: %+v\n", config.Evm)

	// Check if baseSepolia exists
	if config.Evm != nil {
		for key, value := range config.Evm {
			fmt.Printf("Key: %s, Value: %+v\n", key, value)
			if key == "basesepolia" {
				fmt.Printf("  ChainID: %d\n", value.ChainID)
				fmt.Printf("  RPCs: %+v\n", value.RPCs)
			}
		}
	}
}
