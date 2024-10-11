package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	p1 := a - c
	var (
		hasil1 int = 1
		hasil2 int = 1
		hasilc int = 1
		hasild int = 1
		hasil3 int = 1
		hasil4 int = 1
	)

	for i := 1; i <= a; i++ {
		hasil1 *= i
	}

	for i := 1; i <= p1; i++ {
		hasil2 *= i
	}

	for i := 1; i <= c; i++ {
		hasilc *= i
	}

	p2 := b - d

	for i := 1; i <= b; i++ {
		hasil3 *= i
	}

	for i := 1; i <= p2; i++ {
		hasil4 *= i
	}

	for i := 1; i <= d; i++ {
		hasild *= i
	}

	fmt.Println(hasil1/hasil2, hasil1/(hasil2*hasilc))
	fmt.Print(hasil3/hasil4, hasil3/(hasil4*hasild))
}