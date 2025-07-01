package mocks

import (
	"context"

	"github.com/TemoreIO/temore-common/pkg/services/redisdb"
)

var defaultSrvRedisDB *MockSrvRedisDB

type MockSrvRedisDB struct {
	client *MockRedisClient
}

func (s *MockSrvRedisDB) Start(ctx context.Context) error {
	ctrl := GetController(ctx)
	s.client = NewMockRedisClient(ctrl)
	return nil
}

func (s *MockSrvRedisDB) Stop(context.Context) error {
	return nil
}

func (s *MockSrvRedisDB) GetClient(string) (redisdb.Client, bool) {
	return s.client, true
}

func InitMockRedisdb() {
	defaultSrvRedisDB = &MockSrvRedisDB{}
	redisdb.SetDefaultSrvRedisDB(defaultSrvRedisDB)
	TestServiceManager.Register("testSrvRedisDB", defaultSrvRedisDB)
}

func RedisdbGet() *MockRedisClient {
	return defaultSrvRedisDB.client
}
