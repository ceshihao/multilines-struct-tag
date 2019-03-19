# Multilines Struct Tag
Golang struct supports multilines tags

[![Build Status](https://travis-ci.org/ceshihao/multilines-struct-tag.svg?branch=master)](https://travis-ci.org/ceshihao/multilines-struct-tag)
[![Go Report Card](https://goreportcard.com/badge/github.com/ceshihao/multilines-struct-tag)](https://goreportcard.com/report/github.com/ceshihao/multilines-struct-tag)
[![GoDoc](https://godoc.org/github.com/ceshihao/multilines-struct-tag?status.svg)](https://godoc.org/github.com/ceshihao/multilines-struct-tag)
[![Coverage Status](https://coveralls.io/repos/github/ceshihao/multilines-struct-tag/badge.svg?branch=master)](https://coveralls.io/github/ceshihao/multilines-struct-tag?branch=master)

## Background
There was a [propose](https://github.com/golang/go/issues/15893) to accept multilines struct tag in golang, but rejected.

This package is based on Golang standard library.

https://golang.org/pkg/reflect/#StructTag.Get

https://golang.org/pkg/reflect/#StructTag.Lookup

## Notice
Multilines struct tag violates the rule in `go vet`.

