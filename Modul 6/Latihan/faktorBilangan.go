package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	fmt.Print("Faktor dari ", n, " adalah: ")
	faktorBilangan(n, 1)
}

func faktorBilangan(n, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	faktorBilangan(n, i+1)
}