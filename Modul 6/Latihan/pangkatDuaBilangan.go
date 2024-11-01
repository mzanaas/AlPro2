package main
import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)
	fmt.Println("Hasil:", pangkatDuaBilangan(x, y))
}

func pangkatDuaBilangan(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkatDuaBilangan(x, y-1)
}