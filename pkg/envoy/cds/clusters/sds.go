package clusters

import (
	"github.com/deislabs/smc/pkg/envoy"
	xds "github.com/envoyproxy/go-control-plane/envoy/api/v2"
)

func GetSDS() *xds.Cluster {
	return &xds.Cluster{
		ConnectTimeout:       envoy.GetTimeout(),
		ClusterDiscoveryType: &xds.Cluster_Type{Type: xds.Cluster_LOGICAL_DNS},
		Name:                 envoy.SDSClusterName,
		Http2ProtocolOptions: envoy.GetH2(),
		TransportSocket:      envoy.GetTransportSocket(),
		LoadAssignment:       envoy.GetLoadAssignment(envoy.SDSClusterName, envoy.SDSAddress, envoy.SDSPort),
	}
}