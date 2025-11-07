package config

// PostgresConfig stores PostgreSQL connection information
type PostgresConfig struct {
	DSN       string   `yaml:"dsn"`
	Sources   []string `yaml:"sources"`
	Replicas  []string `yaml:"replicas"`
	IsDefault bool     `yaml:"isDefault"`
}

// RedisConfig stores Redis connection information
type RedisConfig struct {
	Address   string `yaml:"address"`
	Password  string `yaml:"password"`
	DB        int    `yaml:"db"`
	IsDefault bool   `yaml:"isDefault"`
}

// EvmConfig stores EVM connection information
type EvmConfig struct {
	ChainID          int64          `yaml:"chainId"`
	RPCs             []EvmConfigRPC `yaml:"rpc"`
	StartBlockNumber int64          `yaml:"startBlockNumber"`
}

type EvmConfigRPC struct {
	URL            string `yaml:"url"`
	LimitPerSecond int    `yaml:"limitPerSecond"`
}

type AsynqConfig struct {
	Redis       string         `yaml:"redis"` // use redis config named default
	Queues      map[string]int `yaml:"queues"`
	Concurrency int            `yaml:"concurrency"`
}
