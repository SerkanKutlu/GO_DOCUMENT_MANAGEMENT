package config

type ApplicationConfig struct {
	Mongo MongoConfig `yaml:"mongo"`
}

type MongoConfig struct {
	ConnectionString string            `yaml:"connectionString"`
	Database         string            `yaml:"database"`
	Collection       map[string]string `yaml:"collection"`
}
