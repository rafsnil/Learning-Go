package main

import "fmt"

type ByteSize int

const (
	// _  = iota
	// KB = 1024 * iota
	// MB = KB * iota
	// GB = MB * iota
	// TB = GB * iota
	// PB = TB * iota
	// EB = PB * iota

	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {

	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)

}
