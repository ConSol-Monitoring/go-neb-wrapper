go-naemon-broker-module
=======================

[![Go Report Card](https://goreportcard.com/badge/github.com/ConSol/go-neb-wrapper)](https://goreportcard.com/report/github.com/ConSol/go-neb-wrapper)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)


What is this
============

This is an NEB module API for [Naemon](http://www.naemon.org)  / Nagios written in Golang. This allows you to write NEB modules in Golang, which is otherwise only possible in C / C++.

It is using CGo therefor this is not plain Go, there you have to be careful, if you want to write a module.

# Limitations

This API works with:
- Naemon
- Nagios 4

Nagios 3 and Icinga 1 are only working **without** the daemon mode!

# Minimalistic-Example
[main.go](https://github.com/ConSol/go-neb-wrapper/blob/master/main.go)

# Modules written with this API

- [Iapetos](https://github.com/Griesbacher/Iapetos)

If you have written a module yourself, let us know.

# Build

The module has to be compiled as shared library so Nagios can load it.

Naemon:
- go build -tags naemon -buildmode=c-shared -ldflags "-s -w"

Nagios 3:
- go build -tags nagios3 -buildmode=c-shared -ldflags "-s -w"

Nagios 4:
- go build -tags nagios4 -buildmode=c-shared -ldflags "-s -w"
