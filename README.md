# Tradly Common

Tradly Common is a shared library containing common utilities, configurations, and services used across Tradly platform services.

## Configuration

The application can be configured using both YAML configuration files and environment variables. Environment variables take precedence over values defined in the YAML file.

### Environment Variables

All environment variables are prefixed with ``.

#### PostgreSQL Configuration

You can configure PostgreSQL connections using the following environment variable pattern:
```
POSTGRES_[NAME]_DSN=...
POSTGRES_[NAME]_SOURCES=...
POSTGRES_[NAME]_REPLICAS=...
POSTGRES_[NAME]_ISDEFAULT=true|false
```

Example:
```bash
POSTGRES_DSN="host=localhost user=tradly password=tradly2025 dbname=tradly port=5432 sslmode=disable TimeZone=Asia/Shanghai"
POSTGRES_ISDEFAULT=true
```

#### Redis Configuration

You can configure Redis connections using the following environment variable pattern:
```
REDIS_[NAME]_ADDRESS=...
REDIS_[NAME]_PASSWORD=...
REDIS_[NAME]_DB=...
REDIS_[NAME]_ISDEFAULT=true|false
```

Example:
```bash
REDIS_TEMORE_ADDRESS="localhost:6379"
REDIS_TEMORE_PASSWORD=""
REDIS_TEMORE_DB=0
REDIS_TEMORE_ISDEFAULT=true
```

#### EVM Configuration

You can configure EVM chains using the following environment variable pattern:
```
EVM_[NAME]_CHAINID=...
EVM_[NAME]_STARTBLOCKNUMBER=...
EVM_[NAME]_RPC=...
```

Example:
```bash
EVM_BSC_CHAINID=56
EVM_BSC_RPC_0_URL="wss://bsc-rpc.publicnode.com"
EVM_BSC_RPC_0_LIMITPERSECOND=50
```

#### Asynq Configuration

You can configure Asynq using the following environment variables:
```
ASYNQ_REDIS=...
ASYNQ_CONCURRENCY=...
ASYNQ_QUEUES_[NAME]=...
```

Example:
```bash
ASYNQ_REDIS=temore
ASYNQ_CONCURRENCY=10
ASYNQ_QUEUES_DEFAULT=10
```

### YAML Configuration File

A sample configuration file is provided at `config/mist.example.yaml`. To use it, copy it to `config/mist.yaml` in your application directory and modify as needed.

The structure of the YAML file is as follows:

```yaml
postgres:
  [name]:
    dsn: "..."
    sources: [...]
    replicas: [...]
    isDefault: true|false
redis:
  [name]:
    address: "..."
    password: "..."
    db: 0
    isDefault: true|false
evm:
  [name]:
    chainId: ...
    startBlockNumber: ...
    rpc:
      - url: "..."
        limitPerSecond: ...
```

Environment variables will override values from the YAML configuration file when both are present.