#!/bin/bash -x
# By: Jens Rantil
# Fm: https://gist.github.com/JensRantil/5073646

#go run SignalSrc.go &
# -- NOK. Catching signals does not properly work with `go run ...`.
go build SignalSrc.go
./SignalSrc &
pid=$!
sleep 2; kill -INT $pid
# Prints the "interrupt" and quits after 5 seconds, which is to be expected.
