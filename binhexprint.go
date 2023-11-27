package main

import "fmt"

func main() {

	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6

	//Printing Decimal
	fmt.Println("\n\nDECIMAL")
	fmt.Printf("a is %d\nb is %d\nc is %d\nd is %d\ne is %d\nf is %d", a, b, c, d, e, f)

	fmt.Println("\n\nHEXADECIMAL")
	//Printing Hexadecimal
	fmt.Printf("a is %x\nb is %x\nb is %x\nc is %x\nd is %x\ne is %x", a, b, c, d, e, f)

	fmt.Println("\n\nBINARY")
	fmt.Printf("a is %b\nb is %b\nb is %b\nc is %b\nd is %b\ne is %b", a, b, c, d, e, f)

	fmt.Printf("a is %b\nb is %b\nb is %b\nc is %b\nd is %b\ne is %b", a, b, c, d, e, f)
}
