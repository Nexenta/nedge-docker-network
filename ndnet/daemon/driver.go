
package daemon

import (
	log "github.com/Sirupsen/logrus"
	"sync"
	dn "github.com/docker/go-plugins-helpers/network"
	"github.com/Nexenta/nedge-docker-network/ndnet/ndnetapi"
)

var (
	DN = "Ndnet "
)

type NdnetDriver struct {
	scope		string
	Client         *ndnetapi.Client
	Mutex          *sync.Mutex
}

func DriverAlloc(cfgFile string) (NdnetDriver, error) {

	client, _ := ndnetapi.ClientAlloc(cfgFile)

	d := NdnetDriver{
		scope:		"local",
		Client:         client,
		Mutex:          &sync.Mutex{},
	}

	return d, nil
}

func (d NdnetDriver) GetCapabilities() (*dn.CapabilitiesResponse,
	error) {
	log.Debug(DN, "Received GetCapabilities req")
	capabilities := &dn.CapabilitiesResponse{
		Scope: d.scope,
	}
	return capabilities, nil
}

func (d NdnetDriver) CreateNetwork(req *dn.CreateNetworkRequest) error {
	log.Debug(DN, "Received CreateNetwork req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) DeleteNetwork(req *dn.DeleteNetworkRequest) error {
	log.Debug(DN, "Received DeleteNetwork req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) CreateEndpoint(req *dn.CreateEndpointRequest) (*dn.CreateEndpointResponse, error) {
	log.Debug(DN, "Received CreateEndpoint req:\n%+v\n", req)

	iface := new(dn.EndpointInterface)

	if req.Interface == nil {
		// TODO: Revise these.
		iface.Address = "1.1.1.1/24"
		iface.AddressIPv6 = "fd00::0000:0000:0000:0000/64"
		iface.MacAddress = "DE:AD:BE:EF:00:00"
	}
	resp := &dn.CreateEndpointResponse{
		Interface: iface,
	}
	return resp, nil
}

func (d NdnetDriver) DeleteEndpoint(req *dn.DeleteEndpointRequest) error {
	log.Debug(DN, "Received DeleteEndpoint req:\n%+v\n", req)
	return nil
}

func (d NdnetDriver) EndpointInfo(req *dn.InfoRequest) (*dn.InfoResponse, error) {
	log.Debug(DN, "Received EndpointOperInfo req:\n%+v\n", req)
	// empty map
	value := make(map[string]string)
	resp := &dn.InfoResponse{
		Value: value,
	}
	return resp, nil
}

func (d NdnetDriver) Join(req *dn.JoinRequest) (*dn.JoinResponse, error) {
	log.Debug(DN, "Received Join req:\n%+v\n", req)

	resp := &dn.JoinResponse{
		InterfaceName: dn.InterfaceName{
			SrcName:   "nedge_rep",
			DstPrefix: "rep",
		},
		DisableGatewayService: false,
	}
	return resp, nil
}

func (d NdnetDriver) Leave(req *dn.LeaveRequest) error {
	log.Debug(DN, "Received Leave req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) DiscoverNew(req *dn.DiscoveryNotification) error {
	log.Debug(DN, "Received DiscoverNew req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Debug(DN, "DiscoveryType: NodeDiscovery\n")
	}
	// Do nothing for now
	return nil
}

func (d NdnetDriver) DiscoverDelete(req *dn.DiscoveryNotification) error {
	log.Debug(DN, "Received DiscoverDelete req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Debug(DN, "DiscoveryType: NodeDiscovery\n")
	}
	// Do nothing for now
	return nil
}

func (d NdnetDriver) ProgramExternalConnectivity(req *dn.ProgramExternalConnectivityRequest) error {
	log.Debug(DN, "Received ProgramExternalConnectivity req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) RevokeExternalConnectivity(req *dn.RevokeExternalConnectivityRequest) error {
	log.Debug(DN, "Received RevokeExternalConnectivity req:\n%+v\n", req)
	return nil
}

