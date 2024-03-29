// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package cake_test

import (
	"flag"
	"testing"
	"time"

	"github.com/suntong/lang/lang/Go/src/async/cake"
)

var defaults cake.Shop

func init() {
	testing.Init()
	flag.Parse()
	defaults = cake.Shop{
		Verbose:      testing.Verbose(),
		Cakes:        20,
		BakeTime:     10 * time.Millisecond,
		NumIcers:     1,
		IceTime:      10 * time.Millisecond,
		InscribeTime: 10 * time.Millisecond,
	}
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber.
	// Each step takes exactly 10ms.  No buffers.
	cakeshop := defaults
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect.
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 224 ms
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	// increases total time due to channel delays.
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.Work(b.N) // 259 ms
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	// delays resulting from variability.
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N) // 244 ms
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	// adds directly to the critical path.
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N) // 1.032 s
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	// to its sequential component, following Amdahl's Law.
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N) // 288ms
}

/*

$ go test -bench=. github.com/suntong/lang/lang/Go/src/async/cake
goos: linux
goarch: amd64
pkg: github.com/suntong/lang/lang/Go/src/async/cake
cpu: Intel(R) Xeon(R) CPU E5-1650 0 @ 3.20GHz
Benchmark-12                                 5   225641032 ns/op
BenchmarkBuffers-12                          5   223561843 ns/op
BenchmarkVariable-12                         4   269953964 ns/op
BenchmarkVariableBuffers-12                  4   250085640 ns/op
BenchmarkSlowIcing-12                        1  1031772493 ns/op
BenchmarkSlowIcingManyIcers-12               4   270255823 ns/op
PASS
ok   github.com/suntong/lang/lang/Go/src/async/cake     11.122s

*/
