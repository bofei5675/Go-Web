package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // provider driver for mysql from driver.go
	"github.com/spf13/viper"
	"log"
)

var DB *sql.DB // global variable

func Init() error {
	var err error
	// construct a sql.DB object
	DB, err = sql.Open("mysql", viper.GetString("mysql.source_name"))
	if nil != err {
		return err
	}

	if nil != err {
		return err
	}
	// max timeout
	DB.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	// contruct connection
	err = DB.Ping()
	if nil != err {
		return err
	} else {
		log.Println("MySQL Startup Normal !")
	}
	return nil
}
