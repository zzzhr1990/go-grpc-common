package config

import (
	"github.com/zzzhr1990/go-grpc-common/client"
	"github.com/zzzhr1990/go-grpc-common/server"
)

//Grpc common config
type Grpc struct {
	Server server.Server
	Client client.Client
}
