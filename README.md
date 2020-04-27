# ttws
Tiny test web server 

This project provides building blocks for creating tiny web servers suitable for running inside containers. It also includes a few sample web servers.

[![Go Report Card](https://goreportcard.com/badge/github.com/rajch/ttws)](https://goreportcard.com/report/github.com/rajch/ttws)


## Packages

|Package|Description|
|-------|-----------|
|webserver|The core webserver. Listens on a port, stops on SIGINT or SIGTERM, sets up handlers defined by the other packages.|
|cpuloadgenerator|Handler. Calculates the square root of 0.0001, one million times, and outputs the result.|
|ipaddresses|Handler. Outputs the host name and ip addresses of the web host.|
|envvars|Handler. Outputs all environment variables of the host process.|
|filesystem|Handler. Outputs directory and file names from the host filesystem. Path and depth can be specified.|


## Web Servers

|Server|Description|
|---|---|
|ttws|A web server which includes all packages|
|ics|A web server which includes the ipaddresses, envvars and filesystem packages. Ipaddresses is default.|

## Options

All servers can have the following options specified on the command line:

|Option|Description|
|---|---|
|-p &lt;port>|The port on which the server listens. Can also be specified via an evironment variable, PORT.|


