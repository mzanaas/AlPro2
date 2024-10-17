package main
import "fmt"

func hitungSoalDanWaktu(input []int, waktuMaks int) (soal, waktu int) {
	for _, t := range input {
		if t < waktuMaks {
			soal++
			waktu += t
		}
	}
	return
}

func check(nama, nama2 string, a, b, c, d, e, f, g, h, a1, b1, c1, d1, e1, f1, g1, h1 int) {
	waktuMaks := 301
	dataA := []int{a, b, c, d, e, f, g, h}
	dataB := []int{a1, b1, c1, d1, e1, f1, g1, h1}

	soalA, waktuA := hitungSoalDanWaktu(dataA, waktuMaks)
	soalB, waktuB := hitungSoalDanWaktu(dataB, waktuMaks)

	if soalB > soalA {
		fmt.Println(nama2, "", soalB, waktuB)
	} else if soalA > soalB {
		fmt.Println(nama, "", soalA, waktuA)
	} else {
		if waktuA < waktuB {
			fmt.Println(nama, "", soalA, waktuA)
		} else {
			fmt.Println(nama2, "", soalB, waktuB)
		}
	}
}

func main() {
	var nama, nama2 string
	var a, b, c, d, e, f, g, h int
	var a1, b1, c1, d1, e1, f1, g1, h1 int
	fmt.Scan(&nama, &a, &b, &c, &d, &e, &f, &g, &h)
	fmt.Scan(&nama2, &a1, &b1, &c1, &d1, &e1, &f1, &g1, &h1)
	check(nama, nama2, a, b, c, d, e, f, g, h, a1, b1, c1, d1, e1, f1, g1, h1)
}
