// Copyright (c) 2020, Marcelo Jorge Vieira
// Licensed under the AGPL-3.0+ License

package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	envReplacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(envReplacer)
	viper.SetEnvPrefix("politicos")
	viper.SetDefault("mongodb.endpoint", "mongodb://localhost:27017")
	viper.SetDefault("mongodb.name", "politicos")
	viper.SetDefault("port", 8888)
	viper.SetDefault("debug", false)
	logLevel := log.InfoLevel
	if viper.GetBool("debug") {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	viper.SetDefault("db.operation.timeout", 30)
	viper.SetDefault("db.operation.per-page", 10)
	viper.SetDefault("collection.operation.timeout", 30)
}
