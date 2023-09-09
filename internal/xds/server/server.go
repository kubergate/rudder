package server

import (
	"context"
	"net"
	"strconv"

	"github.com/KommodoreX/dp-rudder/internal/xds/cache"
	"github.com/KommodoreX/dp-rudder/pkg/logger"
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/service/cluster/v3"
	discoveryv3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/service/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/service/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/service/route/v3"
	runtimev3 "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	secretv3 "github.com/envoyproxy/go-control-plane/envoy/service/secret/v3"
	serverv3 "github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"google.golang.org/grpc"
)

const XdsServerAddress = "0.0.0.0"

type XdsServer struct {
	grpc  *grpc.Server
	cache cache.SnapshotCacheWithCallbacks
}

// func InitXdsServer(config) error {

// }

func (xds *XdsServer) Run(ctx context.Context) error {
	xds.grpc = grpc.NewServer()
	xds.cache = cache.NewSnapshotCache(true, logger.LoggerDragonFly)
	registerServer(serverv3.NewServer(ctx, xds.cache, xds.cache), xds.grpc)
	go xds.serveXdsServer(ctx)
	return nil
}

// registerServer registers the given xDS protocol Server with the gRPC
// runtime.
func registerServer(srv serverv3.Server, g *grpc.Server) {
	// register services
	discoveryv3.RegisterAggregatedDiscoveryServiceServer(g, srv)
	secretv3.RegisterSecretDiscoveryServiceServer(g, srv)
	clusterv3.RegisterClusterDiscoveryServiceServer(g, srv)
	endpointv3.RegisterEndpointDiscoveryServiceServer(g, srv)
	listenerv3.RegisterListenerDiscoveryServiceServer(g, srv)
	routev3.RegisterRouteDiscoveryServiceServer(g, srv)
	runtimev3.RegisterRuntimeDiscoveryServiceServer(g, srv)
}

func (xds *XdsServer) serveXdsServer(ctx context.Context) {
	addr := net.JoinHostPort(XdsServerAddress, strconv.Itoa(9200))
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.LoggerDragonFly.Sugar().Error(err, "failed to listen on address", "address", addr)
		return
	}
	err = xds.grpc.Serve(l)
	if err != nil {
		logger.LoggerDragonFly.Sugar().Error(err, "failed to start grpc based xds server")
	}

	<-ctx.Done()
	logger.LoggerDragonFly.Sugar().Info("grpc server shutting down")
	// We don't use GracefulStop here because envoy
	// has long-lived hanging xDS requests. There's no
	// mechanism to make those pending requests fail,
	// so we forcibly terminate the TCP sessions.
	xds.grpc.Stop()
}
