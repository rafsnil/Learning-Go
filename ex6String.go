package main

import "fmt"

func main() {

	var name string = "niloy"
	var number int = 42
	var decimal float64 = 4.32

	fmt.Printf("%s is a %T type\n", name, name)
	fmt.Printf("%d is a %T type\n", number, number)
	fmt.Printf("%f is a %T type", decimal, decimal)

}
