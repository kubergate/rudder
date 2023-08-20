package dragonfly

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/NomadXD/dragonfly/internal/operator"
	"github.com/NomadXD/dragonfly/internal/server"
)

func Init() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
 
	go operator.InitOperator()
	go server.InitServer()

	for s := range sig {
		switch s {
		case os.Interrupt:
			return
		case syscall.SIGTERM:
			return
		}
	}
}