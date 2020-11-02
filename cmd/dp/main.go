package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dmitsh/device-plugin-example/pkg/dp"

	log "github.com/sirupsen/logrus"
)

func main() {
	c := make(chan os.Signal, 5)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT)

	var p *dp.Plugin
	reset := true
	for {
		if reset {
			if p != nil {
				p.Stop()
			}
			p = dp.NewPlugin()

			if err := p.Run(); err != nil {
				log.Errorf("command returned err: %v", err)
				os.Exit(1)
			}
			reset = false
		}

		select {
		case s := <-c:
			switch s {
			case syscall.SIGHUP:
				log.Infof("Received SIGHUP signal, reseting...")
				reset = true
			default:
				log.Infof("received %v signal, stopping...", s)
				p.Stop()
				os.Exit(1)
			}
		}
	}
}
