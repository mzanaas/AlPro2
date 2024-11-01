package main
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	barisanBilangan(n)
}

func barisanBilangan(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	if n > 1 {
		barisanBilangan(n - 1)
		fmt.Print(n, " ")
	}
}

