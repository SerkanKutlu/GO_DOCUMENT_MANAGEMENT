package config

type ApplicationConfig struct {
	Mongo  MongoConfig `yaml:"mongo"`
	JwtKey JwtKey      `yaml:"jwtKey"`
}

type MongoConfig struct {
	ConnectionString string            `yaml:"connectionString"`
	Database         string            `yaml:"database"`
	Collection       map[string]string `yaml:"collection"`
}

type JwtKey struct {
	SecretKey string `yaml:"secretKey"`
}
