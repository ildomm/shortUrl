package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/ildomm/linx_challenge/config"
	"log"
)

/* Return Mysql session instance */
func Mysql() *dbr.Session {
	return conn.Mysql.NewSession(nil)
}
func mysql() *dbr.Connection {
	csqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.App.Database.Username,
		config.App.Database.Password,
		config.App.Database.Host,
		config.App.Database.Port,
		config.App.Database.Database)

	log.Println("Mysql: " + csqlInfo)

	/* See: https://github.com/gocraft/dbr
	wiki -> https://godoc.org/github.com/gocraft/dbr
	*/
	conn, _ := dbr.Open("mysql", csqlInfo, nil)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(4)

	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return conn
}

func ResetMysqlDB() bool {
	return true
}

func CleanMysqlDB() bool {
	return true
}
