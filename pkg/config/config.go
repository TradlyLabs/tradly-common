package config

// PostgresConfig stores PostgreSQL connection information
type PostgresConfig struct {
	DSN       string   `mapstructure:"dsn" yaml:"dsn"`
	Sources   []string `mapstructure:"sources" yaml:"sources"`
	Replicas  []string `mapstructure:"replicas" yaml:"replicas"`
	IsDefault bool     `mapstructure:"isDefault" yaml:"isDefault"`
}

// RedisConfig stores Redis connection information
type RedisConfig struct {
	Address   string `mapstructure:"address" yaml:"address"`
	Password  string `mapstructure:"password" yaml:"password"`
	DB        int    `mapstructure:"db" yaml:"db"`
	IsDefault bool   `mapstructure:"isDefault" yaml:"isDefault"`
}

// EvmConfig stores EVM connection information
type EvmConfig struct {
	ChainID          int64          `mapstructure:"chainId" yaml:"chainId"`
	RPCs             []EvmConfigRPC `mapstructure:"rpc" yaml:"rpc"`
	StartBlockNumber int64          `mapstructure:"startBlockNumber" yaml:"startBlockNumber"`
}

type EvmConfigRPC struct {
	URL            string `mapstructure:"url" yaml:"url"`
	LimitPerSecond int    `mapstructure:"limitPerSecond" yaml:"limitPerSecond"`
}

type AsynqConfig struct {
	Redis       string         `mapstructure:"redis" yaml:"redis"` // use redis config named default
	Queues      map[string]int `mapstructure:"queues" yaml:"queues"`
	Concurrency int            `mapstructure:"concurrency" yaml:"concurrency"`
}
