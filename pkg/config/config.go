package config

// PostgresConfig stores PostgreSQL connection information
type PostgresConfig struct {
	DSN      string   `json:"dsn"`
	Sources  []string `json:"sources"`
	Replicas []string `json:"replicas"`
}

// RedisConfig stores Redis connection information
type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// EvmConfig stores EVM connection information
type EvmConfig struct {
	ChainID int64    `json:"chainId"`
	RPCs    []string `json:"rpc"`
}

// C returns current configuration
func (c *Config) C() *Config {
	return c
}
