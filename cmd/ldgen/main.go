package main

import (
	"github.com/rajch/ttws/pkg/cpuload"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	webserver.SetRootHandler(cpuload.DefaultPath)
	webserver.ListenAndServe()
}
