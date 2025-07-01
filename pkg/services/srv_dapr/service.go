package srv_dapr

import (
	"context"
	"os"

	"github.com/TemoreIO/temore-common/pkg/runtime"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
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
		appPort = "6005"
	}
	s := daprd.NewService(":" + appPort)
	return &SrvDapr{s}
}

func (s *SrvDapr) Start(ctx context.Context) error {
	return s.s.Start()
}

func (s *SrvDapr) Stop(ctx context.Context) error {
	return s.s.Stop()
}

func AddTopicEventHandler(sub *common.Subscription, fn common.TopicEventHandler) error {
	return DefaultDaprSrv.s.AddTopicEventHandler(sub, fn)
}
