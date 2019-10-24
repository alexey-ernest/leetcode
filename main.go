package main

import "fmt"

func main() {
	// minint := ^int(^uint(0) >> 1)
	maxint := int(^uint(0) >> 1)
	fmt.Printf("%d\n", maxint)
}
