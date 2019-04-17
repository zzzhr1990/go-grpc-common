package client

// Client client entity
type Client struct {
	DefaultEndPoint string
	AutoDiscovery   AutoDiscoveryClient `yaml:"auto-discovery"`
}

// AutoDiscoveryClient AutoDiscoveryClient
type AutoDiscoveryClient struct {
	DiscoveryType string `yaml:"discovery-type"`
	Template      string
	// ManualEndpoints map[string][]string `yaml:"manual-endpoints"`
	ManualEndpoints map[string]string `yaml:"manual-endpoints"`
}
