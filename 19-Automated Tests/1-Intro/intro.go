package main

import (
	"fmt"
	"intro-tests/addresses"
)

func main() {
	addressType := addresses.AddressType("Avenue")
	fmt.Println(addressType)
}