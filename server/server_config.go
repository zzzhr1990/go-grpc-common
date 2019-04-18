package server

//Server entity
type Server struct {
	Port          int
	Endpoint      string
	AutoDiscovery struct {
		DiscoveryType string
		Template      string
		ServerAddress string
	}
}
