package main

import (
	"fmt"
	"math/big"

	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func main() {
	// Print header
	fmt.Printf("%64s %34s %34s\n", "Private", "Public", "Public Compressed")

	// Initialise big numbers with small numbers
	count, one := big.NewInt(0), big.NewInt(1)

	// Create a slice to pad our count to 32 bytes
	padded := make([]byte, 32)

	// Loop forever because we're never going to hit the end anyway
	for {
		// Increment our counter
		count.Add(count, one)

		// Copy count value's bytes to padded slice
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		// Get public key
		_, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)

		// Get compressed and uncompressed addresses
		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		// Print keys
		fmt.Printf("%x %34s %34s\n", padded, uaddr.EncodeAddress(), caddr.EncodeAddress())
	}
}
