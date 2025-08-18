package cli_asynq

import (
	"context"

	"github.com/TradlyLabs/tradly-common/pkg/config"
	"github.com/TradlyLabs/tradly-common/pkg/runtime"
	"github.com/hibiken/asynq"
)

const DEFAULT_NAME = "default"

var defaultCliAsynq *CliAsynq

func init() {
	defaultCliAsynq = &CliAsynq{
		clients: make(map[string]*asynq.Client),
	}
	runtime.Register("CliAsynq", defaultCliAsynq)
}

type CliAsynq struct {
	clients map[string]*asynq.Client
	list    []*asynq.Client
}

// Client returns the asynq client by name.
// If name is empty, it returns the default client.
// Names comming from config file redis section.
func C(args ...interface{}) *asynq.Client {
	name := DEFAULT_NAME
	if len(args) > 0 {
		if n, ok := args[0].(string); ok {
			name = n
		}
	}
	if client, ok := defaultCliAsynq.clients[name]; ok {
		return client
	}
	if client, ok := defaultCliAsynq.clients[DEFAULT_NAME]; ok {
		return client
	}
	panic("asynq client not found: " + name)
}

func (s *CliAsynq) Start(ctx context.Context) error {
	conf := config.C()

	hasDefault := false
	for key, c := range conf.Redis {
		client := asynq.NewClient(asynq.RedisClientOpt{
			Addr:     c.Address,
			Password: c.Password,
			DB:       c.DB,
		})
		s.clients[key] = client
		s.list = append(s.list, client)
	}
	if !hasDefault && len(s.list) > 0 {
		s.clients[DEFAULT_NAME] = s.list[0]
	}
	return nil
}

func (s *CliAsynq) Stop(ctx context.Context) error {
	for _, client := range s.clients {
		if client != nil {
			client.Close()
		}
	}
	return nil
}
