# ttws
Tiny test web server 

This project provides building blocks for creating tiny web servers suitable for running inside containers. It also includes a few sample web servers.

[![Go Report Card](https://goreportcard.com/badge/github.com/rajch/ttws)](https://goreportcard.com/report/github.com/rajch/ttws)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/rajch/ttws)](https://pkg.go.dev/github.com/rajch/ttws)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/rajch/ttws?include_prereleases)

## Packages

|Package|Description|
|-------|-----------|
|webserver|The core webserver. Listens on a port, stops on SIGINT or SIGTERM, sets up handlers defined by the other packages.|
|cpuload|Adds a handler that calculates the square root of 0.0001, one million times, and emits the result.|
|ipaddresses|Adds a handler that emits the host name and ip addresses of the web host.|
|envvars|Adds a handler that emits all environment variables of the host process.|
|filesystem|Adds a handler that emits directory and file names from the host filesystem. Path and depth can be specified.|
|probes|Allows adding _probes_ , which are handlers that can be configured to fail after a specified number of calls, and recover after another specified number.|
|static|Allows adding handlers that serve directories on the web host statically.| 


## Web Servers

|Server|Description|
|---|---|
|ics|A web server which includes the ipaddresses, envvars and filesystem packages. Ipaddresses is default.|
|ldgen|A web server which includes only the cpuload package, which is default.|
|probestest|A web server which includes only the probes package, which is default. Two probes are available on the endpoints '/probes/liveness' and '/probes/readiness'.|
|ttws|A web server which includes all packages. The static package is the default. It serves a directory 'www' under the working directory on the endpoint '/'. Two probes are available on the endpoints '/probes/liveness' and '/probes/readiness'.|

## Options

All servers can have the following options specified either on the command line,  or via an environment 
variable:

|Option|Description|Env Variable|
|---|---|---|
|-p &lt;port>|The port on which the server listens.|PORT|

The probestest and ttws servers allow the following command-line options:

|Option|Description|Env Variable|
|---|---|---|
|-livenessfailafter &lt;calls>|The number of calls after which the liveness probe fails.|LIVENESS_FAIL_AFTER|
|-livenessrecoverafter &lt;calls>|The number of calls post failure after which the liveness probe recovers.|LIVENESS_RECOVER_AFTER|
|-readinessfailafter &lt;calls>|The number of calls after which the readiness probe fails.|READINESS_FAIL_AFTER|
|-readinessrecoverafter &lt;calls>|The number of calls post failure after which the readiness probe recovers.|READINESS_RECOVER_AFTER|
