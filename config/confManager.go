package config

import (
	"github.com/spf13/viper"
)

type ConfigurationManager struct {
	applicationConfig    *ApplicationConfig
	remoteServicesConfig *HttpClientConfig
}

func NewConfigurationManager(path string, file string, env string) *ConfigurationManager {
	viper.AddConfigPath(path)
	viper.SetConfigType("yml")
	appConfig := readApplicationConfigFile(env, file)
	remoteServerConfig := readRemoteServicesConfigFile(env, file)
	return &ConfigurationManager{
		applicationConfig:    appConfig,
		remoteServicesConfig: remoteServerConfig,
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
func readRemoteServicesConfigFile(env string, file string) *HttpClientConfig {
	viper.SetConfigName(file)
	if err := viper.ReadInConfig(); err != nil {
		panic("Can not load application config file")
	}
	var remoteConfig HttpClientConfig
	envSub := viper.Sub(env)
	if err := envSub.Unmarshal(&remoteConfig); err != nil {
		panic(err.Error())
	}
	return &remoteConfig
}

func (cm *ConfigurationManager) GetMongoConfiguration() *MongoConfig {
	return &cm.applicationConfig.Mongo
}
func (cm *ConfigurationManager) GetJwtKey() *JwtKey {
	return &cm.applicationConfig.JwtKey
}
func (cm *ConfigurationManager) GetHttpClientConfig() *HttpClientConfig {
	return cm.remoteServicesConfig
}
