package cpuload

import (
	"fmt"
	"math"
	"net/http"

	"github.com/rajch/ttws/pkg/webserver"
)

// Route is the default path to invoke the cpuloadgenerator module - /loadcpu
const Route = "/loadcpu"

func cpuloadhandler(w http.ResponseWriter, r *http.Request) {
	value := 0.0001
	for i := 0; i <= 1000000; i++ {
		value += math.Sqrt(value)
	}
	fmt.Fprintf(w, "OK:%v!\n", value)
}

// SetHome sets the default route of this module as the web server default
func SetHome() {
	webserver.SetHome(Route)
}

func init() {
	webserver.AddRoute(Route, cpuloadhandler)
}