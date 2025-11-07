package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/TradlyLabs/tradly-common/pkg/runtime"
	"github.com/spf13/viper"
)

var MIST_FILE_PATH = filepath.Join("config", "mist.yaml")

var defaultSrvConfig *SrvConfig

func init() {
	defaultSrvConfig = &SrvConfig{}
	runtime.Register("SrvConfig", defaultSrvConfig)
}

func C() *Config {
	return defaultSrvConfig.conf
}

type Config struct {
	Postgres map[string]*PostgresConfig `yaml:"postgres"`
	Redis    map[string]*RedisConfig    `yaml:"redis"`
	Evm      map[string]*EvmConfig      `yaml:"evm"`
	Asynq    *AsynqConfig               `yaml:"asynq"`
}

type SrvConfig struct {
	conf *Config
}

func (s *SrvConfig) Start(context.Context) error {
	v := viper.New()

	// Set default config file path
	path := os.Getenv("APP_DIR_PATH")
	configPath := filepath.Join(path, MIST_FILE_PATH)

	// Configure viper to read from YAML file
	v.SetConfigFile(configPath)

	// Configure viper to read from environment variables
	v.AutomaticEnv()
	v.SetEnvPrefix("TRADLY")                                     // All env vars will be prefixed with TRADLY_
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_")) // Replace dots and dashes with underscores

	// Read config file if it exists
	if _, err := os.Stat(configPath); err == nil {
		if err := v.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Unmarshal config into struct
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	s.conf = &config
	return nil
}

func (s *SrvConfig) Stop(context.Context) error {
	return nil
}
