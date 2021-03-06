package main

import (
	"github.com/rajch/ttws/pkg/probes"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	// TTWS includes "readiness" and "liveness" probes
	probes.NewProbe("Readiness", 10, 0)
	probes.NewProbe("Liveness", 10, 0)

	webserver.SetRootHandler(probes.DefaultPath)
	webserver.ListenAndServe()
}
