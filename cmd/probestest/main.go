package main

import (
	"github.com/rajch/ttws/pkg/probes"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	webserver.SetRootHandler(probes.DefaultPath)
	webserver.ListenAndServe()
}
