package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var runtimeConfig *CPRuntimeConfig

const (
	EnvCPCoreConfigDir     = "CP_CORE_CONFIG_DIR"
	DefaultCPCoreConfigDir = "/Users/bytedance/Desktop/code/learn/chat-paper-github/cp-core/conf/"
)

type MySQLConfig struct {
	MySQLUsername string `mapstructure:"CP_MYSQL_USERNAME"`
	MySQLPassword string `mapstructure:"CP_MYSQL_PASSWORD"`
	MySQLDatabase string `mapstructure:"CP_MYSQL_DATABASE"`
	MySQLHost     string `mapstructure:"CP_MYSQL_HOST"`
	MySQLPort     string `mapstructure:"CP_MYSQL_PORT"`
}

type CPRuntimeConfig struct {
	MySQLConfig MySQLConfig `mapstructure:"CP_MYSQL_CONFIG"`
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetRuntimeConfig() *CPRuntimeConfig {
	v := viper.New()
	confDir := getEnv(EnvCPCoreConfigDir, DefaultCPCoreConfigDir)
	v.SetConfigName("conf.local")
	v.SetConfigType("yml")
	v.AddConfigPath(confDir)

	// important! do not miss this step
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	runtimeConfig = &CPRuntimeConfig{}
	err := v.Unmarshal(runtimeConfig)
	if err != nil {
		panic(err)
	}
	return runtimeConfig
}
