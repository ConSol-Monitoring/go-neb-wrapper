go-naemon-broker-module
=======================

[![Go Report Card](https://goreportcard.com/badge/github.com/ConSol-Monitoring/go-neb-wrapper)](https://goreportcard.com/report/github.com/ConSol-Monitoring/go-neb-wrapper)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)


What is this
============

This is an NEB module API for [Naemon](http://www.naemon.io) written in Golang. This allows you to write NEB modules in Golang, which is otherwise only possible in C / C++.

It is using CGo therefor this is not plain Go, there you have to be careful, if you want to write a module.

# Minimalistic-Example
[main.go](https://github.com/ConSol-Monitoring/go-neb-wrapper/blob/master/main.go)

# Modules written with this API

- [Naemon Vault Example](https://github.com/sni/naemon-vault-example)
- [Iapetos](https://github.com/Griesbacher/Iapetos)

If you have written a module yourself, let us know.

# Build

The module has to be compiled as shared library so Naemon can load it.

You will have to install the `naemon-dev` package in order to find all required c header files.

Naemon:
- go build -buildmode=c-shared -ldflags "-s -w" -o go-neb-wrapper-example.so
