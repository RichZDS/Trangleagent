package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	names := []string{"admin", "Bradtke1974284", "Denesik4162865", "DuBuque1671714", "Flatley9017515"}
	for _, n := range names {
		h := crc32.ChecksumIEEE([]byte(n))
		fmt.Printf("%s: %d (%%2=%d)\n", n, h, h%2)
	}
}
