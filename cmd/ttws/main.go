package main

import (
	_ "github.com/rajch/ttws/pkg/cpuload"
	_ "github.com/rajch/ttws/pkg/envvars"
	_ "github.com/rajch/ttws/pkg/filesystem"
	_ "github.com/rajch/ttws/pkg/ipaddresses"
	"github.com/rajch/ttws/pkg/probes"
	"github.com/rajch/ttws/pkg/static"
	_ "github.com/rajch/ttws/pkg/static"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	// TTWS includes "readiness" and "liveness" probes
	probes.NewProbe("Readiness", 10, 0)
	probes.NewProbe("Liveness", 10, 0)

	// TTWS serves "www" directory statically
	static.ServeDirectory("/", "./www")

	webserver.ListenAndServe()
}
