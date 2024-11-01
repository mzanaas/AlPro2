package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)
	barisanGanjil(n, 1)
}

func barisanGanjil(n, current int) {
	if current > n {
		return
	}

	fmt.Print(current, " ")
	barisanGanjil(n, current + 2)
}