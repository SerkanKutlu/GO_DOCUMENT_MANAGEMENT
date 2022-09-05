package config

import (
	"github.com/spf13/viper"
)

type ConfigurationManager struct {
	applicationConfig *ApplicationConfig
}

func NewConfigurationManager(path string, file string, env string) *ConfigurationManager {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	appConfig := readApplicationConfigFile(env, file)
	return &ConfigurationManager{
		applicationConfig: appConfig,
	}
}

func readApplicationConfigFile(env string, file string) *ApplicationConfig {

	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var appConfig ApplicationConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&appConfig); err != nil {
		panic(err.Error())
	}
	return &appConfig
}

func (cm *ConfigurationManager) GetMongoConfiguration() *MongoConfig {
	return &cm.applicationConfig.Mongo
}
