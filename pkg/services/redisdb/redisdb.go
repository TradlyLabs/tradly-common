package redisdb

import (
	"context"

	"github.com/TemoreIO/temore-common/pkg/config"
	"github.com/TemoreIO/temore-common/pkg/runtime"
	"github.com/TemoreIO/temore-common/pkg/runtime/service"
	redis "github.com/redis/go-redis/v9"
)

const DEFAULT_NAME = "default"

type SrvRedisDB interface {
	GetClient(string) (Client, bool)
	service.Service
}

// srvRedisDB manages Redis clients
type srvRedisDB struct {
	clients map[string]*redis.Client
	list    []*redis.Client
}

var defaultSrvRedisDB SrvRedisDB

func init() {
	defaultSrvRedisDB = &srvRedisDB{
		clients: make(map[string]*redis.Client),
	}
	runtime.Register("SrvRedisDB", defaultSrvRedisDB)
}

func SetDefaultSrvRedisDB(s SrvRedisDB) {
	defaultSrvRedisDB = s
}

func Get(args ...interface{}) Client {
	if len(args) > 0 {
		switch v := args[0].(type) {
		case string:
			client, ok := defaultSrvRedisDB.GetClient(v)
			if ok {
				return client
			}
			panic("database " + v + " is not configed")
		default:
			panic("database name must be string")
		}
	}
	db, ok := defaultSrvRedisDB.GetClient(DEFAULT_NAME)
	if ok {
		return db
	}
	panic("no database")
}

func (s *srvRedisDB) GetClient(name string) (Client, bool) {
	db, ok := s.clients[name]
	return db, ok
}

// Start start Redis service
func (s *srvRedisDB) Start(ctx context.Context) error {
	conf := config.C()

	hasDefault := false
	for key, c := range conf.Redis {
		client := redis.NewClient(&redis.Options{
			Addr:     c.Address,
			Password: c.Password,
			DB:       c.DB,
		})
		if err := client.Ping(ctx).Err(); err != nil {
			return err
		}
		s.clients[key] = client
		s.list = append(s.list, client)
	}
	if !hasDefault && len(s.list) > 0 {
		s.clients[DEFAULT_NAME] = s.list[0]
	}
	return nil
}

// Stop stop Redis service
func (s *srvRedisDB) Stop(ctx context.Context) error {
	var errors []error
	for _, d := range s.list {
		if d != nil {
			err := d.Close()
			if err != nil {
				errors = append(errors, err)
			}
		}
	}
	if len(errors) > 0 {
		return runtime.MultiError(errors)
	}
	return nil
}
