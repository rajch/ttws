package main

import (
	_ "github.com/rajch/ttws/pkg/envvars"
	_ "github.com/rajch/ttws/pkg/filesystem"
	"github.com/rajch/ttws/pkg/ipaddresses"
	"github.com/rajch/ttws/pkg/webserver"
)

func main() {
	webserver.SetRootHandler(ipaddresses.DefaultPath)
	webserver.ListenAndServe()
}
