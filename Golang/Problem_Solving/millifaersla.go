package main

import "fmt"

func main() {
	a, b, c := 0, 0, 0
	fmt.Scan(&a, &b, &c)

	if a < b && a < c {
		fmt.Println("Monnei")
	} else {
		if b < c {
			fmt.Println("Fjee")
		} else {
			fmt.Println("Dolladollabilljoll")
		}
	}
}
