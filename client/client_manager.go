package client

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	// "github.com/zzzhr1990/go-client-common/config"

	// "github.com/zzzhr1990/go-client-common/config"
	"github.com/zzzhr1990/go-common-entity/confs"
	"google.golang.org/grpc"
)

//Manager manage connections
type Manager struct {
	connections  *sync.Map
	clientConfig *confs.GrpcClientConfig
	lock         *sync.RWMutex
}

//CreateNewManager new Manager
func CreateNewManager(clientConfig *confs.GrpcClientConfig) *Manager {
	return &Manager{
		clientConfig: clientConfig,
		lock:         new(sync.RWMutex),
		connections:  &sync.Map{},
	}
}

//GetConnection for connections
func (manager *Manager) GetConnection(serviceLabel string) (*grpc.ClientConn, error) {
	data, ok := manager.connections.Load(serviceLabel)
	if ok {
		conn, suc := data.(*grpc.ClientConn)
		if !suc {
			return nil, errors.New("Convert ClientConn instance Error")
		}
		return conn, nil
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	data, ok = manager.connections.Load(serviceLabel)
	// dounle checking
	if ok {
		conn, suc := data.(*grpc.ClientConn)
		if !suc {
			return nil, errors.New("Convert ClientConn instance Error")
		}
		return conn, nil
	}
	// stil no
	conn, err := manager.createConnection(serviceLabel)
	if err != nil {
		return nil, err
	}
	manager.connections.Store(serviceLabel, conn)
	return conn, nil
}

func (manager *Manager) createConnection(serviceLabel string) (*grpc.ClientConn, error) {
	endpoint := manager.getEndpoints(serviceLabel)
	wscr := grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(128*1024*1024), grpc.MaxCallSendMsgSize(128*1024*1024))
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure(), wscr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (manager *Manager) getEndpoints(serviceLabel string) string {
	endpoint := "localhost:8964"
	if len(manager.clientConfig.DefaultEndpoint) > 0 {
		endpoint = manager.clientConfig.DefaultEndpoint
	}
	preData := map[string]string{"${service-label}": serviceLabel, "${default-endpoint}": endpoint}
	//manager.clientConfig.DefaultEndPoint
	if manager.clientConfig.AutoDiscovery.DiscoveryType == "default-only" {
		return formatTemplate(&preData, endpoint)
	}
	if val, ok := manager.clientConfig.AutoDiscovery.ManualEndpoints[serviceLabel]; ok {
		return formatTemplate(&preData, val)
	}
	if manager.clientConfig.AutoDiscovery.DiscoveryType == "eureka" || manager.clientConfig.AutoDiscovery.DiscoveryType == "etcd" {
		panic(fmt.Sprintf("Cannot support autoDiscoveryType %v now.", manager.clientConfig.AutoDiscovery.DiscoveryType))
		// return formatTemplate(&preData, endpoint)
	}
	autoDiscoveryTemplate := "${default-endpoint}"
	if len(manager.clientConfig.AutoDiscovery.Template) > 0 {
		autoDiscoveryTemplate = manager.clientConfig.AutoDiscovery.Template
	}
	if manager.clientConfig.AutoDiscovery.DiscoveryType == "none" || manager.clientConfig.AutoDiscovery.DiscoveryType == "istio" {
		return formatTemplate(&preData, autoDiscoveryTemplate)
	}
	return endpoint

}

//Close disconnection
func (manager *Manager) Close() {
	manager.connections.Range(closeConnection)
}

func formatTemplate(preData *map[string]string, template string) string {
	for k, v := range *preData {
		//template = template.replace(k, preData[k]);
		template = strings.Replace(template, k, v, -1)
	}
	return template
}

func closeConnection(key interface{}, value interface{}) bool {
	_, ok := key.(string)
	if ok {
		val, suc := value.(*grpc.ClientConn)
		if suc {
			val.Close()

		}
	}
	return true
}
