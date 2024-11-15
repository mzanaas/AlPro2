package main //2311102044
import "fmt"

func balikkan(teks []rune) []rune {
	n := len(teks)
	balik := make([]rune, n)
	for i := 0; i < n; i++ {
		balik[i] = teks[n-i-1]
	}
	return balik
}

func cekPalindrom(teks []rune) bool {
	n := len(teks)
	for i := 0; i < n/2; i++ {
		if teks[i] != teks[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan teks: ")
	fmt.Scan(&input)

	teks := []rune(input)
	hasilBalik := balikkan(teks)
	fmt.Printf("Reverse: %s\n", string(hasilBalik))

	isPalindrom := cekPalindrom(teks)
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
