// https://pkg.go.dev/hash/crc32#example-MakeTable

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// In this package, the CRC polynomial is represented in reversed notation,
	// or LSB-first representation.
	//
	// LSB-first representation is a hexadecimal number with n bits, in which the
	// most significant bit represents the coefficient of x⁰ and the least significant
	// bit represents the coefficient of xⁿ⁻¹ (the coefficient for xⁿ is implicit).
	//
	// For example, CRC32-Q, as defined by the following polynomial,
	//	x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
	// has the reversed notation 0b11010101100000101000001010000001, so the value
	// that should be passed to MakeTable is 0xD5828281.
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32q))

	crc32c := crc32.MakeTable(crc32.Castagnoli) // 0x82f63b78
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32c))

	crc32k := crc32.MakeTable(crc32.Koopman) // 0xeb31d82e
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32k))
}

/*

$ go run crc32.go
2964d064
72b51f78
b1bcb065

*/
