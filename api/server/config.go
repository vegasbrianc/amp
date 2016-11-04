package server

import (
	"fmt"
)

// Config is used for amplifier configuration settings
type Config struct {
	ServerAddress    string
	ServerPort       string
	EtcdEndpoints    []string
	ElasticsearchURL string
	ClientID         string
	ClientSecret     string
	InfluxURL        string
	DockerURL        string
	DockerVersion    string
	NatsURL          string
}

// String is used to display struct as a string
func (config Config) String() string {
	return fmt.Sprintf("{ Server: %s:%s, EtcdEndpoints: %v, ElasticsearchURL: %s, NatsURL: %s, InfluxURL: %s, Docker: %s}", config.ServerAddress, config.ServerPort, config.EtcdEndpoints, config.ElasticsearchURL, config.NatsURL, config.InfluxURL, config.DockerURL)
}
