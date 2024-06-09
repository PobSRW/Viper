package config

type Configuration struct {
	Environment string
	Server      ConfigurationServer
}

type ConfigurationServer struct {
	Host string
	Port int
}
