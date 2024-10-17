package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 0 && n < 100000 {
		fmt.Print(n, " ")
		cetakDeret(n)
		fmt.Println()
	} else {
		fmt.Println("n < 0 || n > 100000")
	}
}

func cetakDeret(n int) {
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(n, " ")
	}
}