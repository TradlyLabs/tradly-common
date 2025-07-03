package srv_dapr

import (
	"context"
	"os"

	"github.com/TemoreIO/temore-common/pkg/runtime"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"go.uber.org/zap"
)

var DefaultDaprSrv *SrvDapr

func init() {
	DefaultDaprSrv = NewSrvDapr()
	runtime.DefaultManager.Register("SrvDapr", DefaultDaprSrv)
}

type SrvDapr struct {
	s common.Service
}

func NewSrvDapr() *SrvDapr {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000"
	}
	s := daprd.NewService(":" + appPort)
	return &SrvDapr{s}
}

func (s *SrvDapr) Start(ctx context.Context) error {
	go func() {
		err := s.s.Start()
		if err != nil {
			zap.L().Error("dapr srv start error", zap.Error(err))
		}
	}()
	return nil
}

func (s *SrvDapr) Stop(ctx context.Context) error {
	return s.s.Stop()
}

func AddTopicEventHandler(sub *common.Subscription, fn common.TopicEventHandler) error {
	return DefaultDaprSrv.s.AddTopicEventHandler(sub, fn)
}
