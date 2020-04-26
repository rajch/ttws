package main

import (
	_ "github.com/rajch/ttws/pkg/cpuloadgenerator"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	webserver.Serve()
}
