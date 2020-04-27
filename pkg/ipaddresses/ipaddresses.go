package ipaddresses

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/rajch/ttws/pkg/webserver"
)

// DefaultPath is the default path to invoke the ipaddresses module - /ip
const DefaultPath = "/ip"

func ipaddresseshandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = err.Error()
	}

	// Print hostname
	fmt.Fprintf(w, "Hostname: %v\n", hostname)

	// Print IP addresses
	fmt.Fprintln(w, "IP Addresses: ")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Fprint(w, "  %v\n", err.Error())
	} else {
		for _, addr := range addrs {
			fmt.Fprintf(w, "  - %v\n", addr.String())
		}
	}
}

func init() {
	webserver.AddHandler(DefaultPath, ipaddresseshandler)
}
