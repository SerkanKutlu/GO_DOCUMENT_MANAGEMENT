package config

type HttpClientConfig struct {
	UserService HttpClient `yaml:"userService"`
}
type HttpClient struct {
	Name    string `yaml:"name"`
	BaseUrl string `yaml:"baseUrl"`
}
