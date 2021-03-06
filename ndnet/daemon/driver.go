
package daemon

import (
	log "github.com/sirupsen/logrus"
	"sync"
	dn "github.com/docker/go-plugins-helpers/network"
	"github.com/Nexenta/nedge-docker-network/ndnet/ndnetapi"
	"os/exec"
	"os"
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
	log.Debug(DN, "GetCapabilities req ")
	capabilities := &dn.CapabilitiesResponse{
		Scope: d.scope,
	}
	return capabilities, nil
}

func (d NdnetDriver) CreateNetwork(req *dn.CreateNetworkRequest) error {
	log.Debug(DN, "CreateNetwork req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) DeleteNetwork(req *dn.DeleteNetworkRequest) error {
	log.Debug(DN, "DeleteNetwork req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) CreateEndpoint(req *dn.CreateEndpointRequest) (*dn.CreateEndpointResponse, error) {
	log.Debug(DN, "CreateEndpoint req:\n%+v\n", req)

/*
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
*/
	resp := &dn.CreateEndpointResponse{}
	return resp, nil
}

func (d NdnetDriver) DeleteEndpoint(req *dn.DeleteEndpointRequest) error {
	log.Debug(DN, "DeleteEndpoint req:\n%+v\n", req)

	return nil
}

func (d NdnetDriver) EndpointInfo(req *dn.InfoRequest) (*dn.InfoResponse, error) {
	log.Debug(DN, "EndpointInfo req:\n%+v\n", req)
	// empty map
	value := make(map[string]string)
	resp := &dn.InfoResponse{
		Value: value,
	}
	return resp, nil
}

func (d NdnetDriver) Join(req *dn.JoinRequest) (*dn.JoinResponse, error) {
	log.Debug(DN, "Join req:\n%+v\n", req)

	args := "/opt/nedge/src/nmf/nedocker"
	if _, err := os.Stat(args); err != nil {
		args = "/opt/nedge/nmf/nedocker"
	}
	args = args + " ifup-ndnet " + req.EndpointID
	log.Debug(args)
	go exec.Command("/bin/sh", "-c", args).CombinedOutput()

	resp := &dn.JoinResponse{}
	return resp, nil
}

func (d NdnetDriver) Leave(req *dn.LeaveRequest) error {
	log.Debug(DN, "Leave req:\n%+v\n", req)

	args := "/opt/nedge/src/nmf/nedocker"
	if _, err := os.Stat(args); err != nil {
		args = "/opt/nedge/nmf/nedocker"
	}
	args = args + " ifdown-ndnet " + req.EndpointID
	log.Debug(args)
	go exec.Command("/bin/sh", "-c", args).CombinedOutput()

	return nil
}

func (d NdnetDriver) DiscoverNew(req *dn.DiscoveryNotification) error {
	log.Debug(DN, "DiscoverNew req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Debug(DN, "DiscoveryType: NodeDiscovery\n")
	}
	// Do nothing for now
	return nil
}

func (d NdnetDriver) DiscoverDelete(req *dn.DiscoveryNotification) error {
	log.Debug(DN, "DiscoverDelete req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Debug(DN, "DiscoveryType: NodeDiscovery\n")
	}
	// Do nothing for now
	return nil
}

func (d NdnetDriver) ProgramExternalConnectivity(req *dn.ProgramExternalConnectivityRequest) error {
	log.Debug(DN, "ProgramExternalConnectivity req:\n%+v\n", req)
	// Do nothing for now
	return nil
}

func (d NdnetDriver) RevokeExternalConnectivity(req *dn.RevokeExternalConnectivityRequest) error {
	log.Debug(DN, "RevokeExternalConnectivity req:\n%+v\n", req)
	return nil
}

func (d NdnetDriver) AllocateNetwork (r *dn.AllocateNetworkRequest) (*dn.AllocateNetworkResponse, error) {
       log.Debugf("AllocateNetwork Called: [ %+v ]", r)
       return nil, nil
}

func (d NdnetDriver) FreeNetwork(r *dn.FreeNetworkRequest) error {
       log.Debugf("FreeNetwork Called: [ %+v ]", r)
       return nil
}
