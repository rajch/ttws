# ttws
Tiny test web server 

This project provides building blocks for creating tiny web servers suitable for running inside containers.

[![Go Report Card](https://goreportcard.com/badge/github.com/rajch/ttws)](https://goreportcard.com/report/github.com/rajch/ttws)


## Packages

|Package|Description|
|-------|-----------|
|webserver|The core webserver. Listens on a port, stops on SIGINT or SIGTERM, sets up routes defined by the other packages|
|cpuloadgenerator|Registers a route and handler which calculates the square root of 0.0001, one million times.|

## Commands

|Command|Description|
|---|---|
|ttws|A web server which includes all packages|


