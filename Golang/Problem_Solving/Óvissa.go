package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	sum := 0

	for _, i := range s {
		if i == 'u' {
			sum++
		}
	}
	fmt.Println(sum)
}
