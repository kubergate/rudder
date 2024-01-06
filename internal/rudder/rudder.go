package rudder

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	v1alpha1 "github.com/kubergate/rudder/api/v1alpha1/config"
	"github.com/kubergate/rudder/internal/controllers"
	"github.com/kubergate/rudder/internal/server"
	xds "github.com/kubergate/rudder/internal/xds/server"
)

func Init(config v1alpha1.Rudder) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	ctx := context.Background()

	go controllers.InitControllers(config.Spec.DataStoreConfig)
	go server.InitServer()
	go xds.InitXdsServer(ctx, config.Spec.XdsServerConfig)

	for s := range sig {
		switch s {
		case os.Interrupt:
			return
		case syscall.SIGTERM:
			return
		}
	}
}
