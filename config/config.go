package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

// DBConn gorm db handler
var DBConn *gorm.DB

var SingleViper *viper.Viper

func init() {
	SingleViper = viper.New()
	SingleViper.SetConfigFile("config/config.toml")
	err := SingleViper.ReadInConfig()
	if err != nil {
		log.Fatalf("[init] read config.toml error: %s\n", err)
	}
}

func GetHttpListenAddress() string {
	return SingleViper.GetString("http.address")
}

type LoggerConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func GetLogger() *LoggerConfig {
	lc := &LoggerConfig{
		FileName:   SingleViper.GetString("logger.FileName"),
		MaxSize:    SingleViper.GetInt("logger.MaxSize"),
		MaxBackups: SingleViper.GetInt("logger.MaxBackups"),
		MaxAge:     SingleViper.GetInt("logger.MaxAge"),
		Compress:   SingleViper.GetBool("logger.Compress"),
	}
	return lc
}
