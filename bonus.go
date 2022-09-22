package main

import (
	"fmt"
)

func main() {
	output_hash := CalculateHash("hello")
	fmt.Println("Hash of the block data: ", output_hash)
}
