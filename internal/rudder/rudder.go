package dragonfly

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/KommodoreX/dp-rudder/internal/controllers"
	"github.com/KommodoreX/dp-rudder/internal/server"
)

func Init() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go controllers.InitControllers()
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
