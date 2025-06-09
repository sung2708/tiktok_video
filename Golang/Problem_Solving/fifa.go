package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	c := a / b

	fmt.Println(c + 2022)
}
