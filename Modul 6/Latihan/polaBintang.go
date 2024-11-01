package main 
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)
	polaBintang(n)
}

func polaBintang(n int) {
	if n > 0 {
		polaBintang(n-1)
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}