package main
import "fmt"

func main() {
	var bilangan_2311102044 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan_2311102044)

	bilanganStr := fmt.Sprint(bilangan_2311102044)
	length := len(bilanganStr)

	var mid int
	if length%2 == 0 {
	mid = length / 2
	} else {
	mid = (length / 2) + 1
	}

	bilangan1 := bilanganStr[:mid]
	bilangan2 := bilanganStr[mid:]

	var num1, num2 int
	for i := 0; i < len(bilangan1); i++ {
	num1 = num1*10 + int(bilangan1[i]-'0')
	}
	for i := 0; i < len(bilangan2); i++ {
	num2 = num2*10 + int(bilangan2[i]-'0')
	}

	fmt.Println("Bilangan pemotongan:", num1, num2)
	fmt.Println("Hasil penjumlahan:", num1+num2)
}

