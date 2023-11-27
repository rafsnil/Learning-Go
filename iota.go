/*
When used within a constant declaration, iota starts from 0 and increments by 1 for each subsequent
constant. This allows you to define a sequence of related constants without explicitly assigning
values to each one.
*/

package main

import "fmt"

// const (
// 	c0 = iota // c0 == 0
// 	c1 = iota // c1 == 1
// 	c2 = iota // c2 == 2
// )

//When used within a constant declaration, iota automatically increments by 1 for each subsequent constant
const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
)

func main() {
	// fmt.Println(c0, c1, c2)
	// fmt.Println(c5)
	// fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)

}
