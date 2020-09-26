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
	webserver.AddHandler(DefaultPath, probesstatushandler)

	webserver.AddInitFunc(func() {
		for _, probe := range probes {
			probe.parseflags()
		}
	})
}
