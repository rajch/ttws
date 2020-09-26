// Package cpuload registers a handler that calculates the square root of 0.0001, one million times, and emits the result.
package cpuload

import (
	"fmt"
	"math"
	"net/http"

	"github.com/rajch/ttws/pkg/webserver"
)

// DefaultPath is the default path to invoke the cpuloadgenerator module - /loadcpu
const DefaultPath = "/loadcpu"

func cpuloadhandler(w http.ResponseWriter, r *http.Request) {
	value := 0.0001
	for i := 0; i <= 1000000; i++ {
		value += math.Sqrt(value)
	}
	fmt.Fprintf(w, "OK:%v!\n", value)
}

func init() {
	webserver.AddHandler(DefaultPath, cpuloadhandler)
}
