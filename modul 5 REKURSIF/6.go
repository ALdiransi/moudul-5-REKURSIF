package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	} else if y > 0 {
		return x * pangkat(x, y-1)
	} else {
		return 1 / pangkat(x, -y)
	}
}

func main() {
	var base, exponent int
	fmt.Print("Masukkan bilangan pokok (base): ")
	fmt.Scan(&base)
	fmt.Print("Masukkan pangkat (exponent): ")
	fmt.Scan(&exponent)

	result :=
