// Package probes allows the creation of probes, which are handlers that return either success or failure.
// Each probe can be configured to fail (that is, start returning 500 status codes) after a specified
// number of calls, and to recover (start returning 200 status codes) after another specified number.
// Consumers of the package can call the NewProbe method to add any number of probes.
// The package itself registers a handler on the default path, which emits a list of all probes
// with their current status.
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
