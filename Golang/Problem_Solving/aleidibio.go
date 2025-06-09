package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	fmt.Println(c - a - b)
}
