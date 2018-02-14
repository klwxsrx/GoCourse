package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"sync"
)

const dataSourceName = "gocourse:1234@/gocourse"
const mysqlDriverName = "mysql"

var instance *sql.DB
var once sync.Once

func GetConnection() *sql.DB {
	once.Do(func() {
		var err error
		instance, err = sql.Open(mysqlDriverName, dataSourceName)
		if err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
