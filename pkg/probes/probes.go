package probes

import (
	"fmt"
	"net/http"

	"github.com/rajch/ttws/pkg/webserver"
)

var probes = map[string]*probe{}

// DefaultPath is the default path to invoke the probes module - /probes
const DefaultPath = "/probes"

func probesstatushandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PROBE STATUS")

	for _, probe := range probes {
		fmt.Fprintln(w, probe.status())
	}
}

func init() {

	NewProbe("Readiness", 10, 0)
	NewProbe("Liveness", 10, 0)

	webserver.AddHandler(DefaultPath, probesstatushandler)

	for _, probe := range probes {
		probe.parseflags()
	}
}
