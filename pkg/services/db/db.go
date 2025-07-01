package db

import (
	"context"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/TemoreIO/temore-common/pkg/config"
	"github.com/TemoreIO/temore-common/pkg/runtime"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

const DEFAULT_NAME = "default"

var defaultSrvDB *SrvDB

func init() {
	defaultSrvDB = &SrvDB{
		dbs: make(map[string]*gorm.DB),
	}
	runtime.Register("SrvDB", defaultSrvDB)
}

func Get(args ...interface{}) *gorm.DB {
	if len(args) > 0 {
		switch v := args[0].(type) {
		case string:
			db, ok := defaultSrvDB.dbs[v]
			if ok {
				return db
			}
			panic("database " + v + " is not configed")
		default:
			panic("database name must be string")
		}
	}
	db, ok := defaultSrvDB.dbs[DEFAULT_NAME]
	if ok {
		return db
	}
	panic("no database")
}

type SrvDB struct {
	dbs  map[string]*gorm.DB
	list []*gorm.DB
}

func (s *SrvDB) Start(context.Context) error {
	conf := config.C()

	hasDefault := false
	for key, c := range conf.Postgres {
		if key == DEFAULT_NAME {
			hasDefault = true
		}
		dbInst, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				TablePrefix:   getSearchPath(c.DSN),
			},
			Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false,
				Colorful:                  true,
			}),
			CreateBatchSize: 500,
		})
		if err != nil {
			return err
		}
		var (
			sources  []gorm.Dialector
			replicas []gorm.Dialector
		)
		for _, source := range c.Sources {
			sources = append(sources, postgres.Open(source))
		}
		for _, replica := range c.Replicas {
			replicas = append(replicas, postgres.Open(replica))
		}

		if len(sources) > 0 || len(replicas) > 0 {
			dbInst.Use(dbresolver.Register(
				dbresolver.Config{
					Sources:  sources,
					Replicas: replicas,
				},
			))
		}
		s.dbs[key] = dbInst
		s.list = append(s.list, dbInst)
	}
	if !hasDefault && len(s.list) > 0 {
		s.dbs[DEFAULT_NAME] = s.list[0]
	}
	return nil
}

func (s *SrvDB) Stop(ctx context.Context) error {
	var errors []error
	for _, d := range s.list {
		if d != nil {
			db, err := d.DB()
			if err != nil {
				return err
			}
			err = db.Close()
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

func getSearchPath(dsn string) string {
	re := regexp.MustCompile(`search_path=([^ ]+)`)

	match := re.FindStringSubmatch(dsn)
	if len(match) > 1 {
		return match[1] + "."
	}

	return ""
}
