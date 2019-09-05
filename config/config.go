package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

// initialize log info
func LogInfo() {
	// create a file name as date.log
	file := "./" + time.Now().Format("2006-01-02") + ".log"
	// open this file as a writter
	logFile, _ := os.OpenFile(file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0766)
	// set flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// set outputs
	log.SetOutput(logFile)
}

// Init
func Init() error {
	if err := Config() ; err != nil {
		return err
	}
	LogInfo()
	return nil
}
// parse config file by viper
func Config() error{
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil{
		return err
	}
	return nil
}

