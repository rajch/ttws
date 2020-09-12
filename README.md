# ttws
Tiny test web server 

This project provides building blocks for creating tiny web servers suitable for running inside containers. It also includes a few sample web servers.

[![Go Report Card](https://goreportcard.com/badge/github.com/rajch/ttws)](https://goreportcard.com/report/github.com/rajch/ttws)


## Packages

|Package|Description|
|-------|-----------|
|webserver|The core webserver. Listens on a port, stops on SIGINT or SIGTERM, sets up handlers defined by the other packages.|
|cpuload|Handler. Calculates the square root of 0.0001, one million times, and outputs the result.|
|ipaddresses|Handler. Outputs the host name and ip addresses of the web host.|
|envvars|Handler. Outputs all environment variables of the host process.|
|filesystem|Handler. Outputs directory and file names from the host filesystem. Path and depth can be specified.|
|probes|Handler. Provides endpoints for probing liveness and readiness. Each can be configured to fail 
after a specified number of calls, and recover after another specified number.| 


## Web Servers

|Server|Description|
|---|---|
|ttws|A web server which includes all packages|
|ics|A web server which includes the ipaddresses, envvars and filesystem packages. Ipaddresses is default.|
|ldgen|A web server which includes only the cpuload package, which is default.|
|probestest|A web server which includes only the probes package, which is default.|

## Options

All servers can have the following options specified either on the command line,  or via an environment 
variable:

|Option|Description|Env Variable|
|---|---|---|
|-p &lt;port>|The port on which the server listens.|PORT|

The probestest server provides the following command-line options:

|Option|Description|Env Variable|
|---|---|---|
|-livenessfailafter &lt;calls>|The number of calls after which the liveness probe fails.|LIVENESS_FAIL_AFTER|
|-livenessrecoverafter &lt;calls>|The number of calls post failure after which the liveness probe recovers.|LIVENESS_RECOVER_AFTER|
|-readinessfailafter &lt;calls>|The number of calls after which the readiness probe fails.|READINESS_FAIL_AFTER|
|-readinessrecoverafter &lt;calls>|The number of calls post failure after which the readiness probe recovers.|READINESS_RECOVER_AFTER|
