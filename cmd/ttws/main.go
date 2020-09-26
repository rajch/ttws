package main

import (
	_ "github.com/rajch/ttws/pkg/cpuload"
	_ "github.com/rajch/ttws/pkg/envvars"
	_ "github.com/rajch/ttws/pkg/filesystem"
	_ "github.com/rajch/ttws/pkg/ipaddresses"
	"github.com/rajch/ttws/pkg/probes"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	probes.NewProbe("Readiness", 10, 0)
	probes.NewProbe("Liveness", 10, 0)
	webserver.ListenAndServe()
}
