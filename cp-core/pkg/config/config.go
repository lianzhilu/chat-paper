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

type GatewayConfig struct {
	GateWayHost string `mapstructure:"CP_GATEWAY_HOST"`
	GateWayPort string `mapstructure:"CP_GATEWAY_PORT"`
}

type UserServiceConfig struct {
	UserServiceHost string `mapstructure:"CP_USER_SERVICE_HOST"`
	UserServicePort string `mapstructure:"CP_USER_SERVICE_PORT"`
}

type ArticleServiceConfig struct {
	ArticleServiceHost string `mapstructure:"CP_ARTICLE_SERVICE_HOST"`
	ArticleServicePort string `mapstructure:"CP_ARTICLE_SERVICE_PORT"`
}

type CPRuntimeConfig struct {
	MySQLConfig          MySQLConfig          `mapstructure:"CP_MYSQL_CONFIG"`
	GatewayConfig        GatewayConfig        `mapstructure:"CP_GATEWAY_CONFIG"`
	UserServiceConfig    UserServiceConfig    `mapstructure:"CP_USER_SERVICE_CONFIG"`
	ArticleServiceConfig ArticleServiceConfig `mapstructure:"CP_ARTICLE_SERVICE_CONFIG"`
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
