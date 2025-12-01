package db

import "sync"

var (
	toMigrate   = make(map[string][]interface{}, 0)
	migrateLock sync.Mutex
)

func AutoMigrate(dbName string, dst ...interface{}) error {
	if dbName == "" {
		dbName = DEFAULT_NAME
	}
	migrateLock.Lock()
	defer migrateLock.Unlock()
	toMigrate[dbName] = append(toMigrate[dbName], dst...)
	return nil
}
