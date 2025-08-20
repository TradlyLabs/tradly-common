package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TradlyLabs/tradly-common/pkg/runtime"
	"gopkg.in/yaml.v3"
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

type Postgres struct {
	DSN       string   `yaml:"dsn"`
	Sources   []string `yaml:"sources"`
	Replicas  []string `yaml:"replicas"`
	IsDefault bool     `yaml:"isDefault"`
}

type Config struct {
	Postgres map[string]*Postgres    `yaml:"postgres"`
	Redis    map[string]*RedisConfig `yaml:"redis"`
	Evm      map[string]*EvmConfig   `yaml:"evm"`
	Asynq    *AsynqConfig            `yaml:"asynq"`
}

type SrvConfig struct {
	conf *Config
}

func (s *SrvConfig) Start(context.Context) error {
	path := os.Getenv("APP_DIR_PATH")
	data, err := os.ReadFile(filepath.Join(path, MIST_FILE_PATH))
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	s.conf = &config
	return nil
}

func (s *SrvConfig) Stop(context.Context) error {
	return nil
}
