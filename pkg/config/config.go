package config

// PostgresConfig stores PostgreSQL connection information
type PostgresConfig struct {
	DSN      string   `yaml:"dsn"`
	Sources  []string `yaml:"sources"`
	Replicas []string `yaml:"replicas"`
}

// RedisConfig stores Redis connection information
type RedisConfig struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// EvmConfig stores EVM connection information
type EvmConfig struct {
	ChainID int64    `yaml:"chainId"`
	RPCs    []string `yaml:"rpc"`
}

// C returns current configuration
func (c *Config) C() *Config {
	return c
}
