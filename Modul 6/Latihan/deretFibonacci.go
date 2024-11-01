package main 
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	fmt.Printf("Fibonacci ke-%d adalah %d\n", n, deretFibonacci(n))
}

func deretFibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return deretFibonacci(n-1) + deretFibonacci(n-2)
	}
}