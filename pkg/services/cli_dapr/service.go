package cli_dapr

import (
	"context"

	"github.com/TemoreIO/temore-common/pkg/runtime"
	dapr "github.com/dapr/go-sdk/client"
)

var defaultCliDapr *CliDapr

func init() {
	defaultCliDapr = &CliDapr{}
	runtime.Register("CliDapr", defaultCliDapr)
}

type CliDapr struct {
	client dapr.Client
}

func C() dapr.Client {
	return defaultCliDapr.client
}

func (s *CliDapr) Start(ctx context.Context) error {
	client, err := dapr.NewClient()
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

func (s *CliDapr) Stop(ctx context.Context) error {
	if s.client != nil {
		s.client.Close()
	}
	return nil
}
