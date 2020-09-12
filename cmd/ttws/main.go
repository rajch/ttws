package main

import (
	_ "github.com/rajch/ttws/pkg/cpuload"
	_ "github.com/rajch/ttws/pkg/envvars"
	_ "github.com/rajch/ttws/pkg/filesystem"
	_ "github.com/rajch/ttws/pkg/ipaddresses"
	_ "github.com/rajch/ttws/pkg/probes"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	webserver.ListenAndServe()
}
