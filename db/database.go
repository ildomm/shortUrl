package db

import (
	"log"
	"sync"

	"github.com/gocraft/dbr"
)

var conn *pocDATABASES
var once sync.Once
type pocDATABASES struct {
	Mysql *dbr.Connection
}

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

func Setup() *pocDATABASES {
	once.Do(func() {
		conn = instances()
	})
	return conn
}

func instances() *pocDATABASES {
	log.Println("Connecting databases" )

	return &pocDATABASES{
		Mysql: mysql(),
	}
}

func CleanDatabases() {
	CleanMysqlDB()
}