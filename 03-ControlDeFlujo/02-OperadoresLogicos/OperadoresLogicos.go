package main

import "fmt"

func main() {
	//not
	fmt.Println(!true)

	//and
	fmt.Println(true)
	fmt.Println(true && false)
	fmt.Println(false && true)
	fmt.Println(false)

	//or
	fmt.Println(true)
	fmt.Println(true || false)
	fmt.Println(false || true)
	fmt.Println(false)

}
